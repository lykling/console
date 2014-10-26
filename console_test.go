package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCommonConsoleName(t *testing.T) {
	cc := NewCommonConsole("ring")
	if "ring" != cc.Name() {
		t.Error("name doesn't match.")
	}
}

func TestListen(t *testing.T) {
	cc := NewCommonConsole("ring")
	cmd := NewCommonCommand("hello", nil, cc)
	cmd.SetHandler(func(ch chan int, args ...string) {
		fmt.Fprintf(os.Stdout, "hello %s.\n", strings.Join(args, ", "))
	})
	cc.AddCommand(cmd)
	ci, co := make(chan string), make(chan string)
	go func() {
		buffer := bufio.NewReader(os.Stdin)
		_, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
		}
		for line, err := buffer.ReadString('\n'); err == nil; line, err = buffer.ReadString('\n') {
			fmt.Fprintf(os.Stdout, "%s", line)
			ci <- line
		}
		ci <- "exit"
	}()
	cc.Listen(ci, co)
	state := "running"
	for state != "quit" {
		select {
		case state = <-co:
			{
				fmt.Fprintf(os.Stdout, "%s\n", state)
			}
		}
	}
}
