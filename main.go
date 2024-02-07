package main

import (
	"os"
	"smartchoice/infrastructure"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := InitializedApp()

	port := os.Getenv("PORT")
	err := router.Run(port)
	if err != nil {
		panic(err)
	}

}
