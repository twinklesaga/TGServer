package Server

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type Entity struct {
	ctx  *Context
	conn net.Conn
}

func NewEntity(ctx *Context, conn net.Conn) *Entity {

	return &Entity{ctx: ctx, conn: conn}
}

func (entity *Entity) Run() {

	reader := bufio.NewReader(entity.conn)

	for {

		read, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println(read)

			//if read >= 12 {
			//	p , _:= GetPacket(packet)

			fmt.Println("Process")
			fmt.Println(read)
			//	entity.Process(p)
			//}
		}
		time.Sleep(100000000000)
	}
}

func (entity *Entity) Process(packet *Packet) {

	if packet != nil {

		f, err := entity.ctx.GetCmdFunc(packet.Cmd)

		if err == nil {
			f(entity, packet)
		}
	}
}

func (entity *Entity) Send(packet *Packet) error {
	_, err := entity.conn.Write(packet.Bytes())
	return err
}
