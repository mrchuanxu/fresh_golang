package main

import (
	"fmt"
)

// 原始的buildPath函数（有问题）
func buildPathOriginal(prev []int, s, t int) []int {
	path := make([]int, 0)
	cur := t
	for cur != -1 {
		path = append([]int{cur}, path...)
		cur = prev[cur]
	}
	return path
}

// 修复后的buildPath函数
func buildPathFixed(prev []int, s, t int) []int {
	// 检查是否有路径
	if prev[t] == -1 && s != t {
		return nil // 没有找到路径
	}
	
	// 构建路径
	path := make([]int, 0)
	cur := t
	
	// 从终点回溯到起点
	for cur != -1 {
		path = append([]int{cur}, path...)
		cur = prev[cur]
	}
	
	// 验证路径的有效性
	if len(path) > 0 && path[0] == s && path[len(path)-1] == t {
		return path
	}
	
	return nil
}

func main() {
	fmt.Println("=== buildPath函数问题演示 ===")
	
	// 测试用例1: 正常路径
	prev := []int{-1, 0, 1, 2, 3} // 0->1->2->3->4
	fmt.Println("测试用例1: 正常路径")
	fmt.Printf("prev数组: %v\n", prev)
	
	// 使用原始函数
	path1 := buildPathOriginal(prev, 0, 4)
	fmt.Printf("原始函数结果: %v\n", path1)
	
	// 使用修复函数
	path2 := buildPathFixed(prev, 0, 4)
	fmt.Printf("修复函数结果: %v\n", path2)
	
	// 测试用例2: 无路径情况
	prev2 := []int{-1, 0, -1, -1, -1} // 只有0->1，没有到4的路径
	fmt.Println("\n测试用例2: 无路径情况")
	fmt.Printf("prev数组: %v\n", prev2)
	
	// 使用原始函数
	path3 := buildPathOriginal(prev2, 0, 4)
	fmt.Printf("原始函数结果: %v\n", path3)
	
	// 使用修复函数
	path4 := buildPathFixed(prev2, 0, 4)
	fmt.Printf("修复函数结果: %v\n", path4)
	
	// 测试用例3: 相同起点终点
	prev3 := []int{-1, 0, 1, 2, 3}
	fmt.Println("\n测试用例3: 相同起点终点")
	fmt.Printf("prev数组: %v\n", prev3)
	
	// 使用原始函数
	path5 := buildPathOriginal(prev3, 0, 0)
	fmt.Printf("原始函数结果: %v\n", path5)
	
	// 使用修复函数
	path6 := buildPathFixed(prev3, 0, 0)
	fmt.Printf("修复函数结果: %v\n", path6)
	
	fmt.Println("\n=== 问题总结 ===")
	fmt.Println("原始buildPath函数的问题:")
	fmt.Println("1. 没有检查是否存在路径")
	fmt.Println("2. 没有验证路径的有效性")
	fmt.Println("3. 可能返回无效路径")
	
	fmt.Println("\n修复后的buildPath函数的改进:")
	fmt.Println("1. 添加了路径存在性检查")
	fmt.Println("2. 添加了路径有效性验证")
	fmt.Println("3. 正确处理边界情况")
} 