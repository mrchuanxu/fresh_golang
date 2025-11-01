package goroutine_test

import (
	"fmt"
	"iter"
	"sync"
	"testing"
	"time"
	"unicode/utf8"
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

func fib(i int) int{
	if i <=2{
		return i
	}

	return fib(i-1) + fib(i-2)
}


func TestPointers(t *testing.T){
	var iptr *int
	beptr := 9
	iptr = &beptr
	fmt.Println(*iptr)
	*iptr = 10
	fmt.Println(beptr)
}


func TestRune(t *testing.T){
	const s = "สวัสดี"

	fmt.Println("Len: ",len(s))

	for i := 0;i<len(s);i++{
		fmt.Printf("%x",s[i])
	}

	fmt.Println("rune count",utf8.RuneCountInString(s))


	const ss = "hello，我的朋友trans"

	// bytes := []byte(ss)

	// for i,b := range bytes{
	// 	fmt.Printf("字节 %d:%08b (0x%02x) -> %q\n",i,b,b,b)
	// }

	runes := []rune(ss)

	fmt.Println(runes)

	for i, r := range runes {
        fmt.Printf("  字符 %d: %c (U+%04X)\n", i, r, r)
        // 显示该字符的UTF-8编码
        charBytes := []byte(string(r))
        fmt.Printf("        UTF-8编码: % x\n", charBytes)
    }
}


func TestGeneric(t *testing.T){
	type MySlice []int

	myslice := MySlice{1,2,3,4,5}
	SlicesIndex(myslice,3)

}

func SlicesIndex[S ~[]E,E comparable](s S,v E)int{
	for i := range s{
		if v == s[i]{
			return i
		}
	}
	return -1
}

type List[T any] struct{
	head,tail *element[T]
}

type element[T any] struct{
	next *element[T]
	val T
}

func (lst *List[T]) Push(v T){
	if lst.tail == nil{
		lst.head = &element[T]{val:v}
		lst.tail = lst.head
	}else{
		lst.tail.next = &element[T]{val:v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) Allelements()[]T{
	var elems []T
	for e := lst.head;e!= nil;e = e.next{
		elems = append(elems, e.val)
	}
	return elems
}

func (lst *List[T])All()iter.Seq[T]{
	return func(yield func(T) bool){
		for e := lst.head;e!= nil;e = e.next{
			if !yield(e.val){
				return
			}
		}
	}
}


func genFib() iter.Seq[int]{
	return func(yield func(int) bool){
		a,b := 1,1
		for {
			if !yield(a){
				return
			}
			a,b = b, a+b
		}
	}
}


func TestGenericElements(t *testing.T){
	var s = []string{"hello","world","network"}

	lst := List[string]{}

	for i:=0;i<len(s);i++{
		lst.Push(s[i])
	}

	fmt.Println(lst.Allelements())

	for str := range lst.IterAll(){
		// if str == "world"{
		// 	fmt.Println("exist world stop")
		// 	break
		// }
		fmt.Println(str)
	}

}

// 避免内存的浪费
func (lst *List[T])IterAll()iter.Seq[T]{
	return func(yeild func(T)bool){
		if lst == nil{
			return
		}
		for e := lst.head;e != nil;e = e.next{
			if !yeild(e.val){
				return
			}
		}
	}
}


func TestChannelMsgs(t *testing.T){
	msgs := make(chan string,2)

	msgs <- "buffed"
	msgs <- "channle"

	go func(){
		for {
			fmt.Println("wait")
			val := <- msgs
		    msgs <- "str"
			fmt.Println(val,"wait done")
			time.Sleep(5 * time.Second)
			break

	    }
	}()
	time.Sleep(10*time.Second)
}


func TestSelectChannel(t *testing.T){
	msgs := make(chan string)


	msg := "hi"

	go func(){
		fmt.Println(<-msgs)
	}()
    time.Sleep(time.Second)
	select{
	case msgs <- msg:
		fmt.Println("msg recevied")
	default:
		fmt.Println("msg not recevied")
	}
	time.Sleep(2* time.Second)

}

func TestChannelTry(t *testing.T){
	queue := make(chan string,2)

	queue <- "one"
	queue <- "two"
	// go func(){
	// 	ns := time.Now().Second()
	// 	for {
	// 		queue<-"one"
	// 		if time.Now().Second() - ns > 2{
	// 			close(queue)
	// 		}
	// 	}
	// }()
	close(queue)
	for elem := range queue{
		fmt.Println(elem)
	}
}



func TestTicker(t *testing.T){
	timeTicker := time.NewTicker(3 * time.Second)

	doneCh := make(chan int)

	go func(){
		for{
			select{
			case <-doneCh:
				return
			case t:=<-timeTicker.C:
				fmt.Println("timeTicker:",t)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	timeTicker.Stop()
	doneCh<-1

}