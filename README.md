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
