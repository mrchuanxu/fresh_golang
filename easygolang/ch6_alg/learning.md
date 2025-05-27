### learn data and algrithm
![Data and Algorithm](数据结构与算法.webp)

### 时间复杂度
时间复杂度（Time Complexity）和空间复杂度（Space Complexity）常用大 O 符号表示：

- 时间复杂度：$O(f(n))$，表示算法执行所需的基本操作次数随输入规模 $n$ 的增长情况。
- 空间复杂度：$O(g(n))$，表示算法运行时所需的额外存储空间随输入规模 $n$ 的增长情况。

常见复杂度示例：$O(1)$、$O(\log n)$、$O(n)$、$O(n^2)$ 等。<br>
其中计算的时间单位为unit_time,因每个系统执行的时间都不一样。时间复杂度公式只是表示随数据规模增长的变化趋势。<br>

O(log2N)
```go
i := 1
for i<=n{
    i = i*2;
}
// 取 2的x次
```

空间复杂度公式只是表示随着算法的存储空间和数据规模之间的增长关系。
T(n)
```go
    arr := [100]int{}
    for i:= 0;i<100;i++{
        arr[i] = i*i
    }

	for i:= 100-1;i>=0;i--{
		fmt.Println(arr[i])
    }	
```
均摊 平均 最好 最坏的复杂度，可以通过运行程序的阅读计算出来。


### 数组
数组是一种线性表数据结构，用一组连续的内存空间来存储具有相同类型的数据
```golang
arr := [100]int{}
```
```
a[k]_address = base_address + k * type_size
m*n 数组， a[i][j] (i<m,j<n)
a[k][j]_address = base_address + (i*n+j)*type_size
```
从0开始，减少CPU开销。数组支持快速下标随机访问。顺序查找和增删查改操作，都需要O(n),最快也要O(log(n))


### 链表
LRU缓存淘汰算法。双向循环链表，单向链表，循环链表，双向链表。