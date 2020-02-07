package splitstrings

import (
	"reflect"
	"testing"
)

// 测试组
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroups := []testCase{
		testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		testCase{"abcdef", "bc", []string{"a", "def"}},
		testCase{"黄山落叶松叶落山黄", "落叶松", []string{"黄山", "叶落山黄"}},
	}

	for _, tc := range testGroups {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%#v, got:%#v\n", tc.want, got)
		}
	}
}

// 组测试，如果测试用例比较多的时候，我们是没办法一眼看出来具体是哪个测试用例失败了

// 采用map，给每个测试用例加个name，可以解决问题，报错时可以知道哪个测试用例失败
func Test4Split(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroups := map[string]testCase{
		"case1": testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case3": testCase{"abcdef", "bc", []string{"a", "def"}},
		"case4": testCase{"黄山落叶松叶落山黄", "落叶松", []string{"黄山", "叶落山黄"}},
	}

	for name, tc := range testGroups {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("name:%s, excepted:%#v, got:%#v\n", name, tc.want, got)
		}
	}
}

// Go1.7+中新增了子测试，我们可以按照如下方式使用t.Run执行子测试
// 子测试，在测试时会显示每一个用例的结果，更清晰，
// 可以通过-run=RegExp来指定运行的测试用例，如go test -v -run=TestSplit，只会跑TestSplit测试用例
// 还可以通过 / 来指定要运行的子测试用例，例如：go test -v -run=Test5Split/case3只会运行case3对应的子测试用例
func Test5Split(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroups := map[string]testCase{
		"case1": testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case3": testCase{"abcdef", "bc", []string{"a", "def"}},
		"case4": testCase{"黄山落叶松叶落山黄", "落叶松", []string{"黄山", "叶落山黄"}},
	}

	for name, tc := range testGroups {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%s, excepted:%#v, got:%#v\n", name, tc.want, got)
			}
		})
	}
}

func Test1Split(t *testing.T) { //测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("babcbef", "b")         // 程序输出的结果
	want := []string{"", "a", "c", "ef"} // 期望的结果
	if !reflect.DeepEqual(got, want) {   // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%#v, got:%#v\n", want, got) //测试失败输出错误提示
	}
}

func Test2Split(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("excepted:%#v, got:%#v\n", want, got)
	}
}
func Test3Split(t *testing.T) {
	got := Split("abcef", "bc")
	want := []string{"a", "ef"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("excepted:%#v, got:%#v\n", want, got)
	}
}

// benchmark 基准测试

// BenchmarkSplit split
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

/*
基准测试并不会默认执行，需要增加-bench参数，所以我们通过执行go test -bench=Split命令执行基准测试
go test -bench=Split

goos: darwin
goarch: amd64
pkg: github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings
BenchmarkSplit-8         6302373               180 ns/op
PASS
ok      github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings     1.417s

其中BenchmarkSplit-8表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测试很重要。
6302373和180ns/op表示每次调用Split函数耗时180ns，这个结果是6302373次调用的平均值


为基准测试添加-benchmem参数，来获得内存分配的统计数据
go test -bench=Split -benchmem

goos: darwin
goarch: amd64
pkg: github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings
BenchmarkSplit-8         6299938               181 ns/op             112 B/op          3 allocs/op
PASS
ok      github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings     1.642s

112 B/op表示每次操作内存分配了112字节，3 allocs/op则表示每次操作进行了3次内存分配
优化后，内存分配次数减少了三分之二，内存分配减少了一半

goos: darwin
goarch: amd64
pkg: github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings
BenchmarkSplit-8        12538065                85.7 ns/op            48 B/op          1 allocs/op
PASS
ok      github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings     1.557s
*/


//上面的基准测试只能得到给定操作的绝对耗时
//但是在很多性能问题是发生在两个不同操作之间的相对耗时
//比如同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别是多少？
//再或者对于同一个任务究竟使用哪种算法性能最佳？我们通常需要对两个不同算法的实现使用相同的输入来进行基准比较测试。

//性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用。举个例子如下：
//func benchmark(b *testing.B, size int){/* ... */}
//func Benchmark10(b *testing.B){ benchmark(b, 10) }
//func Benchmark100(b *testing.B){ benchmark(b, 100) }
//func Benchmark1000(b *testing.B){ benchmark(b, 1000) }

// go test -bench=. 执行所有性能测试
// go test -bench=Fib2 指定一个测试用例

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

