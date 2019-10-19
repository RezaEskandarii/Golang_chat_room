package main

import (
	"./config"
)

func main() {
	app := &config.App{}
	app.Run(8080)
}
