package main

import "fmt"

func main(){
	//a := "hhdjadja"
	//fmt.Println(unsafe.Sizeof(a))
	//
	//b := "sdfwesdfr"
	//fmt.Println(unsafe.Sizeof(b))
	//
	//c := "asdflhkqwihefokasd"
	//fmt.Println(unsafe.Sizeof(c))

	// 所以GO 允许使用
	//a := string("hello")
	//fmt.Printf("%c", a[1])

	// 但是不允许
	//a := string("hello")
	//a[1] = "w"

	// cannot assign to a[1] strings are immutable
	// 这是因为go认为string是只读的,因为是只读字符串
	// 所以Go认为字符串的编码是可以被共用的
	// 所以如果定义了一个"A世界" 和 "世界", 那么 "世界"部分的编码在Go看来是可以被公用的;
	// 如果此时给变量附上新值
	//a := string("Dota")
	// 根据上面的数据结构,只需要更爱Data字段和Len字段的值就可以了,并没有修改原来的编码
	// 所以也可以讲string强制转化为slice
	a := string("hello")
	b := ([]byte)(a)
	fmt.Println(b)
	b[2] = 65

	fmt.Printf("%c", b)
	// 这里会重新分配一段内存,再拷贝原来的内容,这样也能脱离只读限制
	// slice部分
	// 我们使用一个最基础的声明变量的方法声明一个slice会是什么结构呢?
	//var ints []int
	// 这个时候从最基础的部分分析每个字段是什么
	// Data没有分配任何元素, 所以Data = nil len = 0 cap = 0

	// 如果使用 make 定义数组不仅会定义slice结构, 还会创建一个 array
	var ints = make([]int, 2, 5)
	// 这里除了定义了一个slice机构, 还会创建一个大小为 cap==5 的array, 并且值全为0
	// 这里我把 array 理解为一段连续的内存
	// 这时 ints 的 data 会指向array的起始地址, 并且len = 2, cap = 5
	fmt.Println(ints[1])
	//fmt.Println(ints[2])
	// 执行上面两个会发生什么?
	// ints[1] 会返回0
	// ints[2] 返回 panic: runtime error: index out of range [2] with length 2
	// 添加一个元素
	ints = append(ints, 3)
	fmt.Println(ints[2], len(ints), cap(ints))
	// 返回 3 3 5
	// 基于此就大概能理解 slice 和 array 的关系, 以及 len 和 cap 真正的作用
	// 也就是说在 len 的范围内是可以安全读写的, 超出len 会发生 panic

	// 试一下 new 来创建 slice
	var pr = new([]int)
	// new 创建的 []string 和最基本方式创建的slice一样, 即不会负责底层array的创建
	// 根据之前的数据结构, pr 会作为一个指针,那么它也是data的其实地址,通过append的方式添加元素

	*pr = append(*pr, 1)
	fmt.Println(*pr)


}
