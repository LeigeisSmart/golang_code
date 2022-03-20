// cmpout

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that we can do page 1 of the C book.

package main
import "fmt"
type A struct{
	name string
}
type B struct{
	name string
}
func main() {
	a :=A{}
	a.Print()
	fmt.Println(a.name) //这个是空，因为是值传递没有改变本身
	b :=B{}
	b.Print()
	fmt.Println(b.name) //这个是bb，因为指针传递改变了结果
}
func (a A)Print(){
	a.name = "aa"
	fmt.Println("aa")
}

func (b *B)Print(){
	b.name = "bb"
	fmt.Println("bb")
}

