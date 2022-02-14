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



