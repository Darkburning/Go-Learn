package main

import (
	"fmt"
	"reflect"
	"time"
)

func makeTimeFunc(f interface{}) interface{} {
	tf := reflect.TypeOf(f)
	vf := reflect.ValueOf(f)

	if tf.Kind() != reflect.Func {
		panic("expect a function!")
	}

	wrapper := reflect.MakeFunc(tf, func(args []reflect.Value) (result []reflect.Value) {
		startTime := time.Now()
		// 函数调用
		result = vf.Call(args)
		elapsed := time.Since(startTime)
		fmt.Printf("Call Function Cost: %v\n", elapsed)
		return result
	})

	return wrapper.Interface()
}

func TimeMe() {
	time.Sleep(2 * time.Second)
}
func main() {
	timedFunc := makeTimeFunc(TimeMe).(func())
	timedFunc()
}
