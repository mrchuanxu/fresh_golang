#!/bin/bash

echo "=== 广度优先搜索(BFS)测试 ==="
echo ""

# 运行所有测试
echo "运行所有BFS测试用例..."
go test -v bfs_test.go

echo ""
echo "=== 运行基准测试 ==="
go test -bench=. bfs_test.go

echo ""
echo "=== 运行特定测试用例 ==="
echo "1. 线性图测试"
go test -v -run TestBFSLinearGraph bfs_test.go

echo ""
echo "2. 环形图测试"
go test -v -run TestBFSCyclicGraph bfs_test.go

echo ""
echo "3. 无路径测试"
go test -v -run TestBFSNoPath bfs_test.go

echo ""
echo "=== 测试完成 ===" 