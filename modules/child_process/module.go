package child_process

import (
	"github.com/malivvan/yuki/goja"
	"github.com/malivvan/yuki/modules/errors"
	"github.com/malivvan/yuki/modules/require"
	"os"
	"os/exec"
	"reflect"
)

var (
	reflectTypeChildProcess = reflect.TypeOf((*childProcess)(nil))
)

func toChildProcess(r *goja.Runtime, v goja.Value) *childProcess {
	if v.ExportType() == reflectTypeChildProcess {
		if u := v.Export().(*childProcess); u != nil {
			return u
		}
	}

	panic(errors.NewTypeError(r, errors.ErrCodeInvalidThis, `Value of "this" must be of type ChildProcess`))
}

const ModuleName = "child_process"

type childProcessModule struct {
	r *goja.Runtime
}

type childProcess struct {
	command *exec.Cmd
	process *os.Process
}

func Require(runtime *goja.Runtime, module *goja.Object) {

	// module.exports = {
	// _forkChild,
	// ChildProcess,
	// exec,
	// execFile,
	// execFileSync,
	// execSync,
	// fork,
	// spawn,
	// spawnSync,
	// };
	exports := module.Get("exports").(*goja.Object)
	m := &childProcessModule{
		r: runtime,
	}
	exports.Set("ChildProcess", m.createChildProcessConstructor())
	exports.Set("spawnSync", m.spawnSync)
	//	exports.Set("URLSearchParams", m.createURLSearchParamsConstructor())
	//exports.Set("domainToASCII", m.domainToASCII)
	//	exports.Set("domainToUnicode", m.domainToUnicode)
}

func Enable(runtime *goja.Runtime) {
	m := require.Require(runtime, ModuleName)
	runtime.Set("ChildProcess", m.ToObject(runtime).Get("ChildProcess"))
	runtime.Set(ModuleName, m)
}

func (m *childProcessModule) defineChildProcessAccessorProp(p *goja.Object, name string, getter func(*childProcess) interface{}, setter func(*childProcess, goja.Value)) {
	var getterVal, setterVal goja.Value
	if getter != nil {
		getterVal = m.r.ToValue(func(call goja.FunctionCall) goja.Value {
			return m.r.ToValue(getter(toChildProcess(m.r, call.This)))
		})
	}
	if setter != nil {
		setterVal = m.r.ToValue(func(call goja.FunctionCall) goja.Value {
			setter(toChildProcess(m.r, call.This), call.Argument(0))
			return goja.Undefined()
		})
	}
	p.DefineAccessorProperty(name, getterVal, setterVal, goja.FLAG_FALSE, goja.FLAG_TRUE)
}
func (m *childProcessModule) createChildProcessPrototype() *goja.Object {
	p := m.r.NewObject()

	// host
	m.defineChildProcessAccessorProp(p, "pid", func(u *childProcess) interface{} {
		if u.process != nil {
			return u.process.Pid
		}
		return 0
	}, func(process *childProcess, value goja.Value) {

	})

	//// hash
	//m.defineURLAccessorProp(p, "hash", func(u *nodeURL) interface{} {
	//	if u.url.Fragment != "" {
	//		return "#" + u.url.EscapedFragment()
	//	}
	//	return ""
	//}, func(u *nodeURL, arg goja.Value) {
	//	h := arg.String()
	//	if len(h) > 0 && h[0] == '#' {
	//		h = h[1:]
	//	}
	//	u.url.Fragment = h
	//})
	//
	//// hostname
	//m.defineURLAccessorProp(p, "hostname", func(u *nodeURL) interface{} {
	//	return strings.Split(u.url.Host, ":")[0]
	//}, func(u *nodeURL, arg goja.Value) {
	//	h := arg.String()
	//	if strings.IndexByte(h, ':') >= 0 {
	//		return
	//	}
	//	if _, err := url.ParseRequestURI(u.url.Scheme + "://" + h); err == nil {
	//		if port := u.url.Port(); port != "" {
	//			u.url.Host = h + ":" + port
	//		} else {
	//			u.url.Host = h
	//		}
	//		m.fixURL(u.url)
	//	}
	//})
	//
	//// href
	//m.defineURLAccessorProp(p, "href", func(u *nodeURL) interface{} {
	//	return u.String()
	//}, func(u *nodeURL, arg goja.Value) {
	//	u.url = m.parseURL(arg.String(), true)
	//})
	//
	//// pathname
	//m.defineURLAccessorProp(p, "pathname", func(u *nodeURL) interface{} {
	//	return u.url.EscapedPath()
	//}, func(u *nodeURL, arg goja.Value) {
	//	p := arg.String()
	//	if _, err := url.Parse(p); err == nil {
	//		switch u.url.Scheme {
	//		case "https", "http", "ftp", "ws", "wss":
	//			if !strings.HasPrefix(p, "/") {
	//				p = "/" + p
	//			}
	//		}
	//		u.url.Path = p
	//	}
	//})
	//
	//// origin
	//m.defineURLAccessorProp(p, "origin", func(u *nodeURL) interface{} {
	//	return u.url.Scheme + "://" + u.url.Hostname()
	//}, nil)
	//
	//// password
	//m.defineURLAccessorProp(p, "password", func(u *nodeURL) interface{} {
	//	p, _ := u.url.User.Password()
	//	return p
	//}, func(u *nodeURL, arg goja.Value) {
	//	user := u.url.User
	//	u.url.User = url.UserPassword(user.Username(), arg.String())
	//})
	//
	//// username
	//m.defineURLAccessorProp(p, "username", func(u *nodeURL) interface{} {
	//	return u.url.User.Username()
	//}, func(u *nodeURL, arg goja.Value) {
	//	p, has := u.url.User.Password()
	//	if !has {
	//		u.url.User = url.User(arg.String())
	//	} else {
	//		u.url.User = url.UserPassword(arg.String(), p)
	//	}
	//})
	//
	//// port
	//m.defineURLAccessorProp(p, "port", func(u *nodeURL) interface{} {
	//	return u.url.Port()
	//}, func(u *nodeURL, arg goja.Value) {
	//	setURLPort(u, arg)
	//})
	//
	//// protocol
	//m.defineURLAccessorProp(p, "protocol", func(u *nodeURL) interface{} {
	//	return u.url.Scheme + ":"
	//}, func(u *nodeURL, arg goja.Value) {
	//	s := arg.String()
	//	pos := strings.IndexByte(s, ':')
	//	if pos >= 0 {
	//		s = s[:pos]
	//	}
	//	s = strings.ToLower(s)
	//	if isSpecialProtocol(u.url.Scheme) == isSpecialProtocol(s) {
	//		if _, err := url.ParseRequestURI(s + "://" + u.url.Host); err == nil {
	//			u.url.Scheme = s
	//		}
	//	}
	//})
	//
	//// Search
	//m.defineURLAccessorProp(p, "search", func(u *nodeURL) interface{} {
	//	u.syncSearchParams()
	//	if u.url.RawQuery != "" {
	//		return "?" + u.url.RawQuery
	//	}
	//	return ""
	//}, func(u *nodeURL, arg goja.Value) {
	//	u.url.RawQuery = arg.String()
	//	fixRawQuery(u.url)
	//	if u.searchParams != nil {
	//		u.searchParams = parseSearchQuery(u.url.RawQuery)
	//		if u.searchParams == nil {
	//			u.searchParams = make(searchParams, 0)
	//		}
	//	}
	//})
	//
	//// search Params
	//m.defineURLAccessorProp(p, "searchParams", func(u *nodeURL) interface{} {
	//	if u.searchParams == nil {
	//		sp := parseSearchQuery(u.url.RawQuery)
	//		if sp == nil {
	//			sp = make(searchParams, 0)
	//		}
	//		u.searchParams = sp
	//	}
	//	return m.newURLSearchParams((*urlSearchParams)(u))
	//}, nil)
	//
	p.Set("toString", m.r.ToValue(func(call goja.FunctionCall) goja.Value {
		//	u := toURL(m.r, call.This)
		//	u.syncSearchParams()
		return m.r.ToValue("uwu")
	}))
	//
	//p.Set("toJSON", m.r.ToValue(func(call goja.FunctionCall) goja.Value {
	//	u := toURL(m.r, call.This)
	//	u.syncSearchParams()
	//	return m.r.ToValue(u.url.String())
	//}))

	return p
}

