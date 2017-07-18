package main

import (
	"fmt"
	"net"
	"TGServer/Server"
	"time"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9931")

	if err != nil {

	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {

	}
	defer conn.Close()


	quit := make (chan bool)

	go ClientRead(quit , conn)

	Loop: for {
		var cmd , param string
		fmt.Print("입력:")
		n ,_ := fmt.Scanln(&cmd , &param)
		switch cmd {
		case "0":
			if n > 1 {
				p := Server.MakePacket(0, 0, 0, []byte(param))
				p.Write(conn)

				time.Sleep(2 * time.Second)
			}
		case "Q":
			fmt.Println("insert Q")
			quit <- true
			fmt.Println("insert Q 2")
			break Loop
		}
	}

}

func ClientRead(quit chan bool , conn net.Conn) {
	for {
		select {
		case <-quit:
			fmt.Println("Quit Client")
			return
		default:
			var rP Server.Packet
			rP.Read(conn)
			fmt.Println(string(rP.Body[:]))
		}
	}
}