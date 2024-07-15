package yuki

import (
	"github.com/malivvan/yuki/goja"
	"github.com/malivvan/yuki/goja/ast"
	"github.com/malivvan/yuki/goja/parser"
	"github.com/malivvan/yuki/modules"
	"github.com/malivvan/yuki/modules/eventloop"
)

type Program struct {
	*goja.Program
	sourceJS   string
	syntaxTree *ast.Program
}

// Compile the given project into an ES6 CommonJS module.
func Compile(name string, code string) (*Program, error) {
	var err error
	program := &Program{}
	program.sourceJS = code

	// PARSE
	program.syntaxTree, err = parser.ParseFile(nil, name, program.sourceJS, 0)
	if err != nil {
		return nil, err
	}

	// COMPILE
	program.Program, err = goja.CompileAST(program.syntaxTree, true)
	if err != nil {
		return nil, err
	}
	return program, nil
}

func Run(cfg modules.Config, program *Program) (goja.Value, error) {
	return modules.VM(cfg, func(vm *goja.Runtime, eventLoop *eventloop.EventLoop) (goja.Value, error) {
		return vm.RunProgram(program.Program)
	})
}
