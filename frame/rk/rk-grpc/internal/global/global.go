package global

import "github.com/pebbe/zmq4"

var (
	Sub *zmq4.Socket
)

var Tasks = make(map[string]Task)

type Task interface {
	Execute()
}
