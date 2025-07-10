package main

import (
	"fmt"

	"github.com/emirpasic/gods/trees/redblacktree"
)

// 学生结构体
type Student struct {
	ID   int
	Name string
	Age  int
}

// 学生比较器
func studentComparator(a, b interface{}) int {
	studentA := a.(Student)
	studentB := b.(Student)
	
	if studentA.ID < studentB.ID {
		return -1
	} else if studentA.ID > studentB.ID {
		return 1
	}
	return 0
}

func main() {
	// 基础红黑树操作
	basicOperations()
	
	// 学生信息红黑树
	studentTreeExample()
	
	// 范围查询示例
	rangeQueryExample()
}

func basicOperations() {
	fmt.Println("=== 基础红黑树操作 ===")
	
	tree := redblacktree.NewWithIntComparator()
	
	// 插入操作
	keys := []int{50, 30, 70, 20, 40, 60, 80}
	for _, key := range keys {
		tree.Put(key, fmt.Sprintf("value_%d", key))
	}
	
	fmt.Printf("树的大小: %d\n", tree.Size())
	
	// 查找操作
	if value, found := tree.Get(30); found {
		fmt.Printf("找到键 30: %v\n", value)
	}
	
	// 检查键是否存在
	fmt.Printf("键 25 是否存在: %v\n", tree.Contains(25))
	fmt.Printf("键 50 是否存在: %v\n", tree.Contains(50))
	
	// 获取最小值和最大值
	if min := tree.Left(); min != nil {
		fmt.Printf("最小值: %v\n", min.Key)
	}
	if max := tree.Right(); max != nil {
		fmt.Printf("最大值: %v\n", max.Key)
	}
	
	// 删除操作
	tree.Remove(30)
	fmt.Printf("删除键 30 后，树的大小: %d\n", tree.Size())
	
	// 清空树
	tree.Clear()
	fmt.Printf("清空后，树是否为空: %v\n", tree.Empty())
}

func studentTreeExample() {
	fmt.Println("\n=== 学生信息红黑树 ===")
	
	// 创建使用自定义比较器的红黑树
	tree := redblacktree.NewWith(studentComparator)
	
	// 添加学生信息
	students := []Student{
		{ID: 1001, Name: "张三", Age: 20},
		{ID: 1002, Name: "李四", Age: 21},
		{ID: 1003, Name: "王五", Age: 19},
		{ID: 1004, Name: "赵六", Age: 22},
	}
	
	for _, student := range students {
		tree.Put(student, student.Name)
	}
	
	fmt.Printf("学生树大小: %d\n", tree.Size())
	
	// 查找特定学生
	targetStudent := Student{ID: 1002, Name: "", Age: 0}
	if value, found := tree.Get(targetStudent); found {
		fmt.Printf("找到学生 ID 1002: %v\n", value)
	}
	
	// 遍历所有学生
	fmt.Println("所有学生信息:")
	tree.Each(func(key interface{}, value interface{}) {
		student := key.(Student)
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d\n", student.ID, student.Name, student.Age)
	})
}

func rangeQueryExample() {
	fmt.Println("\n=== 范围查询示例 ===")
	
	tree := redblacktree.NewWithIntComparator()
	
	// 插入一些数据
	for i := 1; i <= 10; i++ {
		tree.Put(i, fmt.Sprintf("item_%d", i))
	}
	
	// 查找大于等于 3 且小于 7 的元素
	fmt.Println("范围查询 [3, 7):")
	it := tree.Iterator()
	for it.Next() {
		key := it.Key().(int)
		if key >= 3 && key < 7 {
			fmt.Printf("键: %d, 值: %v\n", key, it.Value())
		}
	}
	
	// 查找小于等于 5 的元素
	fmt.Println("小于等于 5 的元素:")
	it = tree.Iterator()
	for it.Next() {
		key := it.Key().(int)
		if key <= 5 {
			fmt.Printf("键: %d, 值: %v\n", key, it.Value())
		}
	}
}

// 红黑树性能测试
func performanceTest() {
	fmt.Println("\n=== 性能测试 ===")
	
	tree := redblacktree.NewWithIntComparator()
	
	// 插入大量数据
	for i := 0; i < 10000; i++ {
		tree.Put(i, fmt.Sprintf("value_%d", i))
	}
	
	fmt.Printf("插入 10000 个元素后，树的大小: %d\n", tree.Size())
	
	// 查找测试
	if _, found := tree.Get(5000); found {
		fmt.Println("成功找到元素 5000")
	}
	
	// 删除测试
	tree.Remove(5000)
	fmt.Printf("删除元素 5000 后，树的大小: %d\n", tree.Size())
} 