package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/app-controller/connect"
)

func main() {
	_, err := connect.GetClient()
	if err != nil {
		panic(err)
	}
}
