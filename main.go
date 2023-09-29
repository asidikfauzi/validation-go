package main

import (
	"test-prepare/app"
	"test-prepare/routes"
)

func main() {
	app.InitConfig()

	routes.NewRouter()
}
