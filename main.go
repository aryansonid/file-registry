package main

import (
	"file-registry/config"
	"file-registry/server"
)

func main() {
	viperConfig := config.NewViperConfig("./", "config.json")

	router := server.InitRouter(viperConfig)

	router.Run("localhost:8081")
}
