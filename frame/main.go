package main

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	inject2 "go-palyground/frame/inject"
	"log"
)

func main() {
	engine := gin.Default()
	Configure(engine)
	engine.Run()
}
func Configure(app *gin.Engine) {

	//controller declare
	//inject declare
	//Injection
	var injector inject.Graph

	err := injector.Provide(
		//&inject.Object{Value: &inject2.Api},
		&inject.Object{Value: &inject2.Service},
		//&inject.Object{Value: &inject2.StartService{}},
		&inject.Object{Value: &inject2.StartRepo{}},
	)
	if err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("inject fatal: ", err)
	}

	//database connect
	if err != nil {
		log.Fatal("db fatal:", err)
	}

}
