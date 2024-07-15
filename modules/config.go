package modules

import (
	"github.com/malivvan/yuki/goja"
	"github.com/malivvan/yuki/modules/eventloop"
	"io"
)

type Config struct {
	EventLoop bool
	StdIn     io.Reader
	StdOut    io.Writer
	StdErr    io.Writer
}

func VM(cfg Config, fn func(*goja.Runtime, *eventloop.EventLoop) (goja.Value, error)) (goja.Value, error) {
	var err error
	var value goja.Value
	if cfg.EventLoop {
		eventLoop := eventloop.NewEventLoop(eventloop.EnableConsole(true), eventloop.WithRegistry(Registry.(*registry).Registry))
		eventLoop.Run(func(vm *goja.Runtime) {
			Registry.Enable(vm)
			value, err = fn(vm, eventLoop)
		})
	} else {
		vm := goja.New()
		Registry.Enable(vm)
		value, err = fn(vm, nil)
	}
	return value, err
}
