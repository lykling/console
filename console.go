package console

import (
	"errors"
	"strings"
)

type Console interface {
	Name() string
	GetSignalChannel() chan int
}

type CommonConsole struct {
	state         int
	name          string
	listening     bool
	signalChannel chan int
	commandMap    map[string]Command
}

const ()

func (cc *CommonConsole) Name() string {
	return cc.name
}

func (cc *CommonConsole) GetSignalChannel() chan int {
	return cc.signalChannel
}

func (cc *CommonConsole) Parse(line string) (command Command, args []string, err error) {
	args, err = strings.Split(line, " "), nil
	name := args[0]
	if name != "" {
		command = cc.commandMap[name]
		if command == nil {
			err = errors.New("command: " + name + " not found!")
		}
	}
	return
}

func (cc *CommonConsole) start(in chan string, out chan string) {
	cc.listening = true
	cc.state = 1
	for cc.state != 0 {
		select {
		case line := <-in:
			{
				command, args, err := cc.Parse(line)
				if command != nil {
					command.Exec(args[1:]...)
				} else {
					if err != nil {
						out <- err.Error()
					}
				}
			}
		case signal := <-cc.signalChannel:
			{
				switch signal {
				case 0:
					cc.state = 0
					out <- "quit"
				}
			}
		default:
			{
			}
		}
	}
}

func (cc *CommonConsole) Listen(in chan string, out chan string) {
	if cc.listening {
		return
	}
	go cc.start(in, out)
}

func (cc *CommonConsole) AddCommand(command Command) {
	if command != nil {
		cc.commandMap[command.Name()] = command
	}
}

func (cc *CommonConsole) init() {
	command := NewCommonCommand("exit", nil, cc)
	command.SetHandler(func(ch chan int, args ...string) {
		ch <- 0
	})
	cc.AddCommand(command)
}

func NewCommonConsole(name string) (cc *CommonConsole) {
	cc = &CommonConsole{}
	cc.name = name
	cc.commandMap = make(map[string]Command, 0)
	cc.state = 0
	cc.listening = false
	cc.signalChannel = make(chan int)
	cc.init()
	return cc
}
