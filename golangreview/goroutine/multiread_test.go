package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMultiRead(t *testing.T) {
	
	go func(){
		fmt.Println("multi read test")
		defer fmt.Println("goroutine deferred")
	}()

	time.Sleep(time.Second)
}


// Select: https://draven.co/golang/docs/part2-foundation/ch05-keyword/golang-select/

func TestDoMultipleWays(t *testing.T){
	
	chanA := make(chan int)
	chanQuit := make(chan int)

	// 斐波那契数列
	go func(chanA chan int, chanQuit chan int){
		x,y := 0,1
		for {
			select{
				case chanA <-x:
					x,y = y,x+y
				case <- chanQuit:
					fmt.Println("quit")
					return 
			}
		}
	}(chanA,chanQuit)

	go func (){
		for i:=0;i<10;i++{
			fmt.Println(<-chanA)
		}
		chanQuit <-0
	}()

	time.Sleep(time.Second)
}

// 好的 现在搞清楚了channel的使用 那么就尝试实现一个消息的分发工作

// 先构建一个生产者 生产信息到channel中

func TestProducerConsumer(t *testing.T){
	var wg sync.WaitGroup

	channelMsg := make(chan string,20)
	wg.Add(1)
	go func(){
		defer wg.Done()
		msgs := []string{"hello","world","this","is","a","test","for","producer","and","consumer"}
		index := 0
		for {
			if index >= len(msgs){
				close(channelMsg)
				fmt.Println("producer finished all messages")
				return
			}
			channelMsg <- msgs[index % len(msgs)]
			index++
			}
	}() 
	// 启动消费者
	consumerCount := 2
	wg.Add(consumerCount)
	
	for i:=0;i<consumerCount;i++{
		go func(id int){
			defer wg.Done()
			for {
				select{
					case msg,ok := <- channelMsg:
						if !ok {
							fmt.Printf("consumer %d channel closed, quitting\n", id)
							return
						}
						fmt.Printf("consumer %d received message: %s\n",id,msg)
					case <- time.After(2*time.Second):
						fmt.Printf("consumer %d timeout, quitting\n",id)
						return
				}
			}
		}(i)
	}
	fmt.Println("waiting for all workers to finish")
	wg.Wait()
}


func TestChannelStop(t *testing.T){
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Println("sorry stop now")
		ch <- 1
		fmt.Println("after send")
	}()
	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Printf("num is %d",<-ch)
	}()
	wg.Wait()
	time.Sleep(time.Second)
}


func TestConstants(t *testing.T){
	const str string = "hello constants"
	fmt.Println(str)

	if num := 9;num < 0{
		fmt.Println("num is negative")
	}else if num < 10 {
		fmt.Println("num is a single digit")
	}else{
		fmt.Println("num is multi digit")
	}
}


func TestSwitch(t *testing.T){
    switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	var str interface{} = "golang"

	switch str.(type){
	case string:
		fmt.Println("str is a string")
	case int:
		fmt.Println("str is an int")
	default:
		fmt.Println("unknown type")
	}

	switch {
	case 1 < 0:
		fmt.Println("the world is upside down")
	case 2 > 1:
		fmt.Println("all is normal")
	default:
		fmt.Println("this should never happen")
	}
}

func TestArrays(t *testing.T){
	var arr [5]int

	for i:=0;i<len(arr);i++{
		arr[i] = i
	}
	fmt.Println(arr)

	arr1 := [...]int{1,2,3,4,5:6,7,78,8,8}
	fmt.Println(arr1)

	var twoD [2][3]int

	for i := range 2{
		for j := range 3{
			twoD[i][j] = i+j
		}
	}
	fmt.Println(twoD)
}


func TestSlices(t *testing.T){
	s := make([]string,3)
	s[0] = "my"
	s[1] = "trans"
	s[2] = "god"

	c := make([]string,len(s))

	copy(c,s) // copy没有扩容机制
	fmt.Println("cpy:",c)
}


func TestMaps(t *testing.T){
	// map具有随机性 遍历无法顺序获取
	m := make(map[string]int)
	fmt.Println(m)
	m["s"] = 1
	m["b"] = 2

	delete(m,"s")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)
} 


func TestClosures(t *testing.T){
	fmt.Println(intSeq()())

	fmt.Println(fact(7))

	var fib func(int) int
	// 斐波那契
	fib = func(n int) int{
		if n <= 2{
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(5))
}


func intSeq()func()int{
		i := 0
		return func()int{
			i++
			return i
		}
}


func fact(n int) int{
	// 终止条件
	if n == 0{
		return 1
	}

	return n*fact(n - 1)
}