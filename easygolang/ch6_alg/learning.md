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

### 栈 一种受限的数据结构
内存中的堆栈和数据结构堆栈不是一个概念，可以说内存中的堆栈是真实存在的物理区，数据结构中的堆栈是抽象的数据存储结构。<br>
内存空间在逻辑上分为三部分：代码区、静态数据区和动态数据区，动态数据区又分为栈区和堆区。 <br>
代码区：存储方法体的二进制代码。高级调度（作业调度）、中级调度（内存调度）、低级调度（进程调度）控制代码区执行代码的切换。 <br>
静态数据区：存储全局变量、静态变量、常量，常量包括final修饰的常量和String常量。系统自动分配和回收。 <br>
栈区：存储运行方法的形参、局部变量、返回值。由系统自动分配和回收。 <br>
堆区：new一个对象的引用或地址存储在栈区，指向该对象存储在堆区中的真实数据。<br>

### 队列 循环队列能够利于CAS
```
tail = (tail+1)%length
tail == head empty
(tail+1) % length == head full
(head+1)%length  proceed
```
循环队列的长度设定需要对并发数据有一定的预测，否则会丢失太多请求。

### 递归 推导公式，找到终止条件，避免嵌套过深。 递归还可能会存在堆栈溢出风险，解决方法是避免嵌套过深。

### 二分查找 [LeetCode题目链接](https://leetcode.com/problems/search-in-rotated-sorted-array/submissions/1667082015/)

### 跳表 一种通过多层索引进行逐级下层查找的时间复杂度为mLogn的存储结构 用于redis

### 散列表 支持随机访问且时间复杂度为O(1)，散列函数比较重要，避免散列冲突，通过装载因子判断散列性能
A % B = A & (B - 1)
hashmap的散列函数公式 return  hash ^ (hash >>> 16) &(capacity - 1)
### 二叉树 顺序存储法 父 i 左2*i 右 2*i+1
### 二叉查找树 左节点<根节点<右节点 O(logn) 频繁修改后容易退化为O(n)
```
// ...existing code...

// 二叉树递归查找，返回找到的节点
func BinaryTreeSearchNode(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == val {
        return root
    }
    if val < root.Val {
        return BinaryTreeSearchNode(root.Left, val)
    }
    return BinaryTreeSearchNode(root.Right, val)
}

// ...existing code...

// ...existing code...

// 二叉查找树删除操作，返回删除后的根节点
func DeleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    if key < root.Val {
        root.Left = DeleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = DeleteNode(root.Right, key)
    } else {
        // 找到要删除的节点
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }
        // 左右子树都不为空，找到右子树最小节点替换
        minNode := root.Right
        for minNode.Left != nil {
            minNode = minNode.Left
        }
        root.Val = minNode.Val
        root.Right = DeleteNode(root.Right, minNode.Val)
    }
    return root
}
```
### 红黑树 近似平衡而非绝对平衡 维持时间复杂度能在 O(logn)
红黑树作为一种近似平衡的二叉查找树，通过围绕关键左右旋以及改变颜色操作将整体二叉查找树维持近似平衡，让整体的增删查改维持在O(logn)，不会退化太严重。
维持规则 1. 根节点是黑色的 2. 每个叶子结点都是黑色的空节点 3. 任何相邻的节点都不能同时为红色，一定是红黑。 4. 每个节点，从该节点到达其可达叶子节点的所有路径，都包含相同数量的黑色节点。



