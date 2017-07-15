package main

import "github.com/twinklesaga/worldserver/Server"

func main() {

	ctx := Server.NewContext()

	ctx.Run()
}
