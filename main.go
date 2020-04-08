package main

import (
    "strconv"

    "laojgo/router"
    "laojgo/app"
    "laojgo/app/model"
    "laojgo/pkg/log"
)

func main()  {
    app.Setup()
    log.Setup()
    model.Setup()

    r := router.Setup()
    r.Run(":" + strconv.Itoa(app.Cfg.Server.HttpPort))
}
