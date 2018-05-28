package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", "192.168.1.58:60173")

	for {
		//read stdin input
		//heartmsg := [0x11, 0x00, 0x00, 0x00]
		heartbeat := []byte{0x04, 0x41, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00}

		commandp := []byte{0x4c, 0x41, 0x00, 0x00, 0x20, 0x30, 0x30, 0x30, 0x33,
			0x41, 0x40, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

		//reader := bufio.NewReader(os.Stdin)

		//send to socket
		//fmt.Fprintf(conn, "% s", heartbeat)
		//fmt.Printf("% s", heartbeat)
		//rw, err := Open(conn)

		time.Sleep(10 * time.Second)
		fmt.Fprintf(conn, "% s", commandp)
		fmt.Printf("% s", commandp)
		time.Sleep(10 * time.Second)
		fmt.Fprintf(conn, "% s", heartbeat)
		fmt.Printf("% s", heartbeat)
		//listen for reply
		//message, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Print("Message from server : " + message)

	}

}
