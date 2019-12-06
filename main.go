package main

import (
    "log"
	"laojgo/routers"
	"laojgo/config"
)

func main()  {
	r := routers.InitRouter()
	sec, err := config.App.GetSection("server")
    if err != nil {
        log.Fatal(2, "Fail to get section 'server': %v", err)
    }
	port := sec.Key("HTTP_PORT").MustString("8080")
	r.Run(":" + port)
}
