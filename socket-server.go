package main

import (
	"./config"
	"flag"
)

func main() {
	app := &config.App{}
	port := flag.Int("port", 8080, "run application on given port")
	flag.Parse()
	app.Run(*port)
}
