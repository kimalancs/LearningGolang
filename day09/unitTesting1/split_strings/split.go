package split_strings

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

// Split split a string by a sep
func Split(str, sep string) []string {
	var ret []string
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep) // str里没有sep，Index返回-1，退出循环
	}
	ret = append(ret, str)
	return ret
}
