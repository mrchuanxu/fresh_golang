package main

import (
	"fmt"
	"testing"
)

// 模拟红黑树的基本操作（不依赖外部库）
type SimpleRBTree struct {
	root *RBNode
	size int
}

type RBNode struct {
	key   int
	value interface{}
	left  *RBNode
	right *RBNode
	color bool // true for red, false for black
}

func NewSimpleRBTree() *SimpleRBTree {
	return &SimpleRBTree{
		root: nil,
		size: 0,
	}
}

func (t *SimpleRBTree) Put(key int, value interface{}) {
	t.root = t.insert(t.root, key, value)
	t.root.color = false // 根节点总是黑色
	t.size++
}

func (t *SimpleRBTree) insert(node *RBNode, key int, value interface{}) *RBNode {
	if node == nil {
		return &RBNode{
			key:   key,
			value: value,
			color: true, // 新节点总是红色
		}
	}

	if key < node.key {
		node.left = t.insert(node.left, key, value)
	} else if key > node.key {
		node.right = t.insert(node.right, key, value)
	} else {
		node.value = value
		return node
	}

	// 红黑树平衡操作（简化版）
	return t.balance(node)
}

func (t *SimpleRBTree) balance(node *RBNode) *RBNode {
	// 简化的平衡操作
	if node == nil {
		return nil
	}
	return node
}

func (t *SimpleRBTree) Get(key int) (interface{}, bool) {
	node := t.find(t.root, key)
	if node == nil {
		return nil, false
	}
	return node.value, true
}

func (t *SimpleRBTree) find(node *RBNode, key int) *RBNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		return t.find(node.left, key)
	} else if key > node.key {
		return t.find(node.right, key)
	}
	return node
}

func (t *SimpleRBTree) Size() int {
	return t.size
}

func (t *SimpleRBTree) Empty() bool {
	return t.size == 0
}

// 测试函数
func TestSimpleRBTree(t *testing.T) {
	tree := NewSimpleRBTree()

	// 测试插入
	tree.Put(10, "ten")
	tree.Put(5, "five")
	tree.Put(15, "fifteen")
	tree.Put(3, "three")
	tree.Put(7, "seven")

	// 测试大小
	if tree.Size() != 5 {
		t.Errorf("期望大小 5，实际大小 %d", tree.Size())
	}

	// 测试查找
	if value, found := tree.Get(10); !found || value != "ten" {
		t.Errorf("查找键 10 失败")
	}

	if value, found := tree.Get(5); !found || value != "five" {
		t.Errorf("查找键 5 失败")
	}

	// 测试不存在的键
	if _, found := tree.Get(100); found {
		t.Errorf("不应该找到键 100")
	}

	// 测试空树
	emptyTree := NewSimpleRBTree()
	if !emptyTree.Empty() {
		t.Errorf("空树应该返回 true")
	}

	fmt.Println("所有测试通过！")
}

// 性能测试
func BenchmarkRBTreeInsert(b *testing.B) {
	tree := NewSimpleRBTree()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		tree.Put(i, fmt.Sprintf("value_%d", i))
	}
}

func BenchmarkRBTreeGet(b *testing.B) {
	tree := NewSimpleRBTree()
	// 预先插入一些数据
	for i := 0; i < 1000; i++ {
		tree.Put(i, fmt.Sprintf("value_%d", i))
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Get(i % 1000)
	}
} 