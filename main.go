package main

import (
	"os"
	_ "tech/config"
	_ "tech/database"
	"tech/routes"
)

func main()  {
	r := routes.InitRouter()
	port := os.Getenv("HTTP_PORT")
	r.Run(":" + port)
}
