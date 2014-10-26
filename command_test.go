package console

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type testContext struct {
	sc chan int
}

func (tc *testContext) GetSignalChannel() chan int {
	return tc.sc
}

func newTestContext() (tc *testContext) {
	tc = &testContext{}
	tc.sc = make(chan int)
	return
}

var (
	context = newTestContext()
)

func hello(ch chan int, args ...string) {
	if ch != nil {
	}
	fmt.Fprintf(os.Stdout, "hello %s.\n", strings.Join(args, ", "))
}

func TestCommonCommandName(t *testing.T) {
	cc := NewCommonCommand("ring", hello, nil)
	if "ring" != cc.Name() {
		t.Error("name doesn't match.")
	}
}

func TestCommonCommandSetHandler(t *testing.T) {
	cc := NewCommonCommand("ring", nil, context)
	cc.SetHandler(hello)
	cc.Exec([]string{"asteria", "eunomia"}...)
}

func TestCommonCommandExec(t *testing.T) {
	cc := NewCommonCommand("ring", hello, context)
	cc.Exec([]string{"asteria", "eunomia"}...)
}
