#!/bin/bash

echo "=== Go语言红黑树项目设置 ==="

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "错误: Go未安装，请先安装Go语言环境"
    echo "下载地址: https://golang.org/dl/"
    exit 1
fi

echo "Go版本: $(go version)"

# 初始化模块
echo "初始化Go模块..."
go mod init redblacktree-example

# 安装依赖
echo "安装红黑树库..."
go get github.com/emirpasic/gods

# 下载依赖
echo "下载依赖..."
go mod tidy

echo "=== 设置完成 ==="
echo ""
echo "运行示例:"
echo "  go run redblacktree_advanced.go"
echo ""
echo "运行测试:"
echo "  go test redblacktree_test.go"
echo ""
echo "性能测试:"
echo "  go test -bench=. redblacktree_test.go" 