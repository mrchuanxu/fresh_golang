package main

import (
	"fmt"
	"testing"
)

// 修复后的buildPath函数
func buildPath(prev []int, s, t int) []int {
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

// 测试buildPath函数
func TestBuildPath(t *testing.T) {
	fmt.Println("=== buildPath函数测试 ===")
	
	// 测试用例1: 正常路径
	prev := []int{-1, 0, 1, 2, 3} // 0->1->2->3->4
	path := buildPath(prev, 0, 4)
	expected := []int{0, 1, 2, 3, 4}
	
	if !compareSlices(path, expected) {
		t.Errorf("期望路径 %v，实际路径 %v", expected, path)
	} else {
		fmt.Printf("✓ 正常路径测试通过，路径: %v\n", path)
	}
	
	// 测试用例2: 无路径
	prev2 := []int{-1, 0, -1, -1, -1} // 只有0->1，没有到4的路径
	path2 := buildPath(prev2, 0, 4)
	
	if path2 == nil {
		fmt.Println("✓ 无路径测试通过")
	} else {
		t.Errorf("期望nil，实际得到路径: %v", path2)
	}
	
	// 测试用例3: 相同起点终点
	prev3 := []int{-1, 0, 1, 2, 3}
	path3 := buildPath(prev3, 0, 0)
	expected3 := []int{0}
	
	if !compareSlices(path3, expected3) {
		t.Errorf("期望路径 %v，实际路径 %v", expected3, path3)
	} else {
		fmt.Printf("✓ 相同起点终点测试通过，路径: %v\n", path3)
	}
}

// 辅助函数：比较两个切片
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

// 完整的Graph结构用于测试
type Graph struct {
	points int
	edges  int
	adj    map[int][]int // 简化版本，使用切片而不是红黑树
}

func InitGraph(points int) *Graph {
	return &Graph{
		points: points,
		edges:  0,
		adj:    make(map[int][]int),
	}
}

func (g *Graph) AddEdge(s, t int) {
	if g.adj[s] == nil {
		g.adj[s] = make([]int, 0)
	}
	g.adj[s] = append(g.adj[s], t)
	
	if g.adj[t] == nil {
		g.adj[t] = make([]int, 0)
	}
	g.adj[t] = append(g.adj[t], s)
	g.edges++
}

func (g *Graph) BFS(s, t int) []int {
	if s == t {
		return []int{s}
	}
	
	queue := make([]int, 0)
	queue = append(queue, s)
	visited := make(map[int]bool)
	visited[s] = true
	prev := make([]int, g.points)
	for i := 0; i < g.points; i++ {
		prev[i] = -1
	}
	
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		
		for _, neighbor := range g.adj[cur] {
			if !visited[neighbor] {
				prev[neighbor] = cur
				if neighbor == t {
					return buildPath(prev, s, t)
				}
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	
	return nil
}

// 测试完整的BFS实现
func TestBFSWithFixedBuildPath(t *testing.T) {
	fmt.Println("=== 使用修复后的buildPath测试BFS ===")
	
	graph := InitGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	
	path := graph.BFS(0, 4)
	expected := []int{0, 1, 2, 3, 4}
	
	if !compareSlices(path, expected) {
		t.Errorf("期望路径 %v，实际路径 %v", expected, path)
	} else {
		fmt.Printf("✓ BFS测试通过，路径: %v\n", path)
	}
} 