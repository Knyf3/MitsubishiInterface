package main

import "net"
import "fmt"
import "bufio"
import "strings"

func main() {
	fmt.Println("Launching Server")

	//listen to ll interfaces
	ln, _ := net.Listen("tcp", ":8081")

	//accept connection on port
	conn, _ := ln.Accept()

	//run loop forever
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		//output message received
		fmt.Print("Message Received :", string(message))

		//sample process for string received
		newmessage := strings.ToUpper(message)

		conn.Write([]byte(newmessage + "\n"))
	}
}
