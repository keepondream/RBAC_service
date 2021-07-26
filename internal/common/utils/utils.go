package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadingEnv() {
	APP_ENV := os.Getenv("APP_ENV")

	envFileName := fmt.Sprintf("%s.env", APP_ENV)
	secretFileName := fmt.Sprintf("%s.secret.env", APP_ENV)
	err := godotenv.Load(envFileName, secretFileName)
	if err != nil {
		log.Panicf("Loading env file error : ", err)
	}
}

func Hello() {
	fmt.Println("utils package echo hello")
}
