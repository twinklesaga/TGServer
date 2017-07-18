package Server

import (
	"bufio"
	"fmt"
	"net"
	"io"
)

type Entity struct {
	id  uint64
	ctx  *Context
	conn net.Conn
}

func NewEntity(id uint64 , ctx *Context, conn net.Conn) Entity {

	return Entity{id: id , ctx: ctx, conn: conn}
}

func (entity *Entity) Run() {

	reader := bufio.NewReader(entity.conn)

	defer entity.conn.Close()
	for {
		var p Packet
		err := p.Read(reader)

		if err != nil {
			if err == io.EOF {
				entity.ctx.disconnect <- entity.id
				break
			} else {
				entity.ctx.disconnect <- entity.id
				fmt.Println(err)
				break
			}
		} else {
			fmt.Println("Process")
			entity.Process(&p)
		}
	}
}

func (entity *Entity) Process(packet *Packet) {

	if packet != nil {
		f, err := entity.ctx.GetCmdFunc(packet.H.Cmd)
		if err == nil {
			f(entity, packet)
		}
	}
}

func (entity *Entity) Send(packet *Packet) error {

	return packet.Write(entity.conn)
}
