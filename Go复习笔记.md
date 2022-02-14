# Go复习笔记

## day01(2022年2月14日 上午)：

内容：

1. 复习Go基础

### 1. Go语言的主要特性：

1. 自动立即回收
2. 更丰富的内置类型
3. 函数多返回值
4. 错误处理
5. 匿名函数和闭包
6. 类型和接口
7. 并发编程
8. 反射
9. 语言交互性

### 2. 可见性

1. 声明在函数内部，是函数的本地值，类似于private
2. 声明在函数外部，是对当前包可见的全局值，类似protect
3. 声明在函数外部且首字母大写是所有包可见的全局变量，类似于public

### 3. 语言声明

var(声明变量)

const(声明常量)

type(声明类型)

func(声明函数)

### 内置类型和函数

#### 1. 内置类型

##### 1.1 值类型

bool, 

int(32 or 64), int8, int16, int32, int64

uint(32 or 64), uint8(byte),uint16,uint32,uint64

float32, float64

string

complex64, complex128

array -- 固定长度的数组

##### 1.2引用类型：（指针类型）

slice -- 序列数组(最常用)

map -- 映射

chan -- 管道

#### 2.内置函数

一些不需要导包，可以直接获得编译器支持的函数

- append --用来追加元素到数组。slice中，返回修改后的数组，slice
- close --主要用来关闭channel
- delete -- 从map中删除key对应的value
- panic -- 停止常规的goroutine （panic和recover：用来错误处理）
- real -- 返回complex的实部(complex, real imag:用来创建和操作复数)
- imag -- 返回complex的虚部
- make --用来分配内存，返回type本身(只能用于slice，map，channel)
- new -- 用来内存分配，主要用来分配值类型，比如int，struct。返回Type的指针
- cap -- capacity是容量的意思，用于返回某个类型的最大容量，只能运用于切片和map
- copy -- 用于复制和连接slice，返回复制的数目
- len -- 来求长度，比如string，array，slice，map，channel，返回长度
- print，println -- 底层打印函数，在部署环境中建议使用fmt包

### Init函数和main函数

#### 1.init函数

Go语言中init函数用于包的初始化，该函数的特性：

- init函数是用于函数执行前做包的初始化函数，比如初始化包里的变量等
- 每个包可以用于多个init函数
- 包的每个源文件也可以拥有多个init函数
- 同一个包中多个init函数的执行顺序go语言没有明确的说明
- 不同包的init函数按照包导入的依赖关系决定该初始化的执行顺序
- init函数不能被其他函数调用，而是在main函数执行之前，自动被调用

#### 2.main函数

Go语言程序的默认入口

```go
func main() {
    
}
```

#### 3. init函数和main函数的异同

- 相同点：
  - 两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用
- 不同点：
  - init可以应用于任意包中，且可以重复定义多个
  - main函数只能用于main包中，且只能定义一个

### Go命令

```go
go env用于打印Go语言的环境信息。

go run命令可以编译并运行命令源码文件。

go get可以根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装。

go build命令用于编译我们指定的源码文件或代码包以及它们的依赖包。

go install用于编译并安装指定的代码包及它们的依赖包。

go clean命令会删除掉执行其它命令时产生的一些文件和目录。

go doc命令可以打印附于Go语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的。

go test命令用于对Go语言编写的程序进行测试。

go list命令的作用是列出指定的代码包的信息。

go fix会把指定代码包的所有Go语言源码文件中的旧版本代码修正为新版本的代码。

go vet是一个用于检查Go语言源码中静态错误的简单工具。

go tool pprof命令来交互式的访问概要文件的内容。
```

### 下划线

`_`是特殊的标识符，用来忽略结果

