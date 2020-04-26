package main

import "fmt"

//1
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {

	s.Age = age + 1
}

//3
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

/**
这段代码在 #1 定义了 interface I，
在 #2 用 struct S 实现了 I 定义的两个方法，
接着在 #3 定义了一个函数 f 参数类型是 I，S 实现了 I 的两个方法就说 S 是 I 的实现者，执行 f(&s) 就完了一次 interface 类型的使用。
*/
func main() {
	s := S{}
	f(&s) //4
}
