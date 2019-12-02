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
	// default port number is 8080
	port := flag.Int("port", 8080, "run application on given port")
	// pars flags from cli when app runs
	flag.Parse()
	// run application by port
	app.Run(*port)
}

 
