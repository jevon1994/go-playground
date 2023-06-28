package main

import (
	zmq "github.com/pebbe/zmq4"
	"time"
)

type (
	ZmqSubscriber struct {
		SubAddress    string
		SubFilter     string
		HighWaterMark int
		RecvTimeout   time.Duration

		HeartbeatTime          time.Time
		HeartbeatInterval      time.Duration
		HeartbeatCheckInterval time.Duration
		ReconnectInterval      time.Duration

		Parser Parser
	}

	Parser interface {
		Parse(data [][]byte)
	}
)

func NewZmqSubscriber(subAddress string, subFilter string, highWaterMark int,
	recvTimeout, heartbeatInterval, heartbeatCheckInterval, reconectInterval time.Duration, parser Parser) *ZmqSubscriber {
	return &ZmqSubscriber{
		SubAddress:             subAddress,
		SubFilter:              subFilter,
		HighWaterMark:          highWaterMark,
		RecvTimeout:            recvTimeout,
		HeartbeatInterval:      heartbeatInterval,
		HeartbeatCheckInterval: heartbeatCheckInterval,
		ReconnectInterval:      reconectInterval,
		Parser:                 parser,
	}
}

func (this *ZmqSubscriber) Run() {
	socket, _ := zmq.NewSocket(zmq.SUB)
	//defer socket.Close()

	socket.SetRcvhwm(this.HighWaterMark)
	socket.SetSubscribe(this.SubFilter)

	// connect to service
	for {
		if err := socket.Connect(this.SubAddress); err != nil {
			//logger.Warn("Connect to subscribe address failed.",
			//	zap.Error(err),
			//	zap.String("address", this.SubAddress),
			//	zap.Duration("wait to reconnect", this.ReconnectInterval))
			time.Sleep(this.ReconnectInterval)
		} else {
			this.HeartbeatTime = time.Now()
			break
		}
	}

	nextHeartbeatCheck := time.Now().Add(this.HeartbeatCheckInterval)
	for {
		data, err := socket.RecvMessageBytes(zmq.DONTWAIT)
		if err != nil {
			time.Sleep(this.RecvTimeout)
		} else if len(data) != 2 {
			//logger.Warn("Got broken message", zap.ByteStrings("data", data))
		} else {
			// parse && send to channel
			this.Parser.Parse(data)

			// update heartbeat timestamp
			this.HeartbeatTime = time.Now()
		}

		now := time.Now()
		if nextHeartbeatCheck.Before(now) {
			nextHeartbeatCheck = nextHeartbeatCheck.Add(this.HeartbeatCheckInterval)
			if now.Sub(this.HeartbeatTime).Seconds() > (this.HeartbeatInterval.Seconds() + 1) {
				//logger.Warn("Heartbeat timeout",
				//	zap.Time("last heartbeat time", this.HeartbeatTime),
				//	zap.String("try to reconnect to", this.SubAddress))

				// reinit
				go this.Run()
				return
			}
		}
	}
}
