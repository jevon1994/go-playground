package main

import (
	"fmt"
	"github.com/looplab/fsm"
)

var CallBack = make(map[string]fsm.Callback)

func main() {
	//door := NewDoor("heaven")
	//
	//err := door.FSM.Event("open")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = door.FSM.Event("close")
	//if err != nil {
	//	fmt.Println(err)
	//}
	CallBack["open"] = OpenCallBack
	CallBack["close"] = CloseCallBack
	fm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		CallBack,
	)

	fmt.Println("================ before open =================== ", fm.Current())
	err := fm.Event("open", "1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("================ after open =================== ", fm.Current())

	err = fm.Event("close", "2")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("================ after close=================== ", fm.Current())

}

func OpenCallBack(event *fsm.Event) {
	fmt.Println("================open call back=================== ", event.Args)

}

func CloseCallBack(event *fsm.Event) {
	fmt.Println("================close call back=================== ", event.Args)
}
