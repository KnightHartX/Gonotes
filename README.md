# Go语言学习笔记
<!-- TOC -->

- [Go语言学习笔记](#go语言学习笔记)
    - [Go语言基础](#go语言基础)
        - [Go语言简介](#go语言简介)
        - [文件名、关键字与标识符](#文件名关键字与标识符)
            - [文件名](#文件名)
            - [关键字](#关键字)
        - [预定义标识符](#预定义标识符)
        - [Helloworld](#helloworld)
            - [运行结果](#运行结果)
        - [声明](#声明)
        - [变量](#变量)
            - [运行结果](#运行结果-1)
            - [简短变量声明](#简短变量声明)
            - [运行结果](#运行结果-2)
        - [赋值](#赋值)
            - [运行结果](#运行结果-3)
            - [元祖赋值](#元祖赋值)
            - [运行结果](#运行结果-4)
        - [基础数据类型](#基础数据类型)
            - [整形](#整形)
            - [浮点数](#浮点数)
            - [运行结果](#运行结果-5)
            - [复数](#复数)
            - [布尔型](#布尔型)
            - [运行结果](#运行结果-6)
            - [字符串](#字符串)
            - [运行结果](#运行结果-7)
            - [常量](#常量)
            - [运行结果](#运行结果-8)
        - [Go语言循环](#go语言循环)
            - [运行结果](#运行结果-9)
        - [Go语言条件语句](#go语言条件语句)
            - [if循环](#if循环)
            - [运行结果](#运行结果-10)
            - [switch语句](#switch语句)
            - [运行结果](#运行结果-11)
    - [函数、切片、map](#函数切片map)
        - [函数](#函数)
            - [函数声明](#函数声明)
            - [运行结果](#运行结果-12)
            - [多返回值](#多返回值)
            - [运行结果](#运行结果-13)
            - [函数值](#函数值)
            - [运行结果：](#运行结果)
        - [切片](#切片)
            - [切片的定义和初始化](#切片的定义和初始化)
            - [len() 和 cap() 函数](#len-和-cap-函数)
            - [切片截取](#切片截取)
            - [append() 和 copy() 函数](#append-和-copy-函数)
            - [切片示例代码：](#切片示例代码)
            - [运行结果](#运行结果-14)
        - [map](#map)
            - [map的创建](#map的创建)
            - [map的访问](#map的访问)
            - [map元素的删除](#map元素的删除)
            - [运行结果](#运行结果-15)
        - [面向对象编程](#面向对象编程)
            - [Go语言面向对象过程的实现](#go语言面向对象过程的实现)
            - [结构体的建立和实例的创建](#结构体的建立和实例的创建)
            - [方法](#方法)
            - [继承](#继承)
            - [方法的重写](#方法的重写)
            - [基于只针对象的方法](#基于只针对象的方法)
            - [运行结果](#运行结果-16)
        - [接口](#接口)
            - [运行结果](#运行结果-17)
    - [并发编程](#并发编程)
        - [Goruntimes](#goruntimes)
            - [运行结果](#运行结果-18)
        - [Channels](#channels)
        - [不带缓存的Channel](#不带缓存的channel)
        - [Channel串联](#channel串联)
            - [运行结果](#运行结果-19)
        - [带缓存的channel](#带缓存的channel)
            - [运行结果](#运行结果-20)
            - [运行结果：](#运行结果-1)
    - [网络编程](#网络编程)
        - [TCP编程](#tcp编程)
            - [返回本机\目标主机的网络地址](#返回本机\目标主机的网络地址)
            - [运行结果](#运行结果-21)
            - [运行结果](#运行结果-22)
            - [解析带端口号地址的ip](#解析带端口号地址的ip)
            - [运行结果](#运行结果-23)
        - [UDP编程](#udp编程)
            - [什么是UDP编程](#什么是udp编程)
            - [net包常用的UDP库函数](#net包常用的udp库函数)
    - [其他](#其他)
        - [Go语言范围(Range)](#go语言范围range)
            - [运行结果](#运行结果-24)
        - [Go嵌入C语言代码](#go嵌入c语言代码)
            - [运行结果](#运行结果-25)
        - [反射](#反射)
        - [unsafe](#unsafe)
            - [运行结果](#运行结果-26)

<!-- /TOC -->
## Go语言基础
### Go语言简介
Go是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。
罗伯特·格瑞史莫，罗勃·派克（Rob Pike）及肯·汤普逊于2007年9月开始设计Go语言，稍后Ian Lance Taylor、Russ Cox加入专案。Go语言是基于Inferno作业系统所开发的。Go语言于2009年11月正式宣布推出，成为开放原始码专案，并在Linux及Mac OS X平台上进行了实现，后来追加了Windows系统下的实现。
目前Go语言每半年发布一个二级版本（即升级1.x到1.y）
Go语言的语法接近C语言，但对于变量的声明有所不同。Go语言支持垃圾回收功能。Go语言的并行模型是以东尼·霍尔的交谈循序程式（CSP）为基础，采取类似模型的其他语言包括Occam和Limbo，但它也具有Pi运算的特征，比如通道传输。在1.8版本中开放插件（Plugin）的支持，这意味著现在能从Go语言中动态载入部分函式。
与C++相比，Go语言并不包括如异常处理、继承、泛型、断言、虚函数等功能，但增加了 Slice 型、并发、管道、垃圾回收、接口（Interface）等特性的语言级支持。Google 目前仍正在讨论是否应该支持泛型，其态度还是很开放的，但在该语言的常见问题列表中，对于断言的存在，则持负面态度，同时也为自己不提供型别继承来辩护。
不同于Java，Go语言内嵌了关联数组（也称为哈希表（Hashes）或字典（Dictionaries）），就像字符串类型一样。

### 文件名、关键字与标识符

#### 文件名

Go 的源文件以 .go 为后缀名存储在计算机中，这些文件名均由小写字母组成，如 helloworld.go 。如果文件名由多个部分组成，则使用下划线 _ 对它们进行分隔，如 hello_world.go 。文件名不包含空格或其他特殊字符。

#### 关键字

Go语言中25个关键字或保留字
![](http://omt37wai0.bkt.clouddn.com/17-7-22/85571792.jpg)

### 预定义标识符

除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符，其中包含了基本类型的名称和一些基本的内置函数
![](http://omt37wai0.bkt.clouddn.com/17-7-22/90101005.jpg)

### Helloworld

以一个helloworld开始
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

#### 运行结果

![](http://omt37wai0.bkt.clouddn.com/17-7-22/48077561.jpg)

Go语言支持Unicode，因此能处理所有语言，包括中文。
Go语言的代码通过包（package）组织，包类似于其它语言里的库（libraries）或者模块（modules）。一个包由位于单个目录下的一个或多个.go源代码文件组成, 目录定义包的作用。每个源文件都以一条package声明语句开始，这个例子里就是package main, 表示该文件属于哪个包，紧跟着一系列导入（import）的包，之后是存储在这个文件里的程序语句。
Go的标准库提供了100多个包，以支持常见功能，如输入、输出、排序以及文本处理。比如fmt包，就含有格式化输出、接收输入的函数。Println是其中一个基础函数，可以打印以空格间隔的一个或多个值，并在最后添加一个换行符，从而输出一整行。
main包比较特殊。它定义了一个独立可执行的程序，而不是一个库。在main里的main 函数 也很特殊，它是整个程序执行时的入口7。main函数所做的事情就是程序做的。当然了，main函数一般调用其它包里的函数完成很多工作, 比如fmt.Println。
必须告诉编译器源文件需要哪些包，这就是import声明以及随后的package声明扮演的角色。hello world例子只用到了一个包，大多数程序需要导入多个包。
必须恰当导入需要的包，缺少了必要的包或者导入了不需要的包，程序都无法编译通过。这项严格要求避免了程序开发过程中引入未使用的包。

### 声明

Go语言主要有四种类型的声明语句：**var**、**const**、**type**和**func**，分别对应变量、常量、类型和函数实体对象的声明。

### 变量

var声明语句可以创建一个特定类型的变量，然后给变量附加一个名字，并且设置变量的初始值。变量声明的一般语法如下：
>var 变量名字 类型 = 表达式

其中“类型”或“= 表达式”两个部分可以省略其中的一个。
如果省略的是类型信息，那么将根据初始化表达式来推导变量的类型信息。
如果初始化表达式被省略，那么将用零值初始化该变量。 
数值类型变量对应的零值是0，布尔类型变量对应的零值是false，字符串类型对应的零值是空字符串，接口或引用类型（包括slice、map、chan和函数）变量对应的零值是nil。数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值。
因此，Go语言中不存在未初始化的量，避免了C语言中诸如“野码”的存在。
也可以在一个声明语句中同时声明一组变量，或用一组初始化表达式声明并初始化一组变量。如果省略每个变量的类型，将可以声明多个类型不同的变量（类型由初始化表达式推导）：
```go
package main

import "fmt"

func main() {

    var a, b, c int
    var d, e, f=6, 7.1, "ff"
    fmt.Println(a,b,c,d,e,f)
}
```

#### 运行结果

![](http://omt37wai0.bkt.clouddn.com/17-7-22/24085232.jpg)

#### 简短变量声明

在Go语言中，有一种称为简短变量声明语句的形式可用于声明和初始化局部变量。它以“名字 := 表达式”形式声明变量，变量的类型根据表达式来自动推导。
例如：
```go
package main

import "fmt"

func main()  {
    a:="AAAA"
    b:=6
    c:=7.000000001
    fmt.Println(a,b,c)
}
```

#### 运行结果

![](http://omt37wai0.bkt.clouddn.com/17-7-22/9977048.jpg)

请记住“:=”是一个变量声明语句，而“=‘是一个变量赋值操作。

### 赋值

使用赋值语句可以更新一个变量的值，最简单的赋值语句是将要被赋值的变量放在=的左边，新值的表达式放在=的右边。
也可使用类似于a++的方式进行赋值，例如：
```go
package main

import "fmt"

func main()  {
    a:=1
    a=100
    fmt.Println(a)
    a++
    fmt.Println(a)
}
```

#### 运行结果

![](http://omt37wai0.bkt.clouddn.com/17-7-22/4247672.jpg)

#### 元祖赋值

元组赋值是另一种形式的赋值语句，它允许同时更新多个变量的值。在赋值之前，赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值。
例如，利用元祖赋值计算两整数的最大公约数：
```go
package main

import "fmt"

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x

}

func main()  {
    var x,y int
    fmt.Println("请输入两个整数")
    fmt.Scanln(&x,&y)
    fmt.Printf("%v和%v的最大公约数为%v",x,y,gcd(x,y))
}
```

#### 运行结果

![](http://omt37wai0.bkt.clouddn.com/17-7-23/13939951.jpg)

### 基础数据类型

Go语言将数据类型分为四类：基础类型、复合类型、引用类型和接口类型。

#### 整形

Go语言有int8、int16、int32和int64四种截然不同大小的有符号整形数类型，分别对应8、16、32、64bit大小的有符号整形数，与此对应的是uint8、uint16、uint32和uint64四种无符号整形数类型。

Unicode字符rune类型是和int32等价的类型，通常用于表示一个Unicode码点。这两个名称可以互换使用。同样byte也是uint8类型的等价类型，byte类型一般用于强调数值是一个原始的数据而不是一个小的整数。

#### 浮点数

Go语言提供了两种精度的浮点数，float32和float64。它们的算术规范由IEEE754浮点数国际标准定义，该浮点数规范被所有现代的CPU支持。

常量math.MaxFloat32表示float32能表示的最大数值，大约是 3.4e38；对应的math.MaxFloat64常量大约是1.8e308。它们分别能表示的最小值近似为1.4e-45和4.9e-324。

一个float32类型的浮点数可以提供大约6个十进制数的精度，而float64则可以提供约15个十进制数的精度；通常应该优先使用float64类型，因为float32类型的累计计算误差很容易扩散，并且float32能精确表示的正整数并不是很大，因为float32的有效bit位只有23个，其它的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差。
```go
package main

import "fmt"
import "math"

func main()  {
    fmt.Println(math.MaxFloat32)
    fmt.Println(math.MaxFloat64)
}
```

#### 运行结果

![](http://omt37wai0.bkt.clouddn.com/17-7-23/2857398.jpg)

#### 复数

Go语言提供了两种精度的复数类型：complex64和complex128，分别对应float32和float64两种浮点数精度。内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部，如下例：
```go
package main

import "fmt"

func main()  {
    var a complex128=complex(145,64)
    var b complex128=complex(12,7)
    fmt.Println(a,b,real(a),real(b),imag(a),imag(b))
    fmt.Println(a*b,real(a*b),imag(a*b))
}
```
运行结果

![](http://omt37wai0.bkt.clouddn.com/17-7-23/78208166.jpg)

#### 布尔型

一个布尔类型的值只有两种：true和false。if和for语句的条件部分都是布尔类型的值，并且==和<等比较操作也会产生布尔型的值。一元操作符!对应逻辑非操作，因此!true的值为false。

布尔值可以和&&（AND）和||（OR）操作符结合

布尔值并不会隐式转换为数字值0或1，反之亦然。必须使用一个显式的if语句辅助转换。通常使用一个转换函数。

示例代码如下：
```go
package main

import "fmt"

func main()  {
    var a bool=true
    var b bool=false
    c:=a||b
    d:=a&&b
    fmt.Println(a,b,c,d,convert(c),convert(d))
}

func convert(a bool) int {
    var i int
    if a{
        i=1
    }else {
        i=0
    }
    return i
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-23/30729498.jpg)

#### 字符串

一个字符串是一个不可改变的字节序列，文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列。

内置的len函数可以返回一个字符串中的字节数目，引操作s[i]返回第i个字节的字节值，i必须满足0 ≤ i< len(s)条件约束。如果试图访问超出字符串索引范围的字节将会导致panic异常。

**注意**：第i个字节并不一定是字符串的第i个字符，因为对于非ASCII字符的UTF8编码会要两个或多个字节。

子字符串操作s[i:j]基于原始的s字符串的第i个字节开始到第j个字节（并不包含j本身）生成一个新字符串。生成的新字符串将包含j-i个字节。同样，如果索引超出字符串范围或者j小于i的话将导致panic异常。不管i还是j都可能被忽略，当它们被忽略时将采用0作为开始位置，采用len(s)作为结束的位置。

+操作符将两个字符串链接构造一个新字符串，==和<进行比较；比较通过逐个字节比较完成的，因此比较的结果是字符串自然编码的顺序。

字符串的值是不可变的：一个字符串包含的字节序列永远不会被改变，当然我们也可以给一个字符串变量分配一个新字符串值，但这并不会导致原始的字符串值被改变，因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的。

示例代码如下：
```go
package main

import "fmt"

func main()  {
    var s string="helloworld"
    fmt.Println(s)
    lenth:=len(s)
    fmt.Println(lenth)
    fmt.Println(string(s[1]))
    for i:=0;i<len(s);i++ {
        fmt.Println(string(s[i]))
    }
    s1:=s[1:7]
    s2:=s[0:6]
    println(s1+s2)
}
```

#### 运行结果

![](http://omt37wai0.bkt.clouddn.com/17-7-23/28770000.jpg)

#### 常量

常量表达式的值在编译期计算，而不是在运行期。每种常量的潜在类型都是基础类型：boolean、string或数字。

一个常量的声明语句定义了常量的名字，和变量的声明语法类似，常量的值不可修改，这样可以防止在运行期被意外或恶意的修改。例如，常量比变量更适合用于表达像π之类的数学常数，因为它们的值不会发生变化。

例子：
```go
package main

import "fmt"

func main()  {
    const (
        Pi  =3.1415926
        r=2
    )
    S:=Pi*r*r
    fmt.Println(S)
}
```

#### 运行结果

![](http://omt37wai0.bkt.clouddn.com/17-7-23/24510164.jpg)

### Go语言循环

![](http://omt37wai0.bkt.clouddn.com/17-7-23/93209796.jpg)

Go语言提供了三种for循环方式：

- 和 C 语言的 for 一样：

>for init; condition; post { }

- 和 C 的 while 一样：
>for condition { }

- 和 C 的 for(;;) 一样：
>for { }

其中：
- init： 一般为赋值表达式，给控制变量赋初值；
- condition： 关系表达式或逻辑表达式，循环控制条件；
- post： 一般为赋值表达式，给控制变量增量或减量。

![](http://omt37wai0.bkt.clouddn.com/17-7-23/1255769.jpg)

示例代码：
```go
package main

import "fmt"

func main()  {
    S:=0
    i:=0
    n:=100
    for i:=0;i<=n ;i++  {
        S=S+i
    }
    println(S)
    i=0
    for i<=n{
        S=S+i
        i++
    }
    fmt.Println(S)
}
```

#### 运行结果

![](http://omt37wai0.bkt.clouddn.com/17-7-23/78275968.jpg)

### Go语言条件语句

条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。

![](http://omt37wai0.bkt.clouddn.com/17-7-24/20573835.jpg)

#### if循环
![](http://omt37wai0.bkt.clouddn.com/17-7-24/67514913.jpg)

If 在布尔表达式为 true 时，其后紧跟的语句块执行，如果为 false 则执行 else 语句块。

代码示例：

```go
package main

import "fmt"

func main()  {
    var a bool=true
    var b bool = false
    if  a&&b==true{
        fmt.Println(1)
    }else  {
        fmt.Println(0)
    }
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-24/84379616.jpg)

#### switch语句
![](http://omt37wai0.bkt.clouddn.com/17-7-24/50783758.jpg)

switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。。
switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也**不需要**再加break。

示例代码：
```go
package main

import "fmt"

func main()  {
    fmt.Println("请输入一个正整数n")
    var n int
    fmt.Scanln(&n)
    m:=n%2
    switch m {
    case 0:fmt.Printf("%v为偶数\n",n)
    case 1:fmt.Printf("%v为奇数\n",n)
    }
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-24/56275469.jpg)

## 函数、切片、map
### 函数
#### 函数声明
函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。
```go
func name(parameter-list) (result-list) {
    body
}
```
形式参数列表描述了函数的参数名以及参数类型。这些参数作为局部变量，其值由参数调用者提供。返回值列表描述了函数返回值的变量名以及类型。如果函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的。如果一个函数声明不包括返回值列表，那么函数体执行完毕后，不会返回任何值。

函数的类型被称为函数的标识符。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型和标识符。形参和返回值的变量名不影响函数标识符也不影响它们是否可以以省略参数类型的形式表示。

每一次函数调用都必须按照声明顺序为所有参数提供实参（参数值）。在函数调用时，Go语言没有默认参数值，也没有任何方法可以通过参数名指定形参，因此形参和返回值的变量名对于函数调用者而言没有意义。

在函数体中，函数的形参作为局部变量，被初始化为调用者提供的值。函数的形参和有名返回值作为函数最外层的局部变量，被存储在相同的词法块中。

示例代码如下：
```go
package main

import "fmt"

func sqare(a float64) float64 {
    const pi  =3.1415926
    S:=pi*a*a
    return S
}

func main()  {
    var r float64
    fmt.Println("请输入半径r")
    fmt.Scan(&r)
    fmt.Println(sqare(r))
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-24/53038620.jpg)
#### 多返回值
在Go中，一个函数可以返回多个值。这可以用来确定函数是否得到了正确执行。
示例代码如下：
```go
package main

import "fmt"

func sqare(a float64) (S float64,ok bool) {
    const pi  =3.1415926
    S=pi*a*a
    ok=true
    return S,ok
}

func main()  {
    var r float64
    fmt.Println("请输入半径r")
    fmt.Scan(&r)
    fmt.Println(sqare(r))
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-24/99085590.jpg)
#### 函数值

在Go中，函数被看作第一类值（first-class values）：函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。对函数值（function value）的调用类似函数调用。

示例代码：
```go
package main

import "fmt"

func sqare(a float64) float64 {
    const pi  =3.1415926
    S:=pi*a*a
    return S
}

func main()  {
    var r float64
    fmt.Println("请输入半径r")
    fmt.Scan(&r)
    a:=sqare(r)
    b:=sqare(sqare(r))
    fmt.Println(a)
    fmt.Println(b)
}
```
#### 运行结果：
![](http://omt37wai0.bkt.clouddn.com/17-7-25/18237426.jpg)

### 切片

数组的长度不可改变，在特定场景中这样的集合不适用了，Go中提供了一种灵活，功能强悍的内置类型Slices切片,与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

切片中有两个概念：一是len长度，二是cap容量，长度是指已经被赋过值的最大下标+1，可通过内置函数len()获得。

容量是指切片目前可容纳的最多元素个数，可通过内置函数cap()获得。

切片是引用类型，因此在当传递切片时将引用同一指针，修改值将会影响其他的对象。

#### 切片的定义和初始化
声明一个未指定大小的数组来定义切片：
>var a []type 

切片不需要说明长度。
或使用make()函数来创建切片:
>a := make([]type, len)

初始化切片：

>a :=[] int {1,2,3 }

>a1 := a[startIndex:endIndex] 

或使用make函数定义并初始化：

>s :=make([]int,len,cap) 
#### len() 和 cap() 函数
切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

#### 切片截取
可以通过设置下限及上限来设置截取切片
#### append() 和 copy() 函数
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

使用append()函数则可以向切片中增加新的元素，而copy()可以获得一段切片的片段到新的切片。

#### 切片示例代码：
```go
package main

import "fmt"

func main()  {
    a:=make([]int,5,5)
    a=[]int{1,2,3,4,5}
    b:=make([]int,5,5)
    copy(b,a)
    fmt.Println(len(a))
    fmt.Println(cap(a))
    fmt.Println(a)
    fmt.Println(b)
    a=append(a,6)
    fmt.Println(len(a))
    fmt.Println(cap(a))
    fmt.Println(a)
    a1:=a[0:6]
    fmt.Println(a1)
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-25/75161598.jpg)

###    map
哈希表是一种巧妙并且实用的数据结构。它是一个无序的key/value对的集合，其中所有的key都是不同的，然后通过给定的key可以在常数时间复杂度内检索、更新或删除对应的value。

在Go语言中，一个map就是一个哈希表的引用，map类型可以写为map[K]V，其中K和V分别对应key和value。map中所有的key都有相同的类型，所有的value也有着相同的类型，但是key和value之间可以是不同的数据类型。其中K对应的key必须是支持==比较运算符的数据类型，所以map可以通过测试key是否相等来判断是否已经存在。虽然浮点数类型也是支持相等运算符比较的，但是不建议用浮点数做key，最坏的情况是可能出现的NaN和任何浮点数都不相等。对于V对应的value数据类型则没有任何的限制。

#### map的创建

>a:= make(map[string]int) // mapping from strings to ints

也可以用map字面值的语法创建map，同时还可以指定一些最初的key/value：
>ages := map[string]int{

>    "china": 1,

>  "us": 2,

>}

#### map的访问

Map中的元素通过key对应的下标语法访问：
>a["china"] = 1

>fmt.Println(a["china"]) // "1"

#### map元素的删除

使用内置的delete函数可以删除元素：

>delete(a, "china") // remove element a["china"]

所有这些操作是安全的，即使这些元素不在map中也没有关系；如果一个查找失败将返回value类型对应的零值。

map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作.

要想遍历map中全部的key/value对的话，可以使用range风格的for循环实现。

代码示例：
```go
package main

import "fmt"

func main()  {
    a:=make(map[string]int)
    a["china"]=1
    a["us"]=2
    fmt.Println(a["us"])
    delete(a,"us")
    a["jp"]=3
    for k, v := range a {
        fmt.Printf("%v is no.%v\n", k, v)
    }
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-25/32563292.jpg)

### 面向对象编程
#### Go语言面向对象过程的实现
面向对象主要包括了三个基本特征：封装、继承和多态。
封装，就是指运行的数据和函数绑定在一起，C++中主要是通过this指针来完成的。

继承，就是指class之间可以相互继承属性和函数。

多态，主要就是用统一的接口来处理通用的逻辑，每个class只需要按照接口实现自己的回调函数就可以了。

Go语言中并没有像C++，Java语言中这类的Class，它只含有像C语言中的结构体，用结构体和指针等特性，完成一个类的作用，很巧妙的使用了指针和结构体，不仅是go的面向对象，包括go语言中的map等操作都是借助了结构体。

C++、Java等面向对象的语言中，类的底层实现就是结构体，对象的引用就是指针，只是语言把他们封装起来了而已。

#### 结构体的建立和实例的创建
```go
type circle struct
{
    r float64
    h float64
}
```
如此，定义了一个圆柱体结构体，分别有半径和高两个属性。
可由以下几种方法来创建实例。

```go
circle1:=new(circle)
circle1:=&circle{}
circle1:=&circle{1,2}
circle1:=&circle{r:3,h:4}
```

#### 方法
在函数声明时，在其名字之前放上一个变量，即是一个方法。这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。

例：
```go
func S(p circle) float64 {
    pi:=3.1415926
    return pi*p.r*p.r*p.h
}
```
上面的代码里那个附加的参数p，叫做方法的接收器(receiver)，早期的面向对象语言留下的遗产将调用一个方法称为“向一个对象发送消息”。

在Go语言中，我们并不会像其它语言那样用this或者self作为接收器；我们可以任意的选择接收器的名字。由于接收器的名字经常会被使用到，所以保持其在方法间传递时的一致性和简短性是不错的主意。这里的建议是可以使用其类型的第一个字母，比如这里使用了Point的首字母p。

在方法调用过程中，接收器参数一般会在方法名之前出现。这和方法声明是一样的，都是接收器参数在方法名字之前。

#### 继承
继承：

当一个类型B的某个字段(匿名字段)的类型是另一个类型 A的时候，那么类型 A所拥有的全部字段都被隐式地引入了当前定义的这个类型B。这样就实现了继承。B类型的变量就可以调用A的所有属性和方法。也就是说A继承了B。

则可得到继承circle的类mass，多了一个属性密度D。
```go
type mass struct {
    circle
    D float64
}
```
则mass继承了circle的全部属性和方法。

#### 方法的重写
如果一个类型B实现了作为其属性的类型A中的方法。那么这个类型B的值调用方法的时候调用的是自己类型B的方法，而不是属性类型A的方法。

不同的类型可以有相同的方法名，因此需要在定义的时候消除歧义。

#### 基于只针对象的方法

当调用一个函数时，会对其每一个参数值进行拷贝，如果一个函数需要更新一个变量，或者函数的其中一个参数实在太大我们希望能够避免进行这种默认的拷贝，这种情况下我们就需要用到指针了。对应到我们这里用来更新接收器的对象的方法，当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法。

例：
```go
func SS(p *circle) float64 {
    pi:=3.1415926
    return pi*p.r*p.r*p.h
}
```
以上所有知识点综合后得到的程序：
```go
package main

import "fmt"

type circle struct
{
    r float64
    h float64
}

type mass struct {
    circle
    D float64
}

func S(p circle) float64 {
    pi:=3.1415926
    return pi*p.r*p.r*p.h
}

func M(p mass) float64 {
    pi:=3.1415926
    return p.h*p.r*p.r*pi*p.D
}

func SS(p *circle) float64 {
    pi:=3.1415926
    return pi*p.r*p.r*p.h
}

func main()  {
    circle1:=circle{2.4,3.1}
    fmt.Println(circle1)
    mass1:=mass{circle{2.4 ,3.1},3.23}
    fmt.Println(mass1)
    fmt.Println(S(circle1))
    fmt.Println(M(mass1))
    fmt.Println(SS(&circle1))
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-26/27428608.jpg)

### 接口
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

接口类型是对其它类型行为的抽象和概括；因为接口类型不会和特定的实现细节绑定在一起，通过这种抽象的方式我们可以让我们的函数更加灵活和更具有适应能力。

原则：谁使用，谁实现

形式：
```go
/* 定义接口 */
type interface_name interface {
  method_name1 [return_type]
  method_name2 [return_type]
  method_name3 [return_type]
  ...
  method_namen [return_type]
}

/* 定义结构体 */
type struct_name struct {
  /* variables */
}

/* 实现接口方法 */
func (struct_name_variable struct_name) method_name1() [return_type] {
  /* 方法实现 */
}
...
func (struct_name_variable struct_name) method_namen() [return_type] {
  /* 方法实现*/
}
```
代码示例：
```go
package main

import (
    "fmt"
)

type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

func main() {
    var phone Phone

    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()

}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-26/97316280.jpg)

## 并发编程
### Goruntimes
go中没有多进程多线程的概念，而是使用goroutine概念，我们可以把goroutine当作其他语言中的线程，当在一个函数前加入go关键字，就启动了一个goroutine。

与操作系统或者其它语言提供的线程相比较，可以简单地把goroutine类比作一个线程。

当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。

下面的这个例子说明了如何让主程序在运行斐波那契数列的同时，并发运行一个等待动画：
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go spinner(100 * time.Millisecond)
    const n = 45
    fibN := fib(n) // slow
    fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
    for {
        for _, r := range `-\|/` {
            fmt.Printf("\r%c", r)
            time.Sleep(delay)
        }
    }
}

func fib(x int) int {
    if x < 2 {
        return x
    }
    return fib(x-1) + fib(x-2)
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-26/50564423.jpg)

![](http://omt37wai0.bkt.clouddn.com/17-7-26/50564423.jpg)

![](http://omt37wai0.bkt.clouddn.com/17-7-26/79998804.jpg)

后主函数返回。主函数返回时，所有的goroutine都会被直接打断，程序退出。除了从主函数退出或者直接终止程序之外，没有其它的编程方法能够让一个goroutine来打断另一个的执行。

或者，过goroutine之间的通信来让一个goroutine请求其它的goroutine，并被请求的goroutine自行结束执行。

### Channels
如果说goroutine是Go语音程序的并发体的话，那么channels它们之间的通信机制。一个channels是一个通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息。每个channel都有一个特殊的类型，也就是channels可发送数据的类型。一个可以发送int类型数据的channel一般写为chan int。
创建：
```go
ch := make(chan int) // ch has type 'chan int'
```
一个channel有发送和接受两个主要操作，都是通信行为。一个发送语句将一个值从一个goroutine通过channel发送到另一个执行接收操作的goroutine。发送和接收两个操作都是用<-运算符。在发送语句中，<-运算符分割channel和要发送的值。在接收语句中，<-运算符写在channel对象之前。一个不使用接收结果的接收操作也是合法的。

Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。对一个已经被close过的channel之行接收操作依然可以接受到之前已经成功发送的数据；如果channel中已经没有数据的话讲产生一个零值的数据。

>close(ch)

### 不带缓存的Channel
一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作，当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句。反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channels上执行发送操作。

基于无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作。因为这个原因，无缓存Channels有时候也被称为同步Channels。当通过一个无缓存Channels发送数据时，接收者收到数据发生在唤醒发送者goroutine之前。

### Channel串联
Channels也可以用于将多个goroutine链接在一起，一个Channels的输出作为下一个Channels的输入。这种串联的Channels就是所谓的管道（pipeline）。

程序示例：
```go
package main

import "fmt"

func main() {
    naturals := make(chan int)
    squares := make(chan int)

    // Counter
    go func() {
        for x := 0; ; x++ {
            naturals <- x
        }
    }()

    // Squarer
    go func() {
        for {
            x := <-naturals
            squares <- x * x
        }
    }()

    // Printer (in main goroutine)
    for {
        fmt.Println(<-squares)
    }
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-26/39565201.jpg)

### 带缓存的channel
带缓存的Channel内部持有一个元素队列。队列的最大容量是在调用make函数创建channel时通过第二个参数指定的。

向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列空间。相反，如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。

cap()函数可以获得缓存空间大大小

len()函数可以获得缓存元素的个数

示例程序：
```go
package main

import "fmt"

func main()  {
    responses := make(chan int, 9)
    go func() {responses<-jisuan(3)}()
    go func() {responses<-jisuan(4)}()
    go func() {responses<-jisuan(5)}()
    go func() {responses<-jisuan(6)}()
    go func() {responses<-jisuan(7)}()
    go func() {responses<-jisuan(8)}()
    go func() {responses<-jisuan(9)}()
    go func() {responses<-jisuan(10)}()
    go func() {responses<-jisuan(11)}()
    a:=<-responses
    fmt.Println(a)
}

func jisuan(a int) int {
    s:=a*a
    return s
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-26/21588824.jpg)

实验：并发计算pi
代码：
```go
package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Println("请输入n值")
    fmt.Scanln(&n)
    fmt.Println(pi(n))
}

func pi(n int) float64 {
    ch:=make(chan float64,100)
    for k:=0; k<=n; k++ {
        go term(ch,float64(k))
    }
    var f float64=0
    for k:=0;k<=n;k++ {
        f+=<-ch
    }
    return f
}

func term(ch chan float64, k float64) {
    ch<-4*math.Pow(-1, k) / (2*k+1)
}
```
#### 运行结果：
![](http://omt37wai0.bkt.clouddn.com/17-7-26/73739209.jpg)

## 网络编程
### TCP编程
输控制协议（英语：Transmission Control Protocol，缩写为 TCP）是一种面向连接的、可靠的、基于字节流的传输层通信协议，由IETF的RFC 793定义。在简化的计算机网络OSI模型中，它完成第四层传输层所指定的功能，用户数据报协议（UDP）是同一层内另一个重要的传输协议。

在因特网协议族（Internet protocol suite）中，TCP层是位于IP层之上，应用层之下的中间层。不同主机的应用层之间经常需要可靠的、像管道一样的连接，但是IP层不提供这样的流机制，而是提供不可靠的包交换。

应用层向TCP层发送用于网间传输的、用8位字节表示的数据流，然后TCP把数据流分割成适当长度的报文段（通常受该计算机连接的网络的数据链路层的最大传输单元（MTU）的限制）。之后TCP把结果包传给IP层，由它来通过网络将包传送给接收端实体的TCP层。TCP为了保证不发生丢包，就给每个包一个序号，同时序号也保证了传送到接收端实体的包的按序接收。然后接收端实体对已成功收到的包发回一个相应的确认（ACK）；如果发送端实体在合理的往返时延（RTT）内未收到确认，那么对应的数据包就被假设为已丢失将会被进行重传。TCP用一个校验和函数来检验数据是否有错误；在发送和接收时都要计算校验和。
#### 返回本机\目标主机的网络地址
本机地址：
```go
package main

import (
    "fmt"
    "net"
)

func main()  {
    address,error:=net.InterfaceAddrs()
    if error!=nil{
        fmt.Printf("出现错误，错误信息为%v",error)
    }
    fmt.Printf("本机地址为：%v",address)
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-26/76358011.jpg)

目标主机地址：
```go
package main

import (
    "fmt"
    "net"
)

func main() {
    address, error := net.LookupIP("www.google.com")
    if error != nil {
        fmt.Printf("出现错误，错误信息为%v", error)
    }
    fmt.Printf("目标主机地址为：%v", address)
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-26/11558708.jpg)
#### 解析带端口号地址的ip
```go
package main

import (
    "fmt"
    "net"
)

func main() {
    ip, error := net.ResolveTCPAddr("tcp","www.google.com:80")
    if error != nil {
        fmt.Printf("出现错误，错误信息为%v", error)
    }else {
        fmt.Printf("目标TCP地址为：%v", ip)
    }

```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-26/2670722.jpg)

### UDP编程
#### 什么是UDP编程
- 使用UDP协议(User Datagram Protocol)来传输数据的编程。
- UDP协议→是面向无连接的、不可靠的数据报投递服务。

- 当使用UDP 协议传输信息流时，用户应用程序必须负责解决数据报丢失、重复、排序，差错确认等问题。

- 资源消耗小处理速度快，通常音频、视频和普通数据在传送时使UDP较多。如QQ使用的就是UDP协议。

- UDP适用于一次只传少量数据的环境。

- 从理论上说，包含报头在内的数据报的最大长度为65535字
节。不过，一些实际应用往往会限制数据报的大小，有时会降低到8192字节。

#### net包常用的UDP库函数
- func ResolveUDPAddr(net, addr string) (*UDPAddr, error)
- 把addr 地址字符串，解析成UDPAddr 地址。
- net 可以是”udp”,”udp4”,”udp6”
- addr 是一个地址字符串，由主机名或IP 地址，以及“:”后面跟着的端口号组成。如果是IPv6，主机部分必须在方括内，如[::1]8080
----
- func ListenUDP(net string, laddr *UDPAddr) (*UDPConn, error)
- 在指定的地址(laddr)监听，等待UDP 数据包的到达。
- 返回*UDPConn，可以使用连接的ReadFrom函数来读取UDP 数据，用WriteTo 来向客户端发送数据。
- func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)
- 服务端用来读取UDP 数据。Addr 是发送方的地址。

服务端代码：
```go
package main

import
(
"fmt"
"net"
)

func HandleCilent(conn *net.UDPConn,data []byte,address *net.UDPAddr)  {
    fmt.Println("收到数据："+string(data))
    conn.WriteToUDP([]byte("数据已收到"),address)
}

func main() {
    //解析UDP地址
    address, error := net.ResolveUDPAddr("udp", ":7070")
    if error != nil {
        fmt.Printf("出现错误，错误信息为%v", error)
    } else {
        return
    }
    //7070端口监听
    conn, error := net.ListenUDP("udp", address)
    if error != nil {
        fmt.Printf("出现错误，错误信息为%v", error)
    } else {
        return
    }
    for{//循环接收数据，处理数据
        var buf [1024]byte
        n,address,error:=conn.ReadFromUDP(buf[0:])
        if error != nil {
            fmt.Printf("出现错误，错误信息为%v", error)
        } else {
            return
            //开启新线程处理客户端数据
            go HandleCilent(conn,buf[0:n],address)
    }
}
}
```

客户端代码：
```go
package main

import (
    "net"
    "fmt"
)

func main()  {
    //解析服务器UDP地址
    address, error := net.ResolveUDPAddr("udp", "127.0.0.1:7070")
    if error != nil {
        fmt.Printf("出现错误，错误信息为%v", error)
    } else {
        return
    }
    //连接服务器
    conn,error:=net.DialUDP("udp",nil,address)
    if error != nil {
        fmt.Printf("出现错误，错误信息为%v", error)
    } else {
        return
    }
    defer conn.Close()//关闭连接
    //向服务器发送数据
    conn.Write([]byte("Hello,Server"))
    var buf [1024]byte
    //读取服务器相应信息
    n,_,error:=conn.ReadFromUDP(buf[0:])
    if error != nil {
        fmt.Printf("出现错误，错误信息为%v", error)
    } else {
        return
    }
    fmt.Println(string(buf[0:n]))
}
```
## 其他
### Go语言范围(Range)
Go 语言中 range 关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。
示例代码：
```go
package main

import "fmt"

func main() {
    //这是我们使用range去求一个slice的和。使用数组跟这个很类似
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    //在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    //range也可以用在map的键值对上。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-26/32293184.jpg)

### Go嵌入C语言代码
想在Go代码中使用C语言必须在代码开头注释中写，然后再紧接着的下一行写import "C"，这样就算是导入完成了。这个”C”不是一个真正的包，而是一个类似于命名空间的东西，所有能调用的C的变量、函数都包含在里面。

代码示例：
```go
package main
/*
#include <stdio.h>
#include <stdlib.h>
void say_hello() {
        printf("Hello World!\n");
}
*/
import "C"
func main() {
    C.say_hello()
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-26/40659156.jpg)

### 反射
在计算机科学领域，反射是指一类应用，它们能够自描述和自控制。也就是说，这类应用通过采用某种机制来实现对自己行为的描述（self-representation）和监测（examination），并能根据自身行为的状态和结果，调整或修改应用所描述行为的状态和相关的语义。

### unsafe
Golang的unsafe包是一个很特殊的包。 为什么这样说呢？ 本文将详细解释。

来自go语言官方文档的警告

unsafe包的文档是这么说的：

>导入unsafe的软件包可能不可移植，并且不受Go 1兼容性指南的保护。

Go 1 兼容性指南这么说：

>导入unsafe软件包可能取决于Go实现的内部属性。 我们保留对可能导致程序崩溃的实现进行更改的权利。

当然包名称暗示unsafe包是不安全的。 但这个包有多危险呢？ 让我们先看看unsafe包的作用。

Unsafe包的作用:

unsafe包含以下资源：

三个函数：

func Alignof（variable ArbitraryType）uintptr

func Offsetof（selector ArbitraryType）uintptr

func Sizeof（variable ArbitraryType）uintptr

和一种类型：

类型Pointer * ArbitraryType

这里，ArbitraryType不是一个真正的类型，它只是一个占位符。

与Golang中的大多数函数不同，上述三个函数的调用将始终在编译时求值，而不是运行时。 这意味着它们的返回结果可以分配给常量。

unsafe包中的函数中非唯一调用将在编译时求值。当传递给len和cap的参数是一个数组值时，内置函数和cap函数的调用也可以在编译时被求值。

除了这三个函数和一个类型外，指针在unsafe包也为编译器服务。

出于安全原因，Golang不允许以下之间的直接转换：

两个不同指针类型的值，例如 int64和 float64。

指针类型和uintptr的值。

但是借助unsafe.Pointer，我们可以打破Go类型和内存安全性，并使上面的转换成为可能。这怎么可能发生？让我们阅读unsafe包文档中列出的规则：

任何类型的指针值都可以转换为unsafe.Pointer。

unsafe.Pointer可以转换为任何类型的指针值。

uintptr可以转换为unsafe.Pointer。

unsafe.Pointer可以转换为uintptr。

这些规则与Go规范一致：

底层类型uintptr的任何指针或值都可以转换为指针类型，反之亦然。

规则表明unsafe.Pointer类似于c语言中的void 。当然，void 在C语言里是危险的！

在上述规则下，对于两种不同类型T1和T2，可以使 T1值与unsafe.Pointer值一致，然后将unsafe.Pointer值转换为 T2值（或uintptr值）。通过这种方式可以绕过Go类型系统和内存安全性。当然，滥用这种方式是很危险的。

示例：

```go
package main

import (
    "fmt"
    "unsafe"
)
func main() {
    var n int64 = 5
    var pn = &n
    var pf = (*float64)(unsafe.Pointer(pn))
    // now, pn and pf are pointing at the same memory address
    fmt.Println(*pf) // 2.5e-323
    *pf = 3.14159
    fmt.Println(n) // 4614256650576692846
}
```
#### 运行结果
![](http://omt37wai0.bkt.clouddn.com/17-7-26/66703020.jpg)

在这个例子中的转换可能是无意义的，但它是安全和合法的

因此，资源在unsafe包中的作用是为Go编译器服务，unsafe.Pointer类型的作用是绕过Go类型系统和内存安全。
