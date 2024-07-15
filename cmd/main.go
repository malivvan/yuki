package main

import (
	"github.com/malivvan/yuki"
	"github.com/malivvan/yuki/goja"
	"github.com/malivvan/yuki/modules"
	"github.com/malivvan/yuki/modules/eventloop"
	"os"
)

func main() {
	_, err := modules.VM(modules.Config{
		EventLoop: true,
		StdOut:    os.Stdout,
	}, func(g *goja.Runtime, loop *eventloop.EventLoop) (goja.Value, error) {
		g.RunString(`console.log("Hello, World!")`)
		return nil, nil
	})
	check(err)

	p, err := yuki.Compile("test", `console.log("Hello, World!")`)
	check(err)
	yuki.Run(modules.Config{
		StdOut:    os.Stdout,
		EventLoop: true,
	}, p)

}
func fatal(msg string) {
	println(msg)
	os.Exit(1)
}

func check(err error) {
	if err != nil {
		fatal(err.Error())
	}
}