```go	
package main

import (
    "os"
)

func main() {
    buf := make([]byte, 1024)
    f, _ := os.Open("/Users/***/Desktop/text.txt")
    defer f.Close()
    for {
        n, _ := f.Read(buf)
        if n == 0 {
            break    

        }
        os.Stdout.Write(buf[:n])
    }
}

// 解释
    下划线意思是忽略这个变量.
    比如os.Open，返回值为*os.File，error
    普通写法是f,err := os.Open("xxxxxxx")
    如果此时不需要知道返回的错误值
    就可以用f, _ := os.Open("xxxxxx")
    如此则忽略了error变量

// 另外一种解释
占位符，意思是那个位置本应赋给某个值，但是咱们不需要这个值。
    所以就把该值赋给下划线，意思是丢掉不要。
    这样编译器可以更好的优化，任何类型的单个值都可以丢给下划线。
    这种情况是占位用的，方法返回两个结果，而你只想要一个结果。
    那另一个就用 "_" 占位，而如果用变量的话，不使用，编译器是会报错的。
```

### 变量

#### 变量声明

变量定义格式`var 变量名 变量类型`，行尾无需分号

```go
var name string
var age int
var isOk bool
```

#### **批量声明：**

go语言中还支持批量变量声明

```go
var (
	a string
    b int
    c bool
    d float32
)
```

#### 变量初始化

 整型和浮点型变量的默认值为0。 字符串变量的默认值为空字符串。 布尔型变量默认为`false`。 切片、函数、指针变量的默认为`nil`。

```go
// 格式如下
var 变量名 类型 = 表达式
// 例子
var name string = "suyu"
var sex int = 1
// 一次可以初始化多个值
var name, sex = "suyu", 1
```

#### 类型推导

我们在定义变量的时候可以省略类型，编译器会根据等号右边的值来推导办理的类型完成初始化

```go
var name = "suyu"
var sex = 1
```

#### 短变量声明

在函数内部，可以简单的使用`:=`方式声明并初始化变量。

```go
package main

import "fmt"

// 全局变量
var m = 100

func main() {
	n := 10
	m := 200 // 此时声明变量为m
	fmt.Println(n, m)
}
```

### 常量

常量的声明：

```go
const pi = 3.1415926
const e = 2.7182
```

声明的常量在未来程序运行期间它的值都不会发生改变了。
同时常量也可以像变量一样多个变量一起定义：

```go
const (
	pi = 3.1415926
    e = 2.7182
)

// 当const同时声明多个变量的时候，如果省略了值则表示和上面一行的值相同。
const (
	n1 = 100
    n2
    n3
)
// 这样的话n1, n2, n3 都是100
```

#### iota

`iota`是`go`语言的常量计数器，只能在常量的表达式中使用。`iota`在`const`关键字出现时将被重置为`0`。`const`中每新增一行常量声明将使`iota`计数一次(`iota`可理解为`const`语句块中的行索引)。 使用`iota`能简化定义，在定义枚举时很有用。

```go
const (
	n1 = iota // 0
    n2 // 1
    n3 // 2
    n4 // 3
)
```

使用`_`跳过某些值

```go
const (
	n1 = iota // 0
    n2 // 1
    _
    n4 // 3
)
```

`iota`声明中间插队

```go
const (
	n1 = iota // 0
    n2 = 100 // 100
    n3 = iota // 2
    n4 // 3
)
const n5 = iota // 0
```

定义数量级 （这里的`<<`表示左移操作，`1<<10`表示将`1`的二进制表示向左移`10`位，也就是由`1`变成了`10000000000`，也就是十进制的`1024`。同理`2<<2`表示将`2`的二进制表示向左移`2`位，也就是由`10`变成了`1000`，也就是十进制的`8`。）

```go
const (
	_ = iota
    KB = 1 << (10*iota)
    MB = 1 << (10 * iota)
    GB = 1 << (10 * iota)
    TB = 1 << (10 * iota)
    PB = 1 << (10 * iota)
)
```

多个`iota`定义在一行

```go
const (
	a, b = iota + 1, iota + 2 // 1, 2
    c, d // 2, 3
    e, f // 3, 4
)
```

### 字符串的常用操作

