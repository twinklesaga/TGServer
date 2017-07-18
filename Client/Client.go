package main

import (
	"fmt"
	"net"
	"TGServer/Server"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9931")

	if err != nil {

	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {

	}
	defer conn.Close()

	p := Server.MakePacket(0, 0xFFFF, 0x01010101, []byte("Test"))

	fmt.Println(p.Print())

	_, err = conn.Write([]byte("dddddddddd\n"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("read")

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)

	rp, err := Server.GetPacket(reply)

	fmt.Println(string(rp.Body[:]))
}
