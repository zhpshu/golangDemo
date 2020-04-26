package main

import (
	"errors"
	"fmt"
	"math"
)

type Person struct {
	name string
	age  int
}

func main() {
	x := 5
	var y int = 7
	var result = sum(x, y)
	fmt.Println(result)

	//exception
	result2, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	var m = make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	fmt.Println(m)
	for key, value := range m {
		fmt.Println("key", key, "valu", value)
	}

	var arr = [6]int{1, 2, 3, 4, 5, 6}
	for index, value := range arr {
		fmt.Println("index", index, "valu", value)
	}
	fmt.Println("hello,world,{}", arr)

	person := Person{name: "10", age: 11}
	fmt.Println(person.age)

	//point
	p := 10
	//& 是取地址符号 , 即取得某个变量的地址
	fmt.Println(&p)
	inc(&p)
	fmt.Println(p)

	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
}
func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined number")
	}
	return math.Sqrt(x), nil
}

//指针类型的变量存储的是一个地址，所以又叫指针类型或引用类型，默认值为nil，也就是空地址
func inc(x *int) {
	*x++
}
