# Go语言红黑树使用指南

## 简介

本项目展示了如何在Go语言中使用红黑树数据结构。我们使用 `github.com/emirpasic/gods` 库，这是一个功能强大的Go数据结构库。

## 安装依赖

### 1. 初始化模块
```bash
go mod init redblacktree-example
```

### 2. 安装依赖
```bash
go get github.com/emirpasic/gods
```

### 3. 运行示例
```bash
go run redblacktree_advanced.go
```

## 红黑树特性

红黑树是一种自平衡的二叉搜索树，具有以下特性：

1. **自平衡**: 通过颜色标记和旋转操作保持树的平衡
2. **高效操作**: 插入、删除、查找的时间复杂度都是 O(log n)
3. **有序性**: 中序遍历可以得到有序序列
4. **范围查询**: 支持高效的范围查询操作

## 主要功能

### 基础操作
- `Put(key, value)`: 插入键值对
- `Get(key)`: 查找值
- `Remove(key)`: 删除键值对
- `Contains(key)`: 检查键是否存在
- `Size()`: 获取树的大小
- `Empty()`: 检查是否为空
- `Clear()`: 清空树

### 遍历操作
- `Each(func)`: 遍历所有节点
- `Iterator()`: 获取迭代器
- `Left()`: 获取最小值
- `Right()`: 获取最大值

### 比较器
- `NewWithIntComparator()`: 整数比较器
- `NewWithStringComparator()`: 字符串比较器
- `NewWith(comparator)`: 自定义比较器

## 使用场景

1. **有序数据存储**: 需要保持数据有序的场景
2. **范围查询**: 需要高效范围查询的应用
3. **字典实现**: 键值对的高效存储和检索
4. **事件调度**: 时间相关的事件调度系统
5. **缓存实现**: 需要有序访问的缓存系统

## 性能特点

- **时间复杂度**:
  - 插入: O(log n)
  - 删除: O(log n)
  - 查找: O(log n)
  - 范围查询: O(log n + k)，其中k是结果数量

- **空间复杂度**: O(n)

## 示例代码

项目包含以下示例：

1. **基础操作示例**: 展示红黑树的基本操作
2. **自定义对象**: 使用结构体作为键的示例
3. **范围查询**: 展示如何进行范围查询
4. **性能测试**: 大量数据的性能测试

## 注意事项

1. 确保比较器函数正确实现
2. 自定义对象需要实现合适的比较逻辑
3. 大量数据操作时注意内存使用
4. 并发环境下需要考虑线程安全

## 其他选择

除了 `gods` 库，还有其他选择：

1. **标准库**: Go 1.21+ 可以使用 `golang.org/x/exp/maps`
2. **其他第三方库**: 
   - `github.com/your-org/your-tree`
   - `github.com/another-org/another-tree`

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目。