| 方法                                 | 介绍           |
| ------------------------------------ | -------------- |
| len(str)                             | 求长度         |
| + 或 fmt.Sprintf                     | 拼接字符串     |
| strings.Split                        | 分割           |
| strings.Contains                     | 判断是否包含   |
| strings.HasPrefix, strings.HasSuffix | 前缀/后缀判断  |
| strings.Index(), strings.LastIndex() | 子串出现的位置 |
| strings.Join(a[] stirng, sep string) | join操作       |

### byte(uint8)和rune(int32)类型

组成每个字符串的元素叫做"字符"，可以通过变量或者单个获取字符串元素获得字符。字符用单引号`'`'包裹起来

```go
var a = '中'
var b = 'x'
```

Go语言的字符有以下两种

```go
uint8类型，或者叫byte型，代表了ASCII码的一个字符
rune类型， 代表一个UTF-8字符
```

代码`variableString\main.go`

当需要处理中文、日文或者其他复合字符时，则需要用到`rune`类型。`rune`类型实际是一个`int32`。 Go 使用了特殊的 `rune` 类型来处理 `Unicode`，让基于 `Unicode`的文本处理更为方便，也可以使用 `byte` 型进行默认字符串处理，性能和扩展性都有照顾。

因为UTF8编码下一个中文汉字由`3~4`个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。

字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。

#### 修改字符串

要修改字符串，要先将其转化成`[]rune`和`[]byte`，完成后再转化为`string`。无论那种转化，都会重新分配内存，并复制字节数组

```go
package main

import "fmt"

func changeString() {
	s1 := "hello"
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	s1 = string(byteS1)
	fmt.Println(s1)

	s2 := "我是个程序员!"
	runeS2 := []rune(s2)
	runeS2[0] = '你'
	fmt.Println(string(runeS2))
}
func main() {
	changeString()
}
```

#### 类型转换

Go语言中只有强制类型转化，没有隐式类型转化，该语法只能在两个类型之间支持相互转化的时候使用。

语法：`T(表达式)`



### 数组练习：

1. 求数组所有元素之和
2.  找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，找出两个元素之和等于8的下标分别是（0，4）和（1，2）



### 切片Slice

slice不是数组或数组指针，它通过内部指针和相关属性引用数组片段，以实现变长方案

```	
1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递
2. 切片的长度可以改变，因此，切片是一个可变的数组
3. 切片遍历和数组一样，可以用len()求长度。表示元素数量，读写操作不能超过该限制。
4. cap可以求出slice最大扩张容量，不能超出数组限制。0<=len(slice)<=len(array),其中array是slice引用的数组
5. 切片的定义：var 变量名 []类型，比如 var str []string
6. 如果slice == nil, 那么len，cap结构都等于0
```

#### 创建切片的各种方式

```go
package main

import "fmt"

func main() {
    // 1. 声明切片
	var s1 []int
    if s1 == nil {
        fmt.Println("是空")
    } else {
        fmt.Println("不是空")
    }
    // 2. :=
    s2 := []int{}
    // 3. make
    var s3 []int = make([]int, 0)
    // 4. 初始化赋值
    var s4 []int = make([]int, 0, 0)
    fmt.Println(s4)
    s5 := []int{1, 2, 3}
    fmt.Println(s5)
    // 5. 从数组切片
    arr := [5]int{1, 2, 3, 4, 5}
    var s6 []int
    // 前包含后不包含
    s6 = arr[1:4]
    fmt.Println(s6)
}
```

| 操作            | 含义                                                         |
| --------------- | ------------------------------------------------------------ |
| s[n]            | 切片s中索引位置为n的项                                       |
| s[:]            | 从切片s的索引位置0到len(s)-1处所获得的切片                   |
| s[low:]         | 从切片s的索引位置low到len(s)-1处获得的切片                   |
| s[:high]        | 从切片s的索引位置0到high处获得的切片                         |
| s[low:high]     | 从切片s的索引位置low到high处获得的切片                       |
| s[low:high:max] | 从切片s的索引位置low到high处获得的切片, len=high-low, cap=max-low |
| len(s)          | 切片s的长度，总是<=cap(s)                                    |
| cap(s)          | 切片s的容量，总是>=len(s)                                    |

