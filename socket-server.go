package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//app := &config.App{}
	//port := flag.Int("port", 8080, "run application on given port")
	//flag.Parse()
	//app.Run(*port)

	i := 0
	g := make(chan interface{})

	go func() {
		i++
		data := fmt.Sprintf("%s-%d", "Greet", i)
		greet(g, data)
	}()

	go func() {
		for {
			select {
			case a := <-g:
				fmt.Println(a)
			}
		}
	}()

	var a string
	fmt.Scanln(&a)

}

func greet(c chan interface{}, data interface{}) {
	t := time.Duration(time.Second * 3)
	tick := time.Tick(t)
	for range tick {
		c <- rand.Float32()
	}

	close(c)
}
