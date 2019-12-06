package main

import (
	"os"
	"laojgo/config"
	"laojgo/routers"
)

func main()  {
	r := routers.InitRouter()
	sec, err := config.App.GetSection("server")
	port := sec.Key("HTTP_PORT").MustInt(8080)
	r.Run(":" + port)
}
