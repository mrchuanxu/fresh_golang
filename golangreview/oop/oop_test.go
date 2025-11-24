package oop_test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/mrchuanxu/fresh_golang/golangreview/oop"
)


func TestDogMethods(t *testing.T) {
    dog := oop.NewDog("Buddy", 3, "Woof!", 20, 10)
	
	fmt.Println(dog.GetTransMasteredField())
}


func TestSendMessageInter(t *testing.T) {
    err := oop.SendMessage(&oop.Email{
		Recipient: "hello trans",
	},"This is a test message.")
	t.Error(err)
}


func TestMultiple(t *testing.T){
	MultipleRun()
}

func RunPrint(name string){
	fmt.Printf("%s",name)
}


func MultipleRun()error{
	chanA,chanB,chanC := make(chan struct{}),make(chan struct{}),make(chan struct{})
	ctx,cancel := context.WithCancel(context.Background())

	var totalRun int32 = 300

	var wg sync.WaitGroup

	wg.Add(3)
	go singleRun(ctx,chanA,chanB,&totalRun,&wg,"A",cancel)
	go singleRun(ctx,chanB,chanC,&totalRun,&wg,"B",cancel)
	go singleRun(ctx,chanC,chanA,&totalRun,&wg,"C",cancel)
	chanA<-struct{}{}
	wg.Wait()
	return nil
}


func singleRun(ctx context.Context,in,out chan struct{},totalRun *int32,wg *sync.WaitGroup,name string,cancel context.CancelFunc){
	defer wg.Done()
	for {
		select{
		case <- ctx.Done():
			return
		case _,ok:= <- in:
			if !ok{
				return
			}
			RunPrint(name)
			if atomic.AddInt32(totalRun,-1) == 0{
				cancel()
				return
			}
			select{
			case <- ctx.Done():
				return
			case out <- struct{}{}:
			}
		}
	}
}

func TestContextChild(t *testing.T){
	ctx1,_ := context.WithCancel(context.Background())
	ctx2,cancel1 := context.WithCancel(ctx1)
	ctx3,_ := context.WithCancel(ctx2)

	var wg sync.WaitGroup
	wg.Add(3)
	go func(ctx context.Context){
		defer wg.Done()
		fmt.Printf("%s","ctx1")
		select{
		case <- ctx.Done():
			fmt.Println("ctx1 done")
			return
		}
	}(ctx1)

	go func(ctx context.Context){
		defer wg.Done()
		fmt.Printf("%s","ctx2")
		select{
		case <- ctx.Done():
			fmt.Println("ctx2 done")
			return
		}
	}(ctx2)

	go func(ctx context.Context){
		defer wg.Done()
		fmt.Printf("%s","ctx3")
		select{
		case <- ctx.Done():
			fmt.Println("ctx3 done")
			return
		}
	}(ctx3)
	cancel1()
	wg.Wait()
}


// 电梯调度系统
func TestElevator(t *testing.T){

}


// 获取电梯不允许其他资源占用 一共两台电梯
type Elevator struct{
	Mux sync.Mutex
	sumPeople uint64
	
}

// People 12
func UpPeople(e *Elevator,upNum uint64)error{
	e.Mux.Lock()
	defer e.Mux.Unlock()
	// 如果满员就报错
	if atomic.AddUint64(&e.sumPeople,upNum) >= 12{
		return errors.New("full people")
	}

	return nil
}

func DownPeople(e *Elevator,upNum uint64)error{
	e.Mux.Lock()
	defer e.Mux.Unlock()
	atomic.AddUint64(&e.sumPeople,-upNum)
	return nil
}


type QNode struct{
	Num int
	Next *QNode
	Pre *QNode
}

type TQueue struct{
	mux *sync.Mutex
	Head *QNode
	Tail *QNode
}

func InitQue()*TQueue{
	return &TQueue{
		mux: &sync.Mutex{},
	}
}


func (t *TQueue) InQue(node *QNode){
	defer t.mux.Unlock()
	t.mux.Lock()
	if t.Head == nil{
		t.Head = node
	}
	if t.Tail == nil{
		t.Tail = node
		return
	}
	
	t.Tail.Next = node
	t.Tail = t.Tail.Next
}

func (t *TQueue) DeQue()*QNode{
	defer t.mux.Unlock()
	t.mux.Lock()
	if t.Head != nil{
		deNode := t.Head
		t.Head = t.Head.Next
		return deNode
	}
	return nil
}

func(t *TQueue)HeadQue()*QNode{
	return t.Head
}

func TestQuemulti(t *testing.T){
	tQue := InitQue()
	for i:= 0;i<30;i++{
		go tQue.InQue(&QNode{
			Num: i,
		})
	}

	time.Sleep(5*time.Second)

	for tQue.Head != nil{
		fmt.Println(tQue.DeQue().Num)
	}
}


// 模拟内存 出发gc