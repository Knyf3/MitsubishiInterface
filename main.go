package main

import (
	"fmt"
	"mitsuapi"
	"time"
)

func main() {

	config := mitsuapi.LoadConfiguration("config.json")

	fmt.Println(config)

	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for range ticker.C {
			fmt.Println(mitsuapi.Heartbeat())

		}
	}()
	for {
	}

	//time.Sleep(10 * time.Second)
	//ticker.Stop()
	//fmt.Println("Ticker stopped")
}
