package ch4_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)



func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			t.Logf("count %d", i)
		}(i)
	}

	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	for i := 0;i<5000;i++{
		wg.Add(1)
		go func(){
			defer func(){
			mut.Unlock()
				}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter %d", counter)

}

func service() string {
	// Simulate some work
	time.Sleep(time.Millisecond * 50)
	return "Service completed"
}

func otherTask() {
	// Simulate some other work
	fmt.Println("Doing other task...")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Other task completed")
}


func AsyncService() chan string{
	retCH := make(chan string)

	go func(){
		ret := service()
		fmt.Println("return result")
		retCH <- ret
		fmt.Println("service exited")
	}()
	return retCH
}

func TestAsync(t *testing.T) {
	retCH := AsyncService()
	fmt.Println("service started")
	ret := <- retCH
	fmt.Println("service returned")
	t.Log(ret)
}


func TestSelecAndSwitch(t *testing.T){
	defer func(p interface{}){
		switch p.(type){
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Println("unknown")
		}
	}(1)
}


func CancelChan(chann chan struct{}) bool{
	select{
	case <- chann:
		return true
	default:
		return false
	}
}


func Cancel_1(cancel chan struct{}){
	cancel <- struct{}{}
}

func Cancel_2(cancel chan struct{}){
	close(cancel)
}


func TestCancelChan(t *testing.T) {
	cancel := make(chan struct{})

	for i := 0; i < 10; i++ {
	    go func(i int,cancel chan struct{}) {

			for{
				if CancelChan(cancel) {
					break
				}
			}
			fmt.Println("current goroutine", i, "exiting")
		}(i,cancel)
	}
	Cancel_2(cancel)
	time.Sleep(time.Second * 1)

}



func isCancelled(ctx context.Context)bool{
	select{
	case <-ctx.Done():
		return true
	default:
		return false	
	}
}

func TestCtxCancel(t *testing.T){
	ctx,cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
	    go func(i int,ctx context.Context) {

			for{
				if isCancelled(ctx) {
					break
				}
			}
			fmt.Println("current goroutine", i, "exiting")
		}(i,ctx)
	}
	cancel()
}