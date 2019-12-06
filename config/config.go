package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init(){
	err := godotenv.Load() // 载入 godotenv
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("HTTP_PORT") //  获取.env配置文件里的HTTP_POTRT值
	log.Print(PORT)
}
