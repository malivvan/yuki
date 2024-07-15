package modules

import (
	"github.com/malivvan/yuki/goja"
	"github.com/malivvan/yuki/modules/buffer"
	"github.com/malivvan/yuki/modules/child_process"
	"github.com/malivvan/yuki/modules/console"
	"github.com/malivvan/yuki/modules/fs"
	"github.com/malivvan/yuki/modules/process"
	"github.com/malivvan/yuki/modules/require"
	"github.com/malivvan/yuki/modules/url"
	"github.com/malivvan/yuki/modules/util"
)

var Registry interface {
	List() []string
	Enable(*goja.Runtime)
} = func() *registry {
	r := &registry{}
	r.Registry = require.NewRegistry(require.WithLoader(require.DefaultSourceLoader))
	require.RegisterCoreModule(util.ModuleName, util.Require)
	r.coreModules(
		module{console.ModuleName, console.Require, console.Enable},
		module{buffer.ModuleName, buffer.Require, buffer.Enable},
		module{process.ModuleName, process.Require, process.Enable},
		module{url.ModuleName, url.Require, url.Enable},
		module{fs.ModuleName, fs.Require, fs.Enable},
		module{child_process.ModuleName, child_process.Require, child_process.Enable},
	)
	return r
}()

type registry struct {
	*require.Registry
	modules []module
}

type module struct {
	Name    string
	Require require.ModuleLoader
	Enable  func(*goja.Runtime)
}

func (r *registry) List() []string {
	var list []string
	for _, m := range r.modules {
		list = append(list, m.Name)
	}
	return list
}

func (r *registry) Enable(vm *goja.Runtime) {
	for _, m := range r.modules {
		m.Enable(vm)
	}
}

func (r *registry) coreModules(ms ...module) {
	for _, m := range ms {
		r.modules = append(r.modules, m)
		r.Registry.RegisterNativeModule(m.Name, m.Require)
	}
}