#### 通过make来创建切片

```go
var slice []type = make([]type, len)
slice := make([]type, len)
slice := make([]type, len, cap)
```

![切片](E:\笔记\Go\assets\1.jpg)

切片的内存布局

![切片](E:\笔记\Go\assets\2.jpg)

读写操作实际目的是底层数组，只需要注意索引符号的差别

```go
package main

import (
    "fmt"
)

func main() {
    data := [...]int{0, 1, 2, 3, 4, 5}

    s := data[2:4]
    s[0] += 100
    s[1] += 200

    fmt.Println(s) // [102 203]
    fmt.Println(data) // [0 1 102 203 4 5]
}
```

#### append内置函数操作切片(切片追加)

```go
package main

import "fmt"

func main() {
    var a = []int{1, 2, 3}
    fmt.Printf("slice a: %v\n", a)
    var b = []int{4, 5, 6}
    fmt.Printf("slice b: %v\n", b)
    c := append(a, b...)
    fmt.Printf("slice c: %v\n", c)
    d := append(c, 7)
    fmt.Printf("slice d: %v\n", d)
    e := append(d, 8, 9, 10)
    fmt.Printf("slice e: %v\n", e)
}
```

append：向slice尾部添加数据，返回**新的**slice对象

#### 超出原slice.cap限制，就会重新分配底层数组，即便原数组并未填满。

```go
package main

import "fmt"

func main() {
    data := [...]int{0, 1, 2, 3, 4, 5, 10:0}
    s := data[:2:3]
    
    s = append(s, 100, 200) // 一次append两个值， 超出s.cap限制
    
    fmt.Println(s, data) // 重新分配底层数组，与原数组无关
    fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。
}
```

从输出结果可以看出，append 后的 s 重新分配了底层数组，并复制数据。如果只追加一个值，则不会超过 s.cap 限制，也就不会重新分配。 通常以 2 倍容量重新分配底层数组。在大批量添加数据时，建议一次性分配足够大的空间，以减少内存分配和数据复制开销。或初始化足够长的 len 属性，改用索引号进行操作。及时释放不再使用的 slice 对象，避免持有过期数组，造成 GC 无法回收。

#### 字符串和切片

string底层就是一个byte的数组，因此，也可以进行切片操作

```go
package main
import "fmt"
func main() {
    str := "hello world"
    s1 ；= str[0:5]
    fmt.Println(s1)
    s2 := str[6:]
    fmt.Println(s2)
}
```

string本身是不可变的，因此需要改变string中字符串，需要如下操作：

```go
package main

import (
    "fmt"
)

func main() {
    str := "Hello world"
    s := []byte(str) //中文字符需要用[]rune(str)
    s[6] = 'G'
    s = s[:8]
    s = append(s, '!')
    str = string(s)
    fmt.Println(str)
}
```

a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x

### Slice底层实现

切片常见的操作reslice，append，copy。与此同时，切片还具有可索引，可迭代的优秀特性

#### 切片和数组

在Go中，与C语言不同的是，C数组变量隐式作为指针使用不同，Go数组是指类型，赋值和函数传参操作都会复制整个数组数据

```go
package main

import "fmt"

func main() {
	var arr1 = [2]int{1, 2}
	var arr2 [2]int
	arr2 = arr1
	fmt.Printf("arr1: %p %v \n", &arr1, arr1)
	fmt.Printf("arr2: %p %v \n", &arr2, arr2)
}
arr1: 0xc0000140b0 [1 2] 
arr2: 0xc0000140c0 [1 2] 
```

可以看出在每次调用的时候都会产生不同的地址，也就是说Go中数组在赋值和传参的时候都是复制的。那么这样会损失大量的内存。于是我们可以通过传递数组的指针。

把第一个大数组传递给函数会消耗很多内存，采用切片的方式传参可以避免上述问题。切片是引用传递，所以它们不需要使用额外的内存并且比使用数组更有效率，并非所有时候都适合用切片代替数组，因为切片底层数组可能会在堆上分配内存，而且小数组在栈上拷贝的消耗也未必比 make 消耗大。

