# go-data-structure

go 的数据结构

# string
首先要先了解机string类型在机器编码种的形式

    1、 1 bit 可以是 0/1  8 bit = 1 byte， 可以表示255个数，越多的bit可以表示越大的数字
    2、 stirng类型在机器中是怎么存储的？ --> “A” 在机器中通过编码会用 65 来表示， 01000001
    3、 那怎么表示 "AA" 呢？ 如果不用任何措施就无法知道怎么终止，比如说 0100 0001 0100 0001 要怎么和一个完整的字符区分呢？
    4、 首先最容易想到的一个方法就是使用定长编码，比如说利用编码种最高位来表示所有的字符，也就是说假设现在编码的规则是16个字节，那么所有的
        编码我们都要按照16字节来表示，所以 "AA" 就会表示成：
        0000 0000 0100 0001 0000 0000 0100 0001；
        也很容易发现这个方法会很明显造成内存资源的浪费
    5、 于是出现了变长编码，用0开头表示只占一个字节，110开头占2个字节，1110开头占3个字节，同时除了第一个字节中的标识位，占用多少字节就会
        在后面每个字节开头用10占位：
        比如 用 0100 0001 除去标识符0 剩下 1000001，转化为10进制为65，表示A;
        比如 用 1110 0100 1011 1000 1001 0110 除去 1110 的表示为，表示占用3个bit；再出去后面2个字节的10，剩下的 0100 1110 0001 0110
        转化为十进制是 19990 表示 ”世“
        
        ”界“ 的编码为 30028  二进制为 1110 1010 1001 100 一共15位，至少加上三字节对应的占位/表示为（4+2+2） 一共23位，刚好可以被装下
        为了填满字节所以在前面补0(在最前面补，不会影响值)，所以表示为
        [1110 0111] [10 010101] [10 001100]

>世：
1110 0100 1011 1000 1001 0110
> 
>[1110 0100] [10 111000] [10 010110]
> 
>1110
> 
>0100 1110 0001 0110
> 
>界：
>1110 1010 1001 100  至少加上三字节对应的占位/表示为（4+2+2） 一共23位
>
>[1110 0111] [10 010101] [10 001100]
> 
> 如果表示4个字节就是 11110 的标识符  这里标识符不一定只能四位, 最长能表示到 1111110


    二 现在就可以看字符串是什么结构了
    1 首先要有一段编码的起始地址
    2 知道了开头位置,还需要知道在哪里结尾
    3 C语言中使用\0来表示一个字符串的终止,但是这样会出现问题就是内容中不能出现\0
    4 Go 选择在起始地址后面存储了一个int类型的len字段用于表示string的长度

```go
// test code

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
```
![img.png](string_test.png)
这里我们可以看到这个 a,b,c 的 string 类型的长度都为 16 , 明明我们给它们初始化的长度并不是16

其实在Go里面, 它其实是这样表示的这里使用过一个Data 表示这个 string 的一个开头, 这个是个指针,指向那段编码, 那段编码会存储在另外一段位置, 
就是string类型并不会把这个字符串的编码包裹起来, string 类型只是一个结构体,包含了指向那段编码的指针和它的长度信息
![img.png](string_struct.png)
```go
//字符串是字符串的运行时表示形式。
//它不能安全或便携使用，其表示可能
//在以后的版本中进行更改。
//
//与reflect.StringHeader不同，它的Data字段足以保证
//它引用的数据不会被垃圾收集。

type string struct {
    Data unsafe.Point
    Len  int
}

// 这一段代码在 internal\ursafeheader\unsafeheader.go 种被定义
// 那么A世界在Go种的存储方式就是
// Data 指向编码的起始地址;
// 编码为 [0100 0001] [1110 0100 1011 10000 1001 0110] [1110 0111 1001 0101 1000 1100]
// 在Data后面还存有 int 类型的 7 表示编码的字节数
```
示例图
![img.png](string_show_img.png)
这里data只是一个指针指向真正存储位置的开头, len表示这个字符串占的长度

