package main

import "fmt"

func main() {

	fmt.Println(addInteger(1, 2))
	fmt.Println(addInteger(-1, -2))

	fmt.Println(inSlice([]string{"a", "b", "c"}, "c"))

	s := []int{1, 2, 3}
	// 显示指定参数类型
	Print[[]int](s)

	// 推断参数类型
	Print(s)
	return
}