func (m *childProcessModule) createChildProcessConstructor() goja.Value {
	f := m.r.ToValue(func(call goja.ConstructorCall) *goja.Object {
		command := exec.Cmd{}
		res := m.r.ToValue(&childProcess{command: &command}).(*goja.Object)
		res.SetPrototype(call.This.Prototype())
		return res
	}).(*goja.Object)

	proto := m.createChildProcessPrototype()
	f.Set("prototype", proto)
	proto.DefineDataProperty("constructor", f, goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE)
	return f
}
func (m *childProcessModule) spawnSync(cmd string, args []string) goja.Value {

	if ctor, ok := goja.AssertConstructor(m.r.Get("ChildProcess")); ok {

		obj, err := ctor(nil)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command(cmd, args...)
		if err := cmd.Start(); err != nil {
			return m.r.ToValue(err)
		}

		childProcObj := obj.Export()
		childProc, ok := childProcObj.(*childProcess)
		if !ok {
			return m.r.ToValue(errors.NewTypeError(m.r, errors.ErrCodeInvalidThis, `Value of "this" must be of type ChildProcess`))
		}
		childProc.command = cmd
		childProc.process = cmd.Process

		return obj
	} else {
		panic("Not a constructor")
	}
}

// module.exports = {
// _forkChild,
// ChildProcess,
// exec,
// execFile,
// execFileSync,
// execSync,
// fork,
// spawn,
// spawnSync,
// };
//const ModuleName = "child_process"
//
//type ChildProcess struct {
//	env map[string]string
//}
//
//func Require(runtime *goja.Runtime, module *goja.Object) {
//	p := &Process{
//		env: make(map[string]string),
//	}
//
//	for _, e := range os.Environ() {
//		envKeyValue := strings.SplitN(e, "=", 2)
//		p.env[envKeyValue[0]] = envKeyValue[1]
//	}
//
//	o := module.Get("exports").(*goja.Object)
//
//	//child_process.execFileSync(file[, args][, options])
//	//child_process.execSync(command[, options])
//	//child_process.spawnSync(command[, args][, options])
//	o.Set("execFileSync", p.env)
//}
//
//func Enable(runtime *goja.Runtime) {
//	runtime.Set("child_process", require.Require(runtime, ModuleName))
//}
