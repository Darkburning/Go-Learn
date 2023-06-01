package RpcDemo

import (
	"go/ast"
	"log"
	"reflect"
)

// 定义服务

type methodType struct {
	method    reflect.Method // 方法本身
	ArgType   reflect.Type   // 参数
	ReplyType reflect.Type   // 返回值
}

func (m *methodType) newArgv() reflect.Value {
	var argv reflect.Value
	// 入参可能是值类型也可能是指针类型
	if m.ArgType.Kind() == reflect.Ptr {
		// 创建一个指针类型的反射值(指针的指针的值类型）
		argv = reflect.New(m.ArgType.Elem())
	} else {
		// 先创建一个值类型再取指针（new得出一个指针的值类型再用Elem()再取一次指针）
		argv = reflect.New(m.ArgType).Elem()
	}
	return argv
}

func (m *methodType) newReplyv() reflect.Value {
	// 返回值必须是指针类型
	replyv := reflect.New(m.ReplyType.Elem())
	switch m.ReplyType.Elem().Kind() {
	case reflect.Map:
		// replyv本身是个指针指向的Value故可取地址修改值，且其Value也是指针
		replyv.Elem().Set(reflect.MakeMap(m.ReplyType.Elem()))
	case reflect.Slice:
		replyv.Elem().Set(reflect.MakeSlice(m.ReplyType.Elem(), 0, 0))
	}
	return replyv
}

type service struct {
	name   string                 // 映射的结构体名
	typ    reflect.Type           // 结构体类型
	rcvr   reflect.Value          // 结构体实例本身，调用时需要其作为第0个参数
	method map[string]*methodType // 存储映射的结构体的所有符合条件的方法
}

// 对 net/rpc 而言，一个函数需要能够被远程调用，需要满足如下五个条件：
//
// the method’s type is exported. – 方法所属类型是导出的。
// the method is exported. – 方式是导出的。
// the method has two arguments, both exported (or builtin) types. – 两个入参，均为导出或内置类型。
// the method’s second argument is a pointer. – 第二个入参必须是一个指针。
// the method has return type error. – 返回值为 error 类型

func newService(rcvr interface{}) *service {
	s := new(service)
	s.rcvr = reflect.ValueOf(rcvr)
	s.name = reflect.Indirect(s.rcvr).Type().Name()
	s.typ = reflect.TypeOf(rcvr)
	if !ast.IsExported(s.name) { // 判断字符串是否大写开头也即是否导出
		log.Fatalf("rpc server: %s is not a valid service name", s.name)
	}
	s.registerMethods()
	return s
}

func (s *service) registerMethods() {
	s.method = make(map[string]*methodType)
	for i := 0; i < s.typ.NumMethod(); i++ {
		method := s.typ.Method(i)
		mType := method.Type
		// 类型的入参应该有：方法本身（类似python的self）、参数值、返回值共三个；返回值应该只有一个
		if mType.NumIn() != 3 || mType.NumOut() != 1 {
			log.Printf("rpc server:register failed: Args Num Error")
			continue // 跳过该方法
		}
		// 返回值为错误类型
		if mType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
			log.Printf("rpc server:register failed: Return Value Not Error")
			continue
		}
		// 两个入参，均为导出或内置类型
		argType, replyType := mType.In(1), mType.In(2)
		if !isExportedOrBuiltinType(argType) || !isExportedOrBuiltinType(replyType) {
			log.Printf("rpc server:register failed: Not a ExportedOrBuiltinType")
			continue
		}
		s.method[method.Name] = &methodType{
			method:    method,
			ArgType:   argType,
			ReplyType: replyType,
		}
		log.Printf("rpc server:register %s.%s\n", s.name, method.Name)
	}
}

func (s *service) call(m *methodType, argv, replyv reflect.Value) error {
	f := m.method.Func
	returnValues := f.Call([]reflect.Value{s.rcvr, argv, replyv})
	if errInter := returnValues[0].Interface(); errInter != nil {
		return errInter.(error)
	}
	return nil
}

func isExportedOrBuiltinType(t reflect.Type) bool {
	// 内置类型指的是 Go 语言内置的基本数据类型，如整型、浮点型、字符型等。这些类型不存在于任何一个包中，因此其包路径为空
	return ast.IsExported(t.Name()) || t.PkgPath() == ""
}
