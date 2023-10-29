package functions

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func HandleRequest(connection net.Conn) {
	defer connection.Close()
	data := make([]byte, 1024)
	connection.Read(data)

	log.Println("Request")
	// fmt.Println(string(data))
	method, file := ProcessHeader(string(data))

	switch {
	case method == "POST":
		fmt.Println("POST message: " + file)
		log.Println("POST")
	case method == "GET":
		if file == "html" || file == "html.html" {
			SendHtml(connection)
		} else {
			file, _, _ = strings.Cut(file, ".")
			SendImage(connection, file)
		}
	case file != "":
		SendError(connection)
	case method == "OPTIONS":
		connection.Write([]byte(SendOptions()))
		log.Println("Response | Options")
	default:
		connection.Write([]byte(SendError(connection)))
	}

}

func ProcessHeader(data string) (string, string) {
	before, _, _ := strings.Cut(data, " HTTP")
	method, file, _ := strings.Cut(before, " /")
	return method, file
}
