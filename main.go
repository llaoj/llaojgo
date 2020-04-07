package main

import (
    "strconv"

    "laojgo/router"
    "laojgo/config"
    "laojgo/app/model"
    "laojgo/pkg/log"
)

func main()  {
    config.Setup()
    log.Setup()
    model.Setup()

    r := router.Setup()
    r.Run(":" + strconv.Itoa(config.App.Server.HttpPort))
}
