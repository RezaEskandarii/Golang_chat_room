package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// container
	app := &config.App{}
	
	// get port number from user's cli
	port := flag.Int("port", 8080, "run application on given port")
	flag.Parse()
	app.Run(*port)
}

 
