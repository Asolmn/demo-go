package main

import "fmt"

type number interface {
	int | int32 | uint32 | int64 | uint64 | float32 | float64
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32
}

func Add[T int | float32 | float64](a, b T) T {
	return a + b
}

func addInteger[T Integer](a, b T) T {
	return a + b
}

func inSlice[T comparable](s []T, e T) int {
	for k, v := range s {
		if v == e {
			return k
		}
	}
	return -1
}

func Print[T any](s T) {
	fmt.Println(s)
}

// 类型
type Student1[T int | string] struct {
	Name string
	Data []T
}

type Student2[T []int | []string] struct {
	Name string
	Data T
}

// 匿名函数不支持泛型

// 但可以使用别处定义的类型实参 (% 运算符，需要把类型约束为整型)
func MyFunc[T int](a, b T) {
	fn2 := func(i, j T) T {
		return i % j
	}
	fn2(a, b)
}

// 方法不支持泛型
