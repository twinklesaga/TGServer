package Server

type Context struct {
}

func NewContext() *Context {

	return &Context{}
}

func (ctx *Context) Run() error {

	return nil
}
