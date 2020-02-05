package split_strings

import (
	"reflect"
	"testing"
)

// 测试组
func TestSplit(t *testing.T) {
	type testCase struct {
		str string
		sep string
		want []string
	}

	testGroups := []testCase{
		testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		testCase{"abcdef","bc", []string{"a", "def"}},
		testCase{"黄山落叶松叶落山黄","落叶松", []string{"黄山","叶落山黄"}},
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
		str string
		sep string
		want []string
	}

	testGroups := map[string]testCase{
		"case1": testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case3": testCase{"abcdef","bc", []string{"a", "def"}},
		"case4": testCase{"黄山落叶松叶落山黄","落叶松", []string{"黄山","叶落山黄"}},
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
		str string
		sep string
		want []string
	}

	testGroups := map[string]testCase{
		"case1": testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case3": testCase{"abcdef","bc", []string{"a", "def"}},
		"case4": testCase{"黄山落叶松叶落山黄","落叶松", []string{"黄山","叶落山黄"}},
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
	got := Split("babcbef", "b") // 程序输出的结果
	want := []string{"", "a", "c", "ef"} // 期望的结果
	if !reflect.DeepEqual(got, want) { // 因为slice不能比较直接，借助反射包中的方法比较
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


