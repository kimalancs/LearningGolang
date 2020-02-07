package splitstrings

import "strings"

// TDD test driven development
// 单元测试
// go test命令是一个按照一定约定和组织的测试代码的驱动程序
// 在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中
// 在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。
// 测试函数	函数名前缀为Test	测试程序的一些逻辑行为是否正确
// 基准函数	函数名前缀为Benchmark	测试函数的性能
// 示例函数	函数名前缀为Example	为文档提供示例文档
// go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数
// 然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件

// 就像细胞是构成我们身体的基本单位，一个软件程序也是由很多单元组件构成的
// 单元组件可以是函数、结构体、方法和最终用户可能依赖的任意东西
// 总之我们需要确保这些组件是能够正常运行的
// 单元测试是一些利用各种方法测试单元组件的程序，它会将结果与预期输出进行比较

/*
测试函数的格式
每个测试函数必须导入testing包，测试函数的基本格式（签名）如下：
func TestName(t *testing.T){
    // ...
}

测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头，举几个例子：
func TestAdd(t *testing.T){ ... }
func TestSum(t *testing.T){ ... }
func TestLog(t *testing.T){ ... }

其中参数t用于报告测试失败和附加的日志信息

testing.T的拥有的方法如下：
func (c *T) Error(args ...interface{})
func (c *T) Errorf(format string, args ...interface{})
func (c *T) Fail()
func (c *T) FailNow()
func (c *T) Failed() bool
func (c *T) Fatal(args ...interface{})
func (c *T) Fatalf(format string, args ...interface{})
func (c *T) Log(args ...interface{})
func (c *T) Logf(format string, args ...interface{})
func (c *T) Name() string
func (t *T) Parallel()
func (t *T) Run(name string, f func(t *T)) bool
func (c *T) Skip(args ...interface{})
func (c *T) SkipNow()
func (c *T) Skipf(format string, args ...interface{})
func (c *T) Skipped() bool
*/

/*
测试覆盖率是你的代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。
测试代码覆盖率要在百分之六十以上，不可能百分之百，比如有些if分支永远跑不到，就覆盖不到
测试覆盖率太低，说明你测试用例想的不够，或者代码写的不合理，太多跑不到的分支
测试函数的覆盖率要百分之百，每个函数都要有个测试用例

Go提供内置功能来检查你的代码覆盖率。我们可以使用go test -cover来查看测试覆盖率。
go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/kimalancs/LearningGolang/day09/unitTesting1/split_strings    0.401s
从上面的结果可以看到我们的测试用例覆盖了100%的代码。

Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件。
go test -cover -coverprofile=c.out
上面的命令会将覆盖率相关的信息输出到当前文件夹下面的c.out文件中
然后我们执行go tool cover -html=c.out，使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。
每个用绿色标记的语句块表示被覆盖了，而红色的表示没有被覆盖
*/

// Split split a string by a sep
func Split(str, sep string) []string {
	// var ret []string // 这里不初始化，在for循环里append内置帮ret初始化申请内存，每执行一次for循环会申请一次
	var ret = make([]string, 0, strings.Count(str, sep)+1) // 在此处初始化，并申请足够大的容量，之后就不用申请内存了，优化后，运行时间减少，效率提高
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index]) 
		str = str[index+len(sep):]
		index = strings.Index(str, sep) // str里没有sep，Index返回-1，退出循环
	}
	ret = append(ret, str)
	return ret
}

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}