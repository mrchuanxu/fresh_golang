package ch6alg_test

import (
	"fmt"
	"testing"

	"github.com/mrchuanxu/vito_infra/alg"
)



func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 测试用例
func TestBFSLinearGraph(t *testing.T) {
	fmt.Println("=== 线性图测试 ===")
	
	graph := alg.InitGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	
	path := graph.BFS(0, 4)
	expected := []int{0, 1, 2, 3, 4}
	
	if !compareSlices(path, expected) {
		t.Errorf("期望路径 %v，实际路径 %v", expected, path)
	} else {
		fmt.Printf("✓ 线性图测试通过，路径: %v\n", path)
	}
}

// 环形图测试只能用于单向图，双向图会死循环或者输出错误路径
func TestBFSCyclicGraph(t *testing.T) {
	fmt.Println("=== 环形图测试 ===")
	
	graph := alg.InitGraph(4)
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 0)
	
	path := graph.BFS(0, 3)
	expected := []int{0, 1, 2, 3}
	
	if !compareSlices(path, expected) {
		t.Errorf("期望路径 %v，实际路径 %v", expected, path)
	} else {
		fmt.Printf("✓ 环形图测试通过，路径: %v\n", path)
	}
}

func TestBFSNoPath(t *testing.T) {
	fmt.Println("=== 无路径测试 ===")
	
	graph := alg.InitGraph(4)
	graph.AddEdge(0, 1)
	graph.AddEdge(2, 3)
	
	path := graph.BFS(0, 3)
	
	if path == nil {
		fmt.Println("✓ 无路径测试通过")
	} else {
		t.Errorf("期望nil，实际得到路径: %v", path)
	}
}

func BenchmarkBFS(b *testing.B) {
	graph := alg.InitGraph(100)
	
	for i := 0; i < 50; i++ {
		graph.AddEdge(i, i+1)
		graph.AddEdge(i, i+10)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		graph.BFS(0, 99)
	}
} 