#### 切片的数据结构

切片的本身是一种只读的对象，其工作机制类似数组指针的一种封装
切片(slice)是对数组的一个连续片段的引用，所以切片是一个引用类型(更类似于C家族中的数组类型，Python的list)。这个片段可以使整个数组，或者是由起始和终止索引表示的一些项的子集。需要注意的是，终止索引标识的项不包括在切片内。切片提供一个指向数组的动态窗口。
给定项的切片索引可能比相关数组的相同元素的索引小。和数组不同的是，切片的长度可以在运行时修改，最小为0最大为相关数组的长度：切片是一个长度可变的数组

```go
type slice struct {
    array unsafe.Pointer
    len int
    cap int
}
```

三个部分组成：
Pointer是指向一个数组的指针，
len代表当前切片的长度
cap是当前切片的容量，cap总是大于等于len的

![img](E:\笔记\Go\assets\slice-3.png)

[详情链接](https://www.jianshu.com/p/030aba2bff41)



### Go指针

`&`取地址和`*`根据地址取值

变量，指针地址，指针变量，取地址，取值的互相关系和特性如下：

```
1. 对变量进行取地址(&)操作，可以获得这个变量的指针变量。
2. 指针变量的值是指针地址。
3. 对指针变量进行取值(*)操作，可以获得指针变量指向的原来变量的值。
```

#### new和make的区别

```
1. 二者都是用来做内存分配的。
2. make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身。
3. 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
```

#### 指针小练习

程序定义一个int变量num的地址并打印
将num的地址付给指针ptr，并通过ptr取修改num的值

```go
package main

import "fmt"

func ptr(num *int) {
    *num = 100
}
func main() {
    var num int
    num = 10
    fmt.Printf("num: %p\n", &num)
    ptr(&num)
    fmt.Println(num)
}
```

### Map

#### map的创建

```go
scoreMap := make(map[T]T, len, cap)
scoreMap := map[T]T{....}
```

#### 判断map是否存在

```go
value, ok := map[key]
func main() {
    scoreMap := make(map[string]int)
    scoreMap["张三"] := 90
    scoreMap["小明"] := 100
    // 如果key存在ok为true，v为对应的值：不存在ok为false，v为值类型的零值
    v, ok := scoreMap["张三"]
    if ok {
        fmt.Println(v)
    } else {
        fmt.Println("查无此人")
    }
}
```

#### map的遍历

```go
package main

import "fmt"

func main() {
	// map的遍历
	scoreMap := map[string]int{
		"张三": 100,
		"王五": 110,
	}
	for k, v := range scoreMap {
		fmt.Println("key:", k, "value", v)
	}
	for k := range scoreMap {
		fmt.Println(k)
	}
}
```

#### 使用delete()函数删除键值对

```go
// delete(map, key)
// map: 表示要删除键值对的map
// key: 表示要删除键值对的键
func main() {
    scoreMap := make(map[string]int)
    scoreMap["张三"] = 90
    scoreMap["小明"] = 100
    scoreMap["王五"] = 60
    delete(scoreMap, "小明") // 将小明:100 这对键值对从scoreMap中删除
    for k, v := range scoreMap{
        fmt.Println(k, v)
    }
}
```

#### 按照一定的顺序遍历map

```go
func main() {
    rand.Seed(time.now().UnixNano()) // 初始化随机数种子
    var scoreMap = make(map[string]int, 200)
    for i := 0; i < 100; i++ {
        key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
        value := rand.Intn(100) // 生成0~99的随机数
        scoreMap[key] = value
    }
    // 提取出map中所有的key存入切片keys
    var keys := make([]string, 200)
    for k := range scoreMap {
        keys = append(keys, k)
    }
    // 对切片进行排序
    sort.Strings(keys)
    // 按照排序后的keys遍历map
    for _, key := range keys {
        fmt.Println(key, scoreMap[key])
    }
}
```

[深入Go的Map使用和实现原理](https://cloud.tencent.com/developer/article/1468799)



