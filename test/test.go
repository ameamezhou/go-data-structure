package main

import (
	"fmt"
)

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
	//var ints = make([]int, 2, 5)
	// 这里除了定义了一个slice机构, 还会创建一个大小为 cap==5 的array, 并且值全为0
	// 这里我把 array 理解为一段连续的内存
	// 这时 ints 的 data 会指向array的起始地址, 并且len = 2, cap = 5
	//fmt.Println(ints[1])
	//fmt.Println(ints[2])
	// 执行上面两个会发生什么?
	// ints[1] 会返回0
	// ints[2] 返回 panic: runtime error: index out of range [2] with length 2
	// 添加一个元素
	//ints = append(ints, 3)
	//fmt.Println(ints[2], len(ints), cap(ints))
	// 返回 3 3 5
	// 基于此就大概能理解 slice 和 array 的关系, 以及 len 和 cap 真正的作用
	// 也就是说在 len 的范围内是可以安全读写的, 超出len 会发生 panic

	// 试一下 new 来创建 slice
	var pr = new([]int)
	// new 创建的 []string 和最基本方式创建的slice一样, 即不会负责底层array的创建
	// 根据之前的数据结构, pr 会作为一个指针,那么它也是data的其实地址,通过append的方式添加元素

	*pr = append(*pr, 1)
	fmt.Println(*pr)

	// array
	// 比如:
	//var inta = [5]int{1, 2, 3, 4, 5}
	//ints := inta[2:3]
	// 当 array 已经存在了, 并且基于已经存在的 array 创建slice 就不会指向 array 的起始地址
	// 同理 还可以把其他的 slice 关联到同一个数组
	//ints1 := inta[1:4]
	//ints2 := inta[0:3]
	//ints[0] = 100
	//fmt.Println(inta, ints, ints1, ints2)
	// [1 2 100 4 5] [100] [2 100 4] [1 2 100]
	// 这也是为什么修改一个slice有时候会同时改变其他的slice的原因
	// 因为slice本身并没有保存数据, 只是保存了 开头(data) 和 结尾(len)


	// 数据越界怎么办
	var inta = [5]int{1, 2, 3, 4, 5}
	ints := inta[0:]
	fmt.Println(cap(ints))
	// 这个时候 使用 cap(ints)查看  会发现 cap==5, 此时如果要append一个元素会怎么样
	ints = append(ints, 6)
	fmt.Println(cap(ints))
	// 这里我们会发现cap变成了10
	// 先不去考虑cap的问题, 我们知道array在内存中是一个连续的一段,并且不能扩大;
	// 那么当slice需要表示的len超过了array就会重新给slice创建一个新的array, 再将原数据拷贝过去
	// 至此就能理解为什么会出现cap变成10的原因了;
	// 因为slice是可以扩大的, 如果没append一次就要重新创建数组再copy回来, 那么对于性能的损耗就会比较大
	// 所以 Go 对slice的扩容做了优化

}
