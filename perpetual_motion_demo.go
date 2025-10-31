package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// 基础无限循环永动机
func basicPerpetualMotion() {
	fmt.Println("=== 基础无限循环永动机 ===")
	for {
		fmt.Println("基础永动机运行中...")
		time.Sleep(time.Second)
	}
}

// Goroutine协作永动机
func goroutinePerpetualMotion() {
	fmt.Println("=== Goroutine协作永动机 ===")
	
	var wg sync.WaitGroup
	
	// 启动多个worker goroutine
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				fmt.Printf("Worker %d 正在工作...\n", id)
				time.Sleep(time.Millisecond * 500)
			}
		}(i)
	}
	
	// 主goroutine监控
	for {
		fmt.Printf("主程序监控中... Goroutine数量: %d\n", runtime.NumGoroutine())
		time.Sleep(time.Second)
	}
}

// Channel永动机
func channelPerpetualMotion() {
	fmt.Println("=== Channel永动机 ===")
	
	// 创建有缓冲的环形channel
	ch := make(chan int, 5)
	
	// 生产者
	go func() {
		counter := 0
		for {
			ch <- counter
			fmt.Printf("生产: %d\n", counter)
			counter++
			time.Sleep(time.Millisecond * 300)
		}
	}()
	
	// 消费者
	go func() {
		for {
			value := <-ch
			fmt.Printf("消费: %d\n", value)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	
	// 主程序监控
	for {
		time.Sleep(time.Second)
		fmt.Printf("环形Channel中数据量: %d\n", len(ch))
	}
}

// Select多路复用永动机
func selectPerpetualMotion() {
	fmt.Println("=== Select多路复用永动机 ===")
	
	ch1 := make(chan string)
	ch2 := make(chan int)
	ch3 := make(chan bool)
	
	// 随机发送数据到不同channel
	go func() {
		for {
			switch rand.Intn(3) {
			case 0:
				ch1 <- "字符串数据"
			case 1:
				ch2 <- rand.Intn(100)
			case 2:
				ch3 <- rand.Float64() > 0.5
			}
			time.Sleep(time.Millisecond * 200)
		}
	}()
	
	// 使用select监听多个channel
	for {
		select {
		case msg := <-ch1:
			fmt.Printf("收到字符串: %s\n", msg)
		case num := <-ch2:
			fmt.Printf("收到数字: %d\n", num)
		case flag := <-ch3:
			fmt.Printf("收到布尔值: %t\n", flag)
		case <-time.After(time.Second):
			fmt.Println("超时，继续运行...")
		}
	}
}

// Ticker定时器永动机
func tickerPerpetualMotion() {
	fmt.Println("=== Ticker定时器永动机 ===")
	
	// 创建定时器
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()
	
	counter := 0
	for range ticker.C {
		fmt.Printf("Ticker永动机运行中... 计数: %d\n", counter)
		counter++
	}
}

// 系统监控永动机
func systemMonitorPerpetualMotion() {
	fmt.Println("=== 系统监控永动机 ===")
	
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		
		fmt.Printf("=== 系统监控 ===\n")
		fmt.Printf("Goroutine数量: %d\n", runtime.NumGoroutine())
		fmt.Printf("内存使用: %d MB\n", m.Alloc/(1024*1024))
		fmt.Printf("系统内存: %d MB\n", m.Sys/(1024*1024))
		fmt.Printf("GC次数: %d\n", m.NumGC)
		fmt.Printf("时间: %s\n", time.Now().Format(time.RFC3339))
		fmt.Println("================")
		
		time.Sleep(time.Second)
	}
}

// 优雅退出永动机
func gracefulPerpetualMotion() {
	fmt.Println("=== 优雅退出永动机 ===")
	fmt.Println("按 Ctrl+C 退出程序")
	
	// 创建退出信号通道
	quit := make(chan bool, 1)
	
	// 启动goroutine监听退出信号
	go func() {
		fmt.Println("按回车键退出...")
		fmt.Scanln()
		quit <- true
	}()
	
	// 主循环
	counter := 0
	for {
		select {
		case <-quit:
			fmt.Println("收到退出信号，优雅关闭...")
			return
		default:
			fmt.Printf("优雅永动机运行中... 计数: %d\n", counter)
			counter++
			time.Sleep(time.Second)
		}
	}
}

// 递归永动机
func recursivePerpetualMotion(counter int) {
	fmt.Printf("递归深度: %d\n", counter)
	time.Sleep(time.Millisecond * 100)
	
	// 递归调用自身
	recursivePerpetualMotion(counter + 1)
}

// 内存分配永动机
func memoryAllocationPerpetualMotion() {
	fmt.Println("=== 内存分配永动机 ===")
	
	var m runtime.MemStats
	
	for {
		// 分配内存
		data := make([]byte, 1024*1024) // 1MB
		
		// 获取内存统计
		runtime.ReadMemStats(&m)
		fmt.Printf("分配内存: %d MB, 总内存: %d MB\n", 
			len(data)/(1024*1024), m.Alloc/(1024*1024))
		
		time.Sleep(time.Millisecond * 100)
	}
}

// 主函数 - 选择运行不同的永动机
func main() {
	fmt.Println("Go语言循环永动机演示")
	fmt.Println("=====================")
	fmt.Println("请选择要运行的永动机类型:")
	fmt.Println("1. 基础无限循环")
	fmt.Println("2. Goroutine协作")
	fmt.Println("3. Channel通信")
	fmt.Println("4. Select多路复用")
	fmt.Println("5. Ticker定时器")
	fmt.Println("6. 系统监控")
	fmt.Println("7. 优雅退出")
	fmt.Println("8. 递归永动机")
	fmt.Println("9. 内存分配")
	fmt.Println("0. 退出程序")
	
	var choice int
	fmt.Print("请输入选择 (0-9): ")
	fmt.Scanln(&choice)
	
	switch choice {
	case 1:
		basicPerpetualMotion()
	case 2:
		goroutinePerpetualMotion()
	case 3:
		channelPerpetualMotion()
	case 4:
		selectPerpetualMotion()
	case 5:
		tickerPerpetualMotion()
	case 6:
		systemMonitorPerpetualMotion()
	case 7:
		gracefulPerpetualMotion()
	case 8:
		fmt.Println("=== 递归永动机 ===")
		recursivePerpetualMotion(0)
	case 9:
		memoryAllocationPerpetualMotion()
	case 0:
		fmt.Println("程序退出")
		return
	default:
		fmt.Println("无效选择，运行基础永动机")
		basicPerpetualMotion()
	}
} 