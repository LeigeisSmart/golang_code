// cmpout

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that we can do page 1 of the C book.

package main
import "fmt"
func main() {
	//初始化数组
	var a[3] int;
	var b =[...]int{1,2,3}
	var c = [...] int {2:3,1:2} 
	var d = [...] int {1,2,4:5,6}  

	var times [5][0] int

	//打印int值
	fmt.Println(a[0])
	fmt.Println(b[0],b[1],b[2])
	
	//c:{0,2,3}
	for i:=range c{
		fmt.Printf("c[%d]:%d\n",i,c[i])
	}

	//又一种格式化输出方式
	for  i:=0;i<len(c);i++{
		fmt.Printf("c[%d]:%d\n",i,c[i])
	}

	// d:{1,2,0,0,5,6}
	for i,v :=range d{
		fmt.Println(i,v)
	}

	// for range 
	for range times{
		fmt.Println("hello")
	}

	//append:不能给固定size的append
	var e []int
	e = append(e,100,200,300)
	for i:=range e{
		fmt.Printf("e[%d]:%d\n",i,e[i])
	}
}

