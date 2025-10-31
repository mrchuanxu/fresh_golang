#!/bin/bash

echo "=================================="
echo "    Go语言 Select 语句测试脚本"
echo "=================================="

# 切换到正确的目录
cd easygolang/ch7

echo "正在运行 Select 语句基础功能测试..."
echo ""

echo "1. 基本 Select 用法测试"
go test -v -run TestBasicSelect

echo ""
echo "2. Select Default 用法测试"
go test -v -run TestSelectWithDefault

echo ""
echo "3. Select 随机性测试"
go test -v -run TestSelectPriority

echo ""
echo "4. Select 超时处理测试"
go test -v -run TestSelectTimeout

echo ""
echo "5. Select Context 取消测试"
go test -v -run TestSelectContext

echo ""
echo "6. Select 工作池模式测试"
go test -v -run TestSelectWorkerPool

echo ""
echo "7. Select Fan-in 模式测试"
go test -v -run TestSelectFanInPattern

echo ""
echo "8. Select 竞态条件避免测试"
go test -v -run TestSelectRaceCondition

echo ""
echo "9. Select 内存泄漏避免测试"
go test -v -run TestSelectMemoryLeak

echo ""
echo "10. Select 性能测试"
go test -v -run TestSelectPerformance

echo ""
echo "=================================="
echo "    所有测试完成"
echo "=================================="

# 性能基准测试
echo ""
echo "运行性能基准测试..."
go test -bench=BenchmarkSelect -benchmem

echo ""
echo "如果需要查看详细的实现原理，请查看文件："
echo "  - easygolang/ch7/select_implementation_principle.md"
echo ""
echo "如果需要查看测试代码，请查看文件："
echo "  - easygolang/ch7/select_statement_test.go" 