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

	if port == "" {
		port = "8080"
	}
	err := router.Run("0.0.0.0:" + port)
	if err != nil {
		panic(err)
	}

}
