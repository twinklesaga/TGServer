package main

import (
	"fmt"
	"TGServer/Server"
)

func main() {

	ctx := Server.NewContext("127.0.0.1:9931")

	ctx.AddCmdFunc(0, Echo)
	ctx.Run()
}

func Echo(entity *Server.Entity, packet *Server.Packet) error {
	fmt.Println("Echo")

	entity.Send(packet)
	return nil
}

func Ping(entity *Server.Entity, packet *Server.Packet) error {

	fmt.Println("Pong")
	fmt.Println(packet.Body)

	entity.Send(packet)

	return nil
}
