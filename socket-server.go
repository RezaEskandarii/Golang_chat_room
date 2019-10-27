package main

import (
	"./config"
	"flag"
)

func main() {
	app := &config.App{}
	port:=flag.Int("port",8080,"run allication on port")
	app.Run(8080)
}
