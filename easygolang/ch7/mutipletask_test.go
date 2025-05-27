package mutipletask_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
	"unsafe"
)

type SingleTon struct {
	data string
}

var signTon *SingleTon
var wg sync.WaitGroup
var once sync.Once


func NewSingleTon() *SingleTon {
	once.Do(func() {
	     signTon = &SingleTon{
		data: "singleton data",}
	})
	return signTon
}


func TestBuildSinglton(t *testing.T){
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(){
			defer wg.Done()
			SingleTon := NewSingleTon()
			t.Logf("singlton pointer %x",unsafe.Pointer(SingleTon))

		}()
	}
	wg.Wait()
}


func runTask(id int)string{
	time.Sleep(10*time.Millisecond)
	return fmt.Sprintf("the result of task %d", id)
}

func FirstResponse()string{
	numOfRunner := 10
	ch := make(chan string,numOfRunner)
	for i := 0;i<numOfRunner;i++{
		go func(i int){
			result := runTask(i)
			ch <- result
		}(i)
	}
	return <-ch
}


func TestFirstResponse(t *testing.T) {
	t.Log("before:",runtime.NumGoroutine())
	t.Log("first response:", FirstResponse())
	time.Sleep(1* time.Millisecond) // 等待其他 goroutine 完成
	t.Log("after:", runtime.NumGoroutine())
	runtime.GC()
}

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}