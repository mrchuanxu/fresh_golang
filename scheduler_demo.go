package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 演示Go调度器的工作原理
func schedulerDemo() {
	fmt.Println("=== Go调度器工作原理演示 ===")
	
	// 1. 查看系统信息
	fmt.Printf("CPU核心数: %d\n", runtime.NumCPU())
	fmt.Printf("当前GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	
	// 2. 演示goroutine的创建和调度
	var wg sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d 开始执行\n", id)
			time.Sleep(time.Millisecond * 100)
			fmt.Printf("Goroutine %d 执行完成\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("所有goroutine执行完成")
}

// 演示channel的通信机制
func channelDemo() {
	fmt.Println("\n=== Channel通信机制演示 ===")
	
	// 1. 无缓冲channel - 同步通信
	fmt.Println("1. 无缓冲channel演示:")
	ch := make(chan string)
	
	go func() {
		fmt.Println("  发送方: 准备发送数据...")
		time.Sleep(time.Millisecond * 500)
		ch <- "Hello from sender"
		fmt.Println("  发送方: 数据已发送")
	}()
	
	fmt.Println("  接收方: 等待接收数据...")
	msg := <-ch
	fmt.Printf("  接收方: 收到数据 '%s'\n", msg)
	
	// 2. 有缓冲channel - 异步通信
	fmt.Println("\n2. 有缓冲channel演示:")
	bufferedCh := make(chan string, 3)
	
	// 发送方可以连续发送多个数据
	go func() {
		for i := 0; i < 3; i++ {
			bufferedCh <- fmt.Sprintf("Message %d", i)
			fmt.Printf("  发送方: 发送 Message %d\n", i)
		}
	}()
	
	// 接收方可以连续接收数据
	for i := 0; i < 3; i++ {
		msg := <-bufferedCh
		fmt.Printf("  接收方: 收到 %s\n", msg)
	}
}

// 演示并发安全
func concurrencySafetyDemo() {
	fmt.Println("\n=== 并发安全演示 ===")
	
	// 1. 不安全的并发访问
	fmt.Println("1. 不安全的并发访问:")
	var unsafeCounter int
	var wg sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			unsafeCounter++ // 竞态条件
		}()
	}
	
	wg.Wait()
	fmt.Printf("   不安全计数器结果: %d (期望: 1000)\n", unsafeCounter)
	
	// 2. 使用互斥锁保证安全
	fmt.Println("\n2. 使用互斥锁的并发访问:")
	var safeCounter int
	var mutex sync.Mutex
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			safeCounter++
			mutex.Unlock()
		}()
	}
	
	wg.Wait()
	fmt.Printf("   安全计数器结果: %d (期望: 1000)\n", safeCounter)
}

// 演示select语句
func selectDemo() {
	fmt.Println("\n=== Select语句演示 ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// 启动两个goroutine分别向不同channel发送数据
	go func() {
		time.Sleep(time.Millisecond * 200)
		ch1 <- "来自channel1的消息"
	}()
	
	go func() {
		time.Sleep(time.Millisecond * 100)
		ch2 <- "来自channel2的消息"
	}()
	
	// 使用select监听多个channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("收到: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("收到: %s\n", msg2)
		case <-time.After(time.Second):
			fmt.Println("超时")
		}
	}
}

// 演示goroutine泄漏检测
func goroutineLeakDemo() {
	fmt.Println("\n=== Goroutine泄漏检测 ===")
	
	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("初始goroutine数量: %d\n", initialGoroutines)
	
	// 启动一些goroutine
	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(time.Millisecond * 100)
		}()
	}
	
	time.Sleep(time.Millisecond * 200)
	currentGoroutines := runtime.NumGoroutine()
	fmt.Printf("当前goroutine数量: %d\n", currentGoroutines)
	fmt.Printf("活跃goroutine数量: %d\n", currentGoroutines-initialGoroutines)
}

func main() {
	fmt.Println("Go语言并发机制深度解析")
	fmt.Println("========================")
	
	// 运行各种演示
	schedulerDemo()
	channelDemo()
	concurrencySafetyDemo()
	selectDemo()
	goroutineLeakDemo()
	
	fmt.Println("\n=== 总结 ===")
	fmt.Println("Go语言的并发机制基于以下核心概念:")
	fmt.Println("1. Goroutine: 轻量级用户态线程")
	fmt.Println("2. Channel: goroutine间的通信机制")
	fmt.Println("3. Select: 多路复用机制")
	fmt.Println("4. Sync包: 同步原语(Mutex, WaitGroup等)")
	fmt.Println("5. 调度器: M:N调度模型")
} 