package main

import (
	"strconv"

	"laojgo/router"
	"laojgo/config"
	"laojgo/app/model"
)

func main()  {
	config.Setup()
    model.Setup()

	r := router.Setup()
	r.Run(":" + strconv.Itoa(config.Server.HttpPort))
}
