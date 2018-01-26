package assert

import (
	"errors"
	"log"
	"testing"
)

type st struct {
	a int
	b int64
	c string
	e interface{}
}

func TestOK(t *testing.T) {
	True(true)
	True(1 == 1)
	False(false)
	False(1 == 0)

	err := errors.New("I am error")
	noerr := error(nil)

	True(err != nil)
	False(err == nil)
	True(noerr == nil)
	False(noerr != nil)

	v := &st{
		a: 0,
		b: 15,
		c: "hehe",
		e: nil,
	}
	var p *st

	True(v != nil)
	False(v == nil)
	True(p == nil)
	False(p != nil)
}

func TestFail(t *testing.T) {
	logFunc = log.Printf // avoid process exit.

	// see logs
	True(false)
	False(true)
}
