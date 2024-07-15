package process

import (
	"os"
	"strings"

	"github.com/malivvan/yuki/goja"
	"github.com/malivvan/yuki/modules/require"
)

const ModuleName = "process"

type Process struct {
	env map[string]string
}

func Require(runtime *goja.Runtime, module *goja.Object) {
	p := &Process{
		env: make(map[string]string),
	}

	for _, e := range os.Environ() {
		envKeyValue := strings.SplitN(e, "=", 2)
		p.env[envKeyValue[0]] = envKeyValue[1]
	}

	o := module.Get("exports").(*goja.Object)
	o.Set("env", p.env)
}

func Enable(runtime *goja.Runtime) {
	runtime.Set("process", require.Require(runtime, ModuleName))
}
