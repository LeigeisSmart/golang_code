// cmpout

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that we can do page 1 of the C book.

package main
import "fmt"
//具名函数
func Add0(a,b int) int{
	return a+b
}
//匿名函数
var Add = func(a,b int)int{
	return a+b
}

//多参数多返回值
func Swap(a,b int)(int,int){
	return b,a
}

// 可变数量
func Sum(a int, more ...int) int{
	for _,v:=range more{
		a+=v
	}
	return a
}

func main() {
	fmt.Println(Add0(1,2))

	fmt.Println(Add(1,2))

	fmt.Println(Swap(1,2))

	fmt.Println(Sum(1,2,3,4))
}

