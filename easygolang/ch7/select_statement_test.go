package mutipletask_test

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// TestBasicSelect 演示select的基本用法
func TestBasicSelect(t *testing.T) {
	t.Log("=== Select 基本用法演示 ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// 启动两个goroutine分别向不同channel发送数据
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自channel1的消息"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "来自channel2的消息"
	}()
	
	// 使用select监听多个channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			t.Logf("收到ch1: %s", msg1)
		case msg2 := <-ch2:
			t.Logf("收到ch2: %s", msg2)
		case <-time.After(300 * time.Millisecond):
			t.Log("超时")
		}
	}
}

// TestSelectWithDefault 演示带default的select（非阻塞）
func TestSelectWithDefault(t *testing.T) {
	t.Log("=== Select Default 用法演示 ===")
	
	ch := make(chan string)
	
	// 非阻塞读取
	select {
	case msg := <-ch:
		t.Logf("收到消息: %s", msg)
	default:
		t.Log("channel为空，执行默认操作")
	}
	
	// 非阻塞写入
	select {
	case ch <- "测试消息":
		t.Log("成功发送消息")
	default:
		t.Log("channel已满，无法发送")
	}
}

// TestSelectPriority 演示select的随机性（无优先级）
func TestSelectPriority(t *testing.T) {
	t.Log("=== Select 随机性演示 ===")
	
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	
	// 两个channel都有数据
	ch1 <- "channel1数据"
	ch2 <- "channel2数据"
	
	// 多次测试，观察选择的随机性
	results := make(map[string]int)
	for i := 0; i < 10; i++ {
		ch1 <- fmt.Sprintf("ch1-数据%d", i)
		ch2 <- fmt.Sprintf("ch2-数据%d", i)
		
		select {
		case msg := <-ch1:
			results["ch1"]++
			t.Logf("选择了ch1: %s", msg)
		case msg := <-ch2:
			results["ch2"]++
			t.Logf("选择了ch2: %s", msg)
		}
	}
	
	t.Logf("选择统计: ch1=%d次, ch2=%d次", results["ch1"], results["ch2"])
}

// TestSelectTimeout 演示select超时处理
func TestSelectTimeout(t *testing.T) {
	t.Log("=== Select 超时处理演示 ===")
	
	ch := make(chan string)
	
	// 模拟超时场景
	go func() {
		time.Sleep(2 * time.Second) // 延迟发送
		ch <- "延迟的消息"
	}()
	
	select {
	case msg := <-ch:
		t.Logf("收到消息: %s", msg)
	case <-time.After(1 * time.Second):
		t.Log("操作超时")
	}
}

// TestSelectContext 演示使用context取消
func TestSelectContext(t *testing.T) {
	t.Log("=== Select Context 取消演示 ===")
	
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	
	ch := make(chan string)
	
	// 启动一个可能很慢的操作
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "慢操作完成"
	}()
	
	select {
	case msg := <-ch:
		t.Logf("操作完成: %s", msg)
	case <-ctx.Done():
		t.Logf("操作被取消: %v", ctx.Err())
	}
}

// TestSelectWorkerPool 演示select在工作池中的应用
func TestSelectWorkerPool(t *testing.T) {
	t.Log("=== Select 工作池演示 ===")
	
	const numWorkers = 3
	const numJobs = 10
	
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)
	done := make(chan bool)
	
	// 启动workers
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for {
				select {
				case job := <-jobs:
					// 模拟工作
					time.Sleep(100 * time.Millisecond)
					result := fmt.Sprintf("Worker%d处理Job%d", workerID, job)
					results <- result
				case <-done:
					t.Logf("Worker%d停止", workerID)
					return
				}
			}
		}(i)
	}
	
	// 发送任务
	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	
	// 收集结果
	for i := 0; i < numJobs; i++ {
		result := <-results
		t.Log(result)
	}
	
	// 停止所有workers
	close(done)
	time.Sleep(200 * time.Millisecond)
}

// TestSelectFanInPattern 演示fan-in模式
func TestSelectFanInPattern(t *testing.T) {
	t.Log("=== Select Fan-in 模式演示 ===")
	
	// 创建多个输入channel
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	
	// 启动多个生产者
	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- fmt.Sprintf("生产者1-数据%d", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(ch1)
	}()
	
	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- fmt.Sprintf("生产者2-数据%d", i)
			time.Sleep(150 * time.Millisecond)
		}
		close(ch2)
	}()
	
	go func() {
		for i := 0; i < 3; i++ {
			ch3 <- fmt.Sprintf("生产者3-数据%d", i)
			time.Sleep(200 * time.Millisecond)
		}
		close(ch3)
	}()
	
	// Fan-in: 合并多个channel
	openChannels := 3
	for openChannels > 0 {
		select {
		case msg, ok := <-ch1:
			if !ok {
				ch1 = nil
				openChannels--
				t.Log("ch1已关闭")
			} else {
				t.Logf("从ch1收到: %s", msg)
			}
		case msg, ok := <-ch2:
			if !ok {
				ch2 = nil
				openChannels--
				t.Log("ch2已关闭")
			} else {
				t.Logf("从ch2收到: %s", msg)
			}
		case msg, ok := <-ch3:
			if !ok {
				ch3 = nil
				openChannels--
				t.Log("ch3已关闭")
			} else {
				t.Logf("从ch3收到: %s", msg)
			}
		}
	}
}

