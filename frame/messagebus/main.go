package main

import (
	"fmt"
	messagebus "github.com/vardius/message-bus"
	"sync"
)

var bus = messagebus.New(100)
var wg sync.WaitGroup

func main() {

	wg.Add(2)

	bus.Publish("topic", "false")
	wg.Wait()
	//
	//
	//p := NewPlayer()
	//NewDailyMission()
	//NewAchievement()
	//
	//p.LevelUp()
	//p.LevelUp()
	//p.LevelUp()
	//
	//time.Sleep(100000)
}

func init() {
	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println("1=====", v)
	})

	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println("2=====", v)
	})
}

type Player struct {
	level uint32
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) LevelUp() {
	oldLevel := p.level
	newLevel := p.level + 1
	p.level++

	bus.Publish("UserLevelUp", oldLevel, newLevel)
}

type Achievement struct {
	// ...
}

func NewAchievement() *Achievement {
	a := &Achievement{}
	bus.Subscribe("UserLevelUp", a.OnUserLevelUp)
	return a
}

func (a *Achievement) OnUserLevelUp(oldLevel, newLevel uint32) {
	fmt.Printf("daily mission old level:%d new level:%d\n", oldLevel, newLevel)
}

type DailyMission struct {
	// ...
}

func NewDailyMission() *DailyMission {
	d := &DailyMission{}
	bus.Subscribe("UserLevelUp", d.OnUserLevelUp)
	return d
}

func (d *DailyMission) OnUserLevelUp(oldLevel, newLevel uint32) {
	fmt.Printf("daily mission old level:%d new level:%d\n", oldLevel, newLevel)
}
