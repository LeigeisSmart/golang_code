// cmpout

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that we can do page 1 of the C book.

package main
import(
	"net/http"
	"io/ioutil"
	"log"
)
func echo(wr http.ResponseWriter,r *http.Request){
	msg,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		wr.Write([]byte("echo error"))
		return
	}
	writeLen,err:=wr.Write(msg)
	if err!=nil||writeLen!=len(msg){
		log.Println(err,"write len:",writeLen)
	}
}

func main() {
	http.HandleFunc("/",echo)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal(err)
	}
}

