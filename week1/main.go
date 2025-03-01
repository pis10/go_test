/*
作业：实现切片的删除操作
要求：
1. 实现删除操作
2. 高性能实现
3. 泛型方法
4. 支持缩容机制
*/
package main

import (
	"errors"
	"fmt"
)

func Delete[T any](s []T, i int) ([]T, error) {
	if i < 0 || i >= len(s) { // 修正索引检查条件
		return nil, errors.New("index out of range")
	}
	ns := append(s[:i], s[i+1:]...)
	// 缩容机制：当容量超过两倍长度时缩容
	newLen := len(s) - 1
	if cap(s) >= newLen*2 && newLen >= 4 {
		tmp := make([]T, 0, cap(s)/2)
		ns = append(tmp, ns...)
	}
	return ns, nil
}

func main() {
	// 测试基础功能
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Delete(test, 3)) // 正确删除第4个元素

	// 测试缩容功能
	largeCap := make([]int, 5, 20) // 容量20，长度5
	for i := 0; i < 5; i++ {
		largeCap[i] = i + 1
	}
	result, _ := Delete(largeCap, 2)
	fmt.Printf("缩容测试: len=%d cap=%d\n", len(result), cap(result)) // 预期容量从20缩容到8
}
