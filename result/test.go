package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main(){
	a := make([]byte, 6, 6)
	b := a[:0]
	fmt.Println(b, len(b), cap(b))
	b = append(b, []byte{1, 2, 3}...)
	fmt.Println(a)
	fmt.Println(b)
	c := a[:3]
	fmt.Println(cap(c))
	k := make([]byte, 4, 4)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	fmt.Printf("c: %v \n", c)
	sh.Data = uintptr(unsafe.Pointer(&k[0]))
	fmt.Printf("a: %v \n", a)
	fmt.Printf("b: %v \n", b)
	fmt.Printf("c: %v \n", c)
	fmt.Printf("k: %v \n", k)
	c[1] = 1
	fmt.Printf("c: %v \n", c)
	fmt.Printf("k: %v \n", k)
}
