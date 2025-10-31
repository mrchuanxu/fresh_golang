# Go语言Range循环行为详解

## 为什么"循环永动机"不是真正的永动？

### 问题代码
```go
arr := []int{1, 2, 3}
for i, v := range arr {
    arr = append(arr, v) // 添加新元素
}
```

### 实际执行过程

#### 第一次循环
- **i=0, v=1**
- 执行前：`arr = [1, 2, 3]`, `len=3`
- 执行 `arr = append(arr, 1)`
- 执行后：`arr = [1, 2, 3, 1]`, `len=4`

#### 第二次循环
- **i=1, v=2**
- 执行前：`arr = [1, 2, 3, 1]`, `len=4`
- 执行 `arr = append(arr, 2)`
- 执行后：`arr = [1, 2, 3, 1, 2]`, `len=5`

#### 第三次循环
- **i=2, v=3**
- 执行前：`arr = [1, 2, 3, 1, 2]`, `len=5`
- 执行 `arr = append(arr, 3)`
- 执行后：`arr = [1, 2, 3, 1, 2, 3]`, `len=6`

#### 循环结束！
- 虽然数组变长了，但循环只执行了3次
- 最终数组：`[1, 2, 3, 1, 2, 3]`

## 根本原因

### 1. Range循环的编译时特性
```go
// 编译器会将 range arr 转换为：
for i := 0; i < len(arr); i++ {
    v := arr[i]
    // 循环体
}
```

**关键点：**
- `len(arr)` 在循环开始时就被计算并固定
- 即使数组在循环过程中变长，循环次数也不会改变
- 循环次数 = 循环开始时的数组长度

### 2. 与普通for循环的区别

#### Range循环（固定次数）
```go
arr := []int{1, 2, 3}
for i, v := range arr {
    arr = append(arr, v) // 只执行3次
}
```

#### 普通for循环（动态次数）
```go
arr := []int{1, 2, 3}
for i := 0; i < len(arr); i++ {
    arr = append(arr, arr[i]) // 会一直执行！
}
```

**区别：**
- Range循环：`len(arr)` 在循环开始时计算一次
- 普通for循环：`len(arr)` 在每次循环时都重新计算

## 真正的"永动机"实现

### 方法1：使用普通for循环
```go
arr := []int{1, 2, 3}
for i := 0; i < len(arr); i++ {
    arr = append(arr, arr[i])
    fmt.Printf("i=%d, len=%d\n", i, len(arr))
}
// 这个会一直运行，因为len(arr)在每次循环时都会重新计算
```

### 方法2：使用无限循环
```go
arr := []int{1, 2, 3}
for i := 0; ; i++ {
    if i >= len(arr) {
        break
    }
    arr = append(arr, arr[i])
    fmt.Printf("i=%d, len=%d\n", i, len(arr))
}
```

### 方法3：使用channel
```go
ch := make(chan int, 10)
ch <- 1
ch <- 2
ch <- 3

for v := range ch {
    ch <- v // 这会导致真正的永动！
    fmt.Printf("处理: %d\n", v)
}
```

## 性能对比

### Range循环（安全但有限）
```go
arr := []int{1, 2, 3}
for i, v := range arr {
    arr = append(arr, v)
}
// 执行3次，内存使用可控
```

### 普通for循环（危险但可能永动）
```go
arr := []int{1, 2, 3}
for i := 0; i < len(arr); i++ {
    arr = append(arr, arr[i])
}
// 会一直执行，可能导致内存溢出
```

## 实际应用场景

### 1. 安全的循环
```go
// 推荐：使用range循环
for i, v := range items {
    process(v)
    // 不会因为items变化而出现问题
}
```

### 2. 需要动态长度的循环
```go
// 需要动态长度时，使用普通for循环
for i := 0; i < len(queue); i++ {
    item := queue[i]
    if shouldProcess(item) {
        queue = append(queue, generateNewItem())
    }
}
```

### 3. 真正的永动机（服务器程序）
```go
// HTTP服务器
for {
    conn, err := listener.Accept()
    if err != nil {
        continue
    }
    go handleConnection(conn)
}
```

## 总结

1. **Range循环不是真正的永动机**：循环次数在开始时就被确定
2. **普通for循环可能永动**：因为会重新计算长度
3. **设计目的**：Range循环的设计是为了安全和可预测性
4. **实际应用**：真正的永动机通常用于服务器、监控等场景

### 关键理解
- Range循环 = 编译时确定循环次数
- 普通for循环 = 运行时动态计算循环次数
- 真正的永动机 = 无限循环 + 外部中断机制 