package main

import "fmt"

func main() {
	// array静态数组
	// 大小在声明时就固定
	// 且数组大小不能使用变量，比如
	// var a = 10
	// wrong: var array [a]int
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = i
	}

	// slice动态数组(切片)
	var slice []int // 此时slice为nil
	for i := 0; i < 10; i++ {
		// append会根据数组动态变化
		// append指向的地址可能会产生变化
		// 同时len和cap会动态增长
		// cap不够时正常会以2倍增长
		// 因为扩容需要数据搬迁，影响效率
		slice = append(slice, i)
		fmt.Printf("slice addr: %p len: %d cap: %d\n", slice, len(slice), cap(slice))
	}

	// slice动态数组(切片)
	// 使用make初始化后, 就如同静态数组一样使用
	// 但同样可以使用append, append之后地址也会变化
	var slice2 = make([]int, 10)
	for i := 0; i < 10; i++ {
		slice2[i] = i
		fmt.Printf("slice2 addr: %p len: %d cap: %d\n", slice2, len(slice2), cap(slice2))
	}
}