/*
goos: darwin
goarch: amd64
pkg: github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings
BenchmarkSplit-8        13197652                85.6 ns/op
BenchmarkFib1-8         674709469                1.75 ns/op
BenchmarkFib2-8         232592122                5.09 ns/op
BenchmarkFib3-8         140936972                8.55 ns/op
BenchmarkFib10-8         3688480               321 ns/op
BenchmarkFib20-8           30052             39620 ns/op
BenchmarkFib40-8               2         602472130 ns/op
PASS
ok      github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings     12.535s

这里需要注意的是，默认情况下，每个基准测试至少运行1秒
如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。

最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒
像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果
go test -bench=Fib40 -benchtime=20s

goos: darwin
goarch: amd64
pkg: github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings
BenchmarkFib40-8              38         601440390 ns/op
PASS
ok      github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings     23.956s

这一次BenchmarkFib40函数运行了50次，结果就会更准确一些了。

使用性能比较函数做测试的时候一个容易犯的错误就是把b.N作为输入的大小
错误示范1
func BenchmarkFibWrong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(n)
	}
}

错误示范2
func BenchmarkFibWrong2(b *testing.B) {
	Fib(b.N)
}
*/

/*
重置时间
b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。例如：

func BenchmarkSplit(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
*/


// func (b *B) RunParallel(body func(*PB))会以并行的方式执行给定的基准测试。
// RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行
// 其中goroutine数量的默认值为GOMAXPROCS
// 用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性
// 那么可以在RunParallel之前调用SetParallelism
// RunParallel通常会与-cpu标志一同使用。

func BenchmarkSplitParallel(b *testing.B) {
	//b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("a:b:c", ":")
		}
	})
}

/*
go test -bench=Split

goos: darwin
goarch: amd64
pkg: github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings
BenchmarkSplit-8                13297350                86.0 ns/op
BenchmarkSplitParallel-8        41169416                25.1 ns/op
PASS
ok      github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings     4.971s
*/

/*
Setup与TearDown
测试程序有时需要在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）。

TestMain
通过在*_test.go文件中定义TestMain函数来可以在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）操作。

如果测试文件包含函数:func TestMain(m *testing.M)那么生成的测试会先调用TestMain(m)，然后再运行具体测试
TestMain运行在主goroutine中, 可以在调用m.Run前后做任何设置（setup）和拆卸（teardown）
退出测试的时候应该使用m.Run的返回值作为参数调用os.Exit

一个使用TestMain来设置Setup和TearDown的示例如下：
func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	os.Exit(retCode)                           // 退出测试
}
需要注意的是：在调用TestMain时, flag.Parse并没有被调用
所以如果TestMain依赖于command-line标志 (包括 testing 包的标记), 则应该显示的调用flag.Parse


子测试的Setup与Teardown

有时候我们可能需要为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown。下面我们定义两个函数工具函数如下：

// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}
使用方式如下：

func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdone操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)           // 测试之后执行testdone操作
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}
测试结果如下：

split $ go test -v
=== RUN   TestSplit
=== RUN   TestSplit/simple
=== RUN   TestSplit/wrong_sep
=== RUN   TestSplit/more_sep
=== RUN   TestSplit/leading_sep
--- PASS: TestSplit (0.00s)
    split_test.go:71: 如有需要在此执行:测试之前的setup
    --- PASS: TestSplit/simple (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    --- PASS: TestSplit/wrong_sep (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    --- PASS: TestSplit/more_sep (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    --- PASS: TestSplit/leading_sep (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    split_test.go:73: 如有需要在此执行:测试之后的teardown
=== RUN   ExampleSplit
--- PASS: ExampleSplit (0.00s)
PASS
ok      github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings       0.006s
*/

/*
示例函数的格式

被go test特殊对待的第三种函数就是示例函数，它们的函数名以Example为前缀。它们既没有参数也没有返回值。标准格式如下：

func ExampleName() {
    // ...
}

下面的代码是我们为Split函数编写的一个示例函数：

func ExampleSplit() {
	fmt.Println(split.Split("a:b:c", ":"))
	fmt.Println(split.Split("沙河有沙又有河", "沙"))
	// Output:
	// [a b c]
	// [ 河有 又有河]
}
为你的代码编写示例代码有如下三个用处：

示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。

示例函数只要包含了// Output:也是可以通过go test运行的可执行测试。

split $ go test -run Example
PASS
ok      github.com/kimalancs/LearningGolang/day09/unitTesting1/splitstrings      0.006s
示例函数提供了可以直接运行的示例代码，可以直接在golang.org的godoc文档服务器上使用Go Playground运行示例代码。

*/