package main

import (
	"flag"
	"fmt"
	"strings"
)

type Users []string

/* 实现Value接口即可自定义数据类型进行解析
type Value interface {
	String() string
	Set(string) error
}
*/

func (u *Users) Set(val string) error {
	*u = strings.Split(val, ",")
	return nil
}

func (u *Users) String() string {
	str := "["
	for _, v := range *u {
		str += v
	}
	return str + "]"
}

func main() {
	var u Users
	flag.Var(&u, "u", "用户列表")
	flag.Parse()

	for _, v := range u {
		fmt.Println(v)
	}
}
