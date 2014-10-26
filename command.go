package console

type Command interface {
	Name() string
	SetContext(context Context)
	SetHandler(handler HandlerFunc)
	Exec(args ...string)
}

type Context interface {
	GetSignalChannel() chan int
}

type HandlerFunc func(ch chan int, args ...string)

type CommonCommand struct {
	name    string
	handler HandlerFunc
	context Context
}

func (cc *CommonCommand) Name() string {
	return cc.name
}

func (cc *CommonCommand) SetHandler(handler HandlerFunc) {
	cc.handler = handler
}

func (cc *CommonCommand) Exec(args ...string) {
	if cc.handler != nil {
		context := cc.GetContext()
		if context != nil {
			go cc.handler(context.GetSignalChannel(), args...)
		} else {
			go cc.handler(nil, args...)
		}
	}
}

func (cc *CommonCommand) GetContext() (context Context) {
	return cc.context
}

func (cc *CommonCommand) SetContext(context Context) {
	cc.context = context
}

func NewCommonCommand(name string, handler HandlerFunc, context Context) (cc *CommonCommand) {
	cc = &CommonCommand{}
	cc.name = name
	cc.handler = handler
	cc.context = context
	return cc
}
