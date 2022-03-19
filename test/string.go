//todo
//【1】:= 和 var区别


package main
import "fmt"
import "utf8"
func main() {
	//字符串数组
	var s1 = [2]string{"hello","world"}
	var s2 = [...]string{"你好","世界"}
	//var s3 = [...]string{1:"世界",0:"hello"}

	//结构体数组

	for i,v:=range s1{
		fmt.Println(i,v)
	}

	for i:=0;i<len(s2);i++{
		fmt.Println(i,s2[i])
	}


	//切片
	var s = "hello world,世界"
	fmt.Println(s[1])
	fmt.Printf("%#v\n",s)
}

