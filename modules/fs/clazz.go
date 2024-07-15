package fs

//
//type module struct {
//	vm *goja.Runtime
//}
//
//type class[T any] struct {
//	Module *module
//	Impl   T
//	Type   reflect.Type
//}
//
//func (m *module) createConstructor(newImpl func() any, defProto func(*goja.Object)) goja.Value {
//	f := m.vm.ToValue(func(call goja.ConstructorCall) *goja.Object {
//		res := m.vm.ToValue(newImpl()).(*goja.Object)
//		res.SetPrototype(call.This.Prototype())
//		return res
//	}).(*goja.Object)
//
//	proto := m.vm.NewObject()
//	defProto(proto)
//	f.Set("prototype", proto)
//	proto.DefineDataProperty("constructor", f, goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE)
//	return f
//}
//
//func (c *class[T]) toImpl(r *goja.Runtime, v goja.Value) T {
//	if v.ExportType() == c.Type {
//		if u := v.Export().(T); u != nil {
//			return u
//		}
//	}
//	panic(errors.NewTypeError(r, errors.ErrCodeInvalidThis, `Value of "this" must be of type %T`, c.Impl))
//}
