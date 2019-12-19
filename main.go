package main

import (
	"strconv"

	"laojgo/router"
	"laojgo/config"
	"laojgo/app/model"
	"laojgo/pkg/log"
)

func main()  {
    log.Setup()
	config.Setup()
    model.Setup()

	r := router.Setup()
	r.Run(":" + strconv.Itoa(config.Server.HttpPort))
}
