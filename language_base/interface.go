// cmpout

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that we can do page 1 of the C book.

package main
import "fmt"
//定义接口
type People interface{
	GetName() string
}
type Handsome interface{
	Handsome() bool
}
//定义结构体
type Student struct{
	Name string
}

// 一个类型实现多个接口
func (stu Student) GetName() string{
	return stu.Name
}
func (stu Student) Handsome() string{
	return "yes"
}

//x.(T)断言
func CheckPeople(p interface{}){
	if _,ok:=p.(People);ok{
		fmt.Println("it is people")
	}
}
func main() {
	cbs:=Student{Name:"leige"}
	//var a People
	//a = cbs
	name:=cbs.GetName()
	fmt.Println(name)
	fmt.Println(cbs.Handsome())
	CheckPeople(cbs)

	
}

