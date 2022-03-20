// cmpout

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that we can do page 1 of the C book.

package main
import (
	"fmt"
	"sync"
	"sync/atomic"
)
var total uint64
func worker(wg *sync.WaitGroup){
	defer wg.Done()
	var i uint64
	for i=0;i<100;i++{
		atomic.AddUint64(&total,i)
		fmt.Println(total)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker(&wg)
	go worker(&wg)
	wg.Wait()	
}

