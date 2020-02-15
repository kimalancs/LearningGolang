# Learn Golang

## 学习资料

* 《Go语言入门经典》
* 《Go in Action》
* 《Go语言高级编程》

## 方向

1. 区块链（币圈、链圈）
2. 开发运维平台
3. 云平台
4. web开发

## web开发框架

### beego

github.com/astaxie/beego

中文文档
坑比较多

### gin

github.com/gin-gonic/gin

轻量级
bug较少

## notes

### go module

#### Go依赖管理发展

1. 初期所有第三方库都放在GOPATH目录下，导致同一个库只能保存一个版本
2. 为解决不同的项目依赖与同一第三方库的不同版本，Go1.5开始引入vendor机制，将第三方库放到项目目录的vendor目录中，go工具链会优先使用vendor目录下的第三方包进行编译、测试等。godep是基于vendor模式实现第三方依赖管理的工具。
3. Go1.11推出了官方版本管理工具go module，并从Go1.13开始默认使用

#### GO111MODULE

设置GO111MODULE环境变量，开启或关闭模块支持。共有三个可选值：`on`、`off`、`auto`，默认值是`auto`  

* `off`，禁用模块支持，编译时从GOPATH和vendor目录中查找包  
* `on`，启用模块支持，编译时忽略GOPATH和vendor目录，只根据go.mod下载依赖  
* `auto`，当项目在$GOPATH/src外，且项目目录下有go.mod时，开启模块支持  

#### GOPROXY

设置GOPROXY，默认为`https://proxy.golang.org`，国内无法访问，可以改为`https://goproxy.cn`，修改方式如下：  

* `export GOPROXY=https://goproxy.cn`
* `go env -w GOPROXY=https://goproxy.cn`

#### go mod命令

命令 | 含义
:---:|:---:
go mod download | 下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录
go mod edit | 编辑go.mod文件
go mod graph | 打印模块依赖图
go mod init | 初始化当前文件夹，创建go.mod文件
go mod tidy | 增加缺少的module，删除无用的module
go mod vendor | 将依赖复制到vendor下
go mod verify | 校验依赖
go mod why | 解释为什么需要依赖

#### go.mod

go.mod文件记录了项目所有的依赖信息，其结构大致如下：

```
module github.com/Q1mi/studygo/blogger

go 1.12

require (
    github.com/DeanThompson/ginpprof v0.0.0-20190408063150-3be636683586
    github.com/gin-gonic/gin v1.4.0
    github.com/go-sql-driver/mysql v1.4.1
    github.com/jmoiron/sqlx v1.2.0
    github.com/satori/go.uuid v1.2.0
    google.golang.org/appengine v1.6.1 // indirect
)
```

1. module用来定义包名
2. require用来定义依赖包及版本
3. indirect表示间接引用

#### 依赖的版本

go mod支持语义化版本号，比如`go get foo@v1.2.3`，也可以跟git的分支或tag，比如`go get foo@master`，当然也可以跟git提交哈希，比如`go get foo@e3702bed2`。关于依赖的版本支持以下几种格式：

```
gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
gopkg.in/vmihailenco/msgpack.v2 v2.9.1
gopkg.in/yaml.v2 <=v2.2.1
github.com/tatsushid/go-fastping v0.0.0-20160109021039-d7bb493dee3e
latest
```

#### replace

在国内访问golang.org/x的各个包都需要翻墙，你可以在go.mod中使用replace替换成github上对应的库。

```
replace (
    golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
    golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
    golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
```

#### go get

在项目中执行go get命令可以下载依赖包，并且还可以指定下载的版本。

运行`go get -u`将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
运行`go get -u=patch`将会升级到最新的修订版本
运行`go get package@version`将会升级到指定的版本号version
如果下载所有依赖可以使用`go mod download`命令。

#### 整理依赖

我们在代码中删除依赖代码后，相关的依赖库并不会在go.mod文件中自动移除。这种情况下我们可以使用`go mod tidy`命令更新go.mod中的依赖关系。

#### go mod edit

格式化
因为我们可以手动修改go.mod文件，所以有些时候需要格式化该文件。Go提供了一下命令：  
`go mod edit -fmt`

添加依赖项
`go mod edit -require=golang.org/x/text`

移除依赖项
如果只是想修改go.mod文件中的内容，那么可以运行`go mod edit -droprequire=package path`，比如要在go.mod中移除golang.org/x/text包，可以使用如下命令：
`go mod edit -droprequire=golang.org/x/text`
关于go mod edit的更多用法可以通过go help mod edit查看。

#### 在项目中使用go module

既有项目
如果需要对一个已经存在的项目启用go module，可以按照以下步骤操作：

1. 在项目目录下执行go mod init，生成一个go.mod文件。
2. 执行go get，查找并记录当前项目的依赖，同时生成一个go.sum记录每个依赖库的版本和哈希值。

新项目
对于一个新创建的项目，我们可以在项目文件夹下按照以下步骤操作：

1. 执行go mod init 项目名命令，在当前项目文件夹下创建一个go.mod文件。
2. 手动编辑go.mod中的require依赖项或执行go get自动发现、维护依赖。

### 没有泛型

[为什么Go语言没有泛型](https://draveness.me/whys-the-design-go-generics?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io)

Go，语法元素少，设计简单的编程语言，而简单的设计意味着较弱的表达能力，需要使用更多时间编写重复的逻辑

```go
package sort

func Float64s(a []float64)
func Strings(a []string)
func Ints(a []int)
...
```

上述sort包中类似功能、底层实现逻辑也相同的函数，但由于传入类型的不同，需要重复实现

```java
public class ArraySortViaComparabale {
    public <E extends Comparable> void insert insertionSort(E[] a) {
        for (int i = 1; i < a.length; i++) {
            Comparable itemToInsert = a[i];
            int j =i;
            while (j != 0 && greaterThan(a[j-1], itemToInsert)) {
                a[j] = a[j-1]
                j = j - 1
            };
            a[j] = itemToInsert;
        }
    }
    private static boolean greaterThan(E left, Object right) {return left.compareTo(right) == 1;}
}
```

上面这段Java代码使用泛型数组作为参数实现了通用的数组排序逻辑，任意类型只要实现了Comparable接口，insertionSort函数就能排序由该对象组成的数组

使用泛型能够减少重复的代码和逻辑，提供更强的表达能力从而提升效率

#### Go的泛型困境

必须要在开发效率、编译速度和运行速度三者中平衡

* C语言没有泛型，牺牲开发效率。但不引入泛型也有好处，降低编译器的实现复杂度，保证源代码的编译速度

* C++使用编译期间类型特化实现泛型，提供了强大的抽象能力。虽然提高了开发效率，不需要手写同一逻辑的相似实现，但是编译器的实现变得非常复杂，泛型展开会生成大量重复代码，导致最终的二进制文件膨胀和编译缓慢，往往需要链接器来解决代码重复的问题

* Java在1.5版本引入了泛型，使用类型擦除实现泛型，泛型只在编译期间有效，正确检验泛型结果之后，类型擦除会删除泛型的相关信息，并且在对象进入和离开方法的边界处添加类型检查和类型转换的方法。泛型信息不会进入到运行时阶段。Java类型的装箱和拆箱会降低程序的执行效率

#### 不紧急不完善

Go语言团队认为加入泛型并不紧急，更重要的是完善运行时机制，包括调度器、垃圾收集器等功能。并且使用Go语言时没有太多对泛型的需求，只是在提供一些通用抽象逻辑时不得不使用interface{}空接口作为方法的参数
