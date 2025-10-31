package main

import "fmt"

func main() {
	fmt.Println("=== Range循环行为测试 ===")
	
	// 测试1：基础range循环
	fmt.Println("\n1. 基础range循环:")
	arr1 := []int{1, 2, 3}
	fmt.Printf("初始数组: %v, 长度: %d\n", arr1, len(arr1))
	
	for i, v := range arr1 {
		fmt.Printf("循环 %d: i=%d, v=%d, 当前长度=%d\n", i+1, i, v, len(arr1))
		arr1 = append(arr1, v)
		fmt.Printf("添加后: %v, 长度: %d\n", arr1, len(arr1))
	}
	fmt.Printf("最终数组: %v\n", arr1)
	
	// 测试2：使用索引遍历
	fmt.Println("\n2. 使用索引遍历:")
	arr2 := []int{1, 2, 3}
	fmt.Printf("初始数组: %v, 长度: %d\n", arr2, len(arr2))
	
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("循环 %d: i=%d, v=%d, 当前长度=%d\n", i+1, i, arr2[i], len(arr2))
		arr2 = append(arr2, arr2[i])
		fmt.Printf("添加后: %v, 长度: %d\n", arr2, len(arr2))
	}
	fmt.Printf("最终数组: %v\n", arr2)
	
	// 测试3：真正的"永动机"
	fmt.Println("\n3. 真正的永动机:")
	arr3 := []int{1, 2, 3}
	fmt.Printf("初始数组: %v, 长度: %d\n", arr3, len(arr3))
	
	// 使用无限循环 + 索引
	for i := 0; ; i++ {
		if i >= len(arr3) {
			break
		}
		fmt.Printf("循环 %d: i=%d, v=%d, 当前长度=%d\n", i+1, i, arr3[i], len(arr3))
		arr3 = append(arr3, arr3[i])
		fmt.Printf("添加后: %v, 长度: %d\n", arr3, len(arr3))
	}
	fmt.Printf("最终数组: %v\n", arr3)
	
	// 测试4：使用for循环遍历
	fmt.Println("\n4. 使用for循环遍历:")
	arr4 := []int{1, 2, 3}
	fmt.Printf("初始数组: %v, 长度: %d\n", arr4, len(arr4))
	
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("循环 %d: i=%d, v=%d, 当前长度=%d\n", i+1, i, arr4[i], len(arr4))
		arr4 = append(arr4, arr4[i])
		fmt.Printf("添加后: %v, 长度: %d\n", arr4, len(arr4))
	}
	fmt.Printf("最终数组: %v\n", arr4)
} 