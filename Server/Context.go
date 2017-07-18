package Server

import (
	"fmt"
	"net"
)

type PacketFunc func(entity *Entity, packet *Packet) error

type Context struct {
	address string
	funcMap map[Command]PacketFunc
}

func NewContext(address string) *Context {

	return &Context{address: address,
		funcMap: make(map[Command]PacketFunc)}
}

func (ctx *Context) Run() error {

	listener, err := net.Listen("tcp", ctx.address)

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {

		conn, err := listener.Accept()

		if err != nil {

		} else {

			entity := NewEntity(ctx, conn)

			fmt.Println("New Entity")
			go entity.Run()
		}
	}
	return nil
}

func (ctx *Context) AddCmdFunc(cmd Command, f PacketFunc) {

	ctx.funcMap[cmd] = f

}

func (ctx *Context) GetCmdFunc(cmd Command) (PacketFunc, error) {

	f, ok := ctx.funcMap[cmd]
	if ok {
		return f, nil
	}

	return nil, fmt.Errorf("Cannot find match function : %d", cmd)
}
