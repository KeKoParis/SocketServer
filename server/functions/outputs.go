package functions

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func SendHtml(connection net.Conn) {
	file, err := os.ReadFile("html.html")
	if err != nil {
		fmt.Println(err.Error())
		SendError(connection)
		log.Println("ERROR: Invalid request html | GET")
		return
	}
	header := fmt.Sprintf("HTTP/1.1 200 \r\nContent-Type: text/html; charset=utf-8\r\nSize: %d\r\n\r\n", len(file))
	connection.Write([]byte(header))

	for i := 0; i <= len(file)/1024; i++ {
		if (i+1)*1023 > len(file) {
			connection.Write(file[i*1023:])
		} else {
			connection.Write(file[i*1023 : (i+1)*1023])
		}
	}

	time.Sleep(10 * time.Second)

	log.Println("Response html | GET")
}

func SendImage(connection net.Conn, imgType string) {
	file, err := os.ReadFile(imgType + "." + imgType)
	fmt.Println(imgType)
	if err != nil {
		fmt.Println(err.Error())
		SendError(connection)
		log.Println("ERROR: Invalid request image | GET")
		return
	}
	fmt.Println(imgType)
	if imgType == "svg" {
		imgType += "+xml"
	}
	header := fmt.Sprintf("HTTP/1.1 200 \r\nContent-Type: image/%s;\r\nSize: %d\r\n\r\n", imgType, len(file))
	connection.Write([]byte(header))
	fmt.Println(header)

	for i := 0; i <= len(file)/1024; i++ {

		if (i+1)*1024 > len(file) {
			connection.Write(file[i*1024:])
		} else {
			connection.Write(file[i*1024 : (i+1)*1024])
		}
	}
	time.Sleep(10000 * time.Millisecond)

	log.Println("Response to image | GET")
}

func SendError(connection net.Conn) string {
	header := "HTTP/1.1 404 Not Found\r\nContent-Type: text/html\r\n\r\n"

	errorHtml, err := os.ReadFile("error.html")
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	connection.Write([]byte(header))
	connection.Write(errorHtml)

	return header
}

func SendOptions() string {
	header := "HTTP/1.1 200 OK\r\nAllow: GET, POST, OPTIONS\r\nServer: My-Little-Server"
	return header
}
