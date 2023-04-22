package main

import (
	"fmt"
	"unsafe"
)

func main(){
	a := "hhdja"
	fmt.Println(unsafe.Sizeof(a))

	b := "sdfwesdfr"
	fmt.Println(unsafe.Sizeof(b))

	c := "asdflhkqwihefokasd"
	fmt.Println(unsafe.Sizeof(c))
}
