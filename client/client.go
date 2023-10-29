package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	host    = "localhost"
	port    = "5005"
	netType = "tcp"
)

func main() {
	arg1 := flag.String("method", "GET", "method: GET, POST, OPTIONS")
	arg2 := flag.String("type", "", "file to get")

	flag.Parse()
	method := *arg1
	dataType := *arg2

	addr, err := net.ResolveTCPAddr(netType, host+":"+port)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	connection, err := net.DialTCP(netType, nil, addr)

	connection.Write([]byte(method + " /" + dataType + " HTTP/1.1"))

	data := make([]byte, 1024)
	connection.Read(data)
	fmt.Println(string(data))

	_, sizeS, _ := strings.Cut(string(data), "Size: ")
	sizeS, _, _ = strings.Cut(sizeS, "\r\n\r\n")
	sizeInt, _ := strconv.Atoi(sizeS)

	if method == "GET" {

		var fileData []byte
		for i := 0; i <= sizeInt/1024; i++ {
			_, err = connection.Read(data)
			fileData = append(fileData, data...)
		}

		fileData = fileData[:sizeInt]

		file, _ := os.OpenFile(dataType+"."+dataType, os.O_CREATE|os.O_RDWR, 0666)

		file.Write(fileData)

		file.Close()

	}

}
