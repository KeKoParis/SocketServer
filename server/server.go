package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"server/functions"
)

const (
	host    = "localhost"
	port    = "5005"
	netType = "tcp"
)

func main() {
	logFile, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Sever started")

	listen, err := net.Listen(netType, host+":"+port)
	if err != nil {
		fmt.Println("Listen error", err.Error())
		os.Exit(1)
	}
	defer listen.Close()

	log.Println("Listening connections")

	for {
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println("connection error", err.Error())
			return
		}
		log.Println("Connection accepted")

		go functions.HandleRequest(connection)
	}

}
