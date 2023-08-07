package task

import (
	"fmt"
	"rk-grpc/internal/global"
)

type ReceiverTask struct{}

func (rt *ReceiverTask) Execute() {
	//接收广播
	fmt.Println("invoke receive======================>")
	resp, err := global.Sub.RecvBytes(0)
	if err == nil {
		//resp = resp[len(""):] //去掉前缀
		fmt.Printf("receive [%s] ============================== \n", resp)
		//if resp == "END" {
		//	break
		//}
	} else {
		fmt.Println(err)
	}
}

func init() {
	global.Tasks["receiver"] = &ReceiverTask{}
}