这里长度为16是因为指针长度为8, 字符长度为8, 所以组合起来他们的长度就为8

Go string 注意点:
```go
// 所以GO 允许使用
a := string("hello")
fmt.Printf("%c", a[1])

// 但是不允许
a := string("hello")
a[1] = "w"

// cannot assign to a[1] strings are immutable
// 这是因为go认为string是只读的,因为是只读字符串
// 所以Go认为字符串的编码是可以被共用的
// 所以如果定义了一个"A世界" 和 "世界", 那么 "世界"部分的编码在Go看来是可以被公用的;
// 如果此时给变量附上新值
a := string("Dota")
// 根据上面的数据结构,只需要更爱Data字段和Len字段的值就可以了,并没有修改原来的编码
// 所以也可以讲string强制转化为slice
a := string("hello")
b := ([]byte)(a)
fmt.Println(b)
// [104 101 108 108 111]
b[2] = 65
fmt.Printf("%c", b)
// [h e A l o]
// 这里会重新分配一段内存,再拷贝原来的内容,这样也能脱离只读限制
// slice部分
```

# slice
    slice 是什么结构? 元素存在哪里?存了多少元素 len? 可以存多少 buf?

```go
//切片是切片的运行时表示。
//它不能安全或便携使用，其表示可能
//在以后的版本中进行更改。
//
//与reflect.SliceHeader不同，它的Data字段足以保证
//它引用的数据不会被垃圾收集。
type Slice struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
// 这一段代码在 internal\ursafeheader\unsafeheader.go 种被定义
// 我也不确定就只在这里有定义  就用做了解
// Data 是一个指针  指向一个array
// len表示slice长度 有效区域, cap表示支持存多少 我理解为一个buf缓冲区
// 我们使用一个最基础的声明变量的方法声明一个slice会是什么结构呢?
var ints []int
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
```
# array
    数组: int 型的 slice 底层就是 int 型的数组, 但是 slice 的 data 并不一定指向 array 的开头

```go
// array
// 比如:
var inta = [5]int{1, 2, 3, 4, 5}
ints := inta[2:3]
// 当 array 已经存在了, 并且基于已经存在的 array 创建slice 就不会指向 array 的起始地址
// 同理 还可以把其他的 slice 关联到同一个数组
ints1 := inta[1:4]
ints2 := inta[0:3]
ints[0] = 100
fmt.Println(inta, ints, ints1, ints2)
// [1 2 100 4 5] [100] [2 100 4] [1 2 100]
// 这也是为什么修改一个slice有时候会同时改变其他的slice的原因
// 因为slice本身并没有保存数据, 只是保存了 开头(data) 和 结尾(len)
```

数组越界
```go
var inta = [5]int{1, 2, 3, 4, 5}
ints := inta[0:]
fmt.Println(cap(ints))
// 这个时候 使用 cap(ints)查看  会发现 cap==5, 此时如果要append一个元素会怎么样
ints = append(ints, 6)
fmt.Println(cap(ints))
// 这里我们会发现cap变成了10
// 先不去考虑cap的问题, 我们知道array在内存中是一个连续的一段,并且不能扩大;
// 那么当slice需要表示的len超过了array就会重新给slice创建一个新的array, 再将元数据拷贝过去
// 至此就能理解为什么会出现cap变成10的原因了;
// 因为slice是可以扩大的, 如果没append一次就要重新创建数组再copy回来, 那么对于性能的损耗就会比较大
// 所以 Go 对slice的扩容做了优化
```

newcap 大小的数组需要分配多大的内存?
> newcap * 元素类型大小 == 分配的内存大小?
> 
> 再Go中 申请内存并不会直接和操作系统直接交互, 而是由Go一次性申请一大块内存, 再由Go按照需要分配
> 
> 这里涉及到了Go的malloc 这个环节比较多, 以后继续学习
> 
> 简单来说就是会匹配到足够大切接近的规格(8\16\32...) -- 类似于操作系统里的一种分块方法, 定义一些规格然后去匹配最接近的规格
# 内存对齐
# map
