package main

import (
	"fmt"
	"log"

	"gopkg.in/zeromq/goczmq.v4"
)

func main() {
	router, err := goczmq.NewRouter("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer router.Destroy()
	log.Println("Router created and bound")

	addr := []byte("")
	msg := []byte("")
	for {
		request, more, err := router.RecvFrame()
		if err != nil {
			log.Fatal(err)
		}
		if more != 0 {
			addr = request
			continue
		} else {
			msg = request
		}
		log.Printf("Echoed %v\n", addr)
		fmt.Printf("%s [%d]\n", string(msg), len(msg))

		err = router.SendFrame(addr, goczmq.FlagMore)
		if err != nil {
			log.Fatal(err)
		}
		err = router.SendFrame(msg, goczmq.FlagNone)
		if err != nil {
			log.Fatal(err)
		}
	}
}
