package console

// Command interface
type Command interface {
	Name() string
	SetContext(context Context)
	SetHandler(handler HandlerFunc)
	Exec(args ...string)
}

// Context interface
type Context interface {
	GetSignalChannel() chan int
}

// HandlerFunc shorcut for handler type
type HandlerFunc func(ch chan int, args ...string)

// CommonCommand implement of Command
type CommonCommand struct {
	name    string
	handler HandlerFunc
	context Context
}

// Name return command name
func (cc *CommonCommand) Name() string {
	return cc.name
}

// SetHandler set command handler
func (cc *CommonCommand) SetHandler(handler HandlerFunc) {
	cc.handler = handler
}

// Exec execute command
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

// GetContext return context
func (cc *CommonCommand) GetContext() (context Context) {
	return cc.context
}

// SetContext set context
func (cc *CommonCommand) SetContext(context Context) {
	cc.context = context
}

// NewCommonCommand create instance of CommonCommand
func NewCommonCommand(name string, handler HandlerFunc, context Context) (cc *CommonCommand) {
	cc = &CommonCommand{}
	cc.name = name
	cc.handler = handler
	cc.context = context
	return cc
}