// TestSelectRaceCondition 演示避免竞态条件
func TestSelectRaceCondition(t *testing.T) {
	t.Log("=== Select 避免竞态条件演示 ===")
	
	var counter int
	var mu sync.Mutex
	incrementCh := make(chan bool)
	decrementCh := make(chan bool)
	resultCh := make(chan int)
	done := make(chan bool)
	
	// 安全的计数器操作
	go func() {
		for {
			select {
			case <-incrementCh:
				mu.Lock()
				counter++
				mu.Unlock()
				t.Log("计数器递增")
			case <-decrementCh:
				mu.Lock()
				counter--
				mu.Unlock()
				t.Log("计数器递减")
			case resultCh <- func() int {
				mu.Lock()
				defer mu.Unlock()
				return counter
			}():
				// 发送当前值
			case <-done:
				return
			}
		}
	}()
	
	// 并发操作
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementCh <- true
			time.Sleep(10 * time.Millisecond)
			decrementCh <- true
		}()
	}
	
	wg.Wait()
	
	// 获取最终结果
	finalResult := <-resultCh
	t.Logf("最终计数器值: %d", finalResult)
	
	close(done)
}

// TestSelectMemoryLeak 演示避免内存泄漏
func TestSelectMemoryLeak(t *testing.T) {
	t.Log("=== Select 内存泄漏避免演示 ===")
	
	initialGoroutines := runtime.NumGoroutine()
	t.Logf("初始goroutine数量: %d", initialGoroutines)
	
	done := make(chan bool)
	
	// 启动多个goroutine，但有正确的退出机制
	for i := 0; i < 5; i++ {
		go func(id int) {
			ch := make(chan string)
			
			// 模拟可能永远不会收到数据的channel
			go func() {
				time.Sleep(time.Duration(id) * 100 * time.Millisecond)
				if id < 3 { // 只有部分goroutine会发送数据
					ch <- fmt.Sprintf("数据来自goroutine %d", id)
				}
			}()
			
			select {
			case msg := <-ch:
				t.Logf("收到消息: %s", msg)
			case <-done: // 确保有退出机制
				t.Logf("Goroutine %d 正常退出", id)
				return
			case <-time.After(300 * time.Millisecond):
				t.Logf("Goroutine %d 超时退出", id)
				return
			}
		}(i)
	}
	
	time.Sleep(400 * time.Millisecond)
	close(done) // 通知所有goroutine退出
	time.Sleep(100 * time.Millisecond)
	
	finalGoroutines := runtime.NumGoroutine()
	t.Logf("最终goroutine数量: %d", finalGoroutines)
	t.Logf("goroutine增量: %d", finalGoroutines-initialGoroutines)
}

// TestSelectPerformance 演示select性能特点
func TestSelectPerformance(t *testing.T) {
	t.Log("=== Select 性能测试演示 ===")
	
	const iterations = 100000
	
	// 测试单channel的性能
	t.Log("单channel性能测试:")
	start := time.Now()
	ch := make(chan int)
	go func() {
		for i := 0; i < iterations; i++ {
			ch <- i
		}
		close(ch)
	}()
	
	count := 0
	for range ch {
		count++
	}
	singleChannelTime := time.Since(start)
	t.Logf("单channel处理%d条消息用时: %v", count, singleChannelTime)
	
	// 测试select多channel的性能
	t.Log("多channel select性能测试:")
	start = time.Now()
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan bool)
	
	go func() {
		for i := 0; i < iterations/2; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	
	go func() {
		for i := 0; i < iterations/2; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	
	count = 0
	openChannels := 2
	for openChannels > 0 {
		select {
		case _, ok := <-ch1:
			if !ok {
				ch1 = nil
				openChannels--
			} else {
				count++
			}
		case _, ok := <-ch2:
			if !ok {
				ch2 = nil
				openChannels--
			} else {
				count++
			}
		case <-done:
			return
		}
	}
	
	multiChannelTime := time.Since(start)
	t.Logf("多channel select处理%d条消息用时: %v", count, multiChannelTime)
	t.Logf("性能比较: 单channel/多channel = %.2f", 
		float64(singleChannelTime)/float64(multiChannelTime))
} 