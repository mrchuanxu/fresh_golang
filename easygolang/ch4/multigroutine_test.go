package ch4_test

import (
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


	select {
	case ret := <- retCh1:
		t.Logf("result %s",&ret)
	
	}
}