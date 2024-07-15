package fs

import (
	"github.com/malivvan/yuki/goja"
	"github.com/malivvan/yuki/modules/require"
	"os"
)

const ModuleName = "fs"

type fsModule struct {
	r   *goja.Runtime
	fds map[int]*os.File
}

func Require(runtime *goja.Runtime, module *goja.Object) {

	exports := module.Get("exports").(*goja.Object)
	m := &fsModule{
		r: runtime,
	}
	m.defineConstants(exports)

}

func Enable(runtime *goja.Runtime) {
	m := require.Require(runtime, ModuleName)
	//runtime.Set("ChildProcess", m.ToObject(runtime).Get("ChildProcess"))
	runtime.Set(ModuleName, m)
}
