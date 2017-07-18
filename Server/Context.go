package Server

import (
	"fmt"
	"net"
)

type PacketFunc func(entity *Entity, packet *Packet) error

type Context struct {
	address string
	baseID uint64
	disconnect chan uint64
	entityMap map[uint64]Entity
	funcMap map[Command]PacketFunc
}

func NewContext(address string) *Context {

	return &Context{address: address,
		baseID: 100000,
		entityMap: make(map[uint64]Entity),
		funcMap: make(map[Command]PacketFunc) ,
		}
}

func (ctx *Context) Run() error {

	listener, err := net.Listen("tcp", ctx.address)

	if err != nil {
		panic(err)
	}

	ctx.disconnect = make(chan uint64)
	go func() {
		id := <- ctx.disconnect
		fmt.Println("disconnect" , id)

		_,ok := ctx.entityMap[id]

		if ok {
			delete(ctx.entityMap , id)
		}
	}()

	defer listener.Close()

	for {

		conn, err := listener.Accept()

		if err != nil {

		} else {
			entity := NewEntity(ctx.baseID , ctx, conn)
			ctx.entityMap[ctx.baseID] = entity

			fmt.Println("New Entity" , ctx.baseID )
			go entity.Run()

			ctx.baseID++
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
