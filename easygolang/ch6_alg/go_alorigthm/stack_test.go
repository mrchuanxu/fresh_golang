package goalorigthm_test

import (
	"fmt"
	"testing"

	"github.com/mrchuanxu/vito_infra/alg"
)

func TestStack(t *testing.T){
	ts := alg.Init()
	ts.Push(123)
	ts.Pop()
}


func TestLinkStack(t *testing.T){

	tls := alg.InitLinkStack()
	tls.Push("aht")
	tls.Push("aht")
	tls.Push(3)
	tls.Push("2")
	tls.Push("1")
	fmt.Println(tls.Pop())
	fmt.Println(tls.Pop())
	fmt.Println(tls.Pop())
	fmt.Println(tls.Pop())

	fmt.Println(tls.Pop())
	fmt.Println(tls.Pop())
	fmt.Println(tls.Pop())

}



func TestCircleQue(t *testing.T){
	tqs := alg.InitTransQue(6)
	tqs.EnQue(1)
	tqs.EnQue(2)
	_,err := tqs.EnQue(3)
	if err != nil{
		fmt.Println(err)
	}
	_,err = tqs.EnQue(4)
	if err != nil{
		fmt.Println(err)
	}
	_,err = tqs.EnQue(5)
	if err != nil{
		fmt.Println(err)
	}
	_,err = tqs.EnQue(6)
	if err != nil{
		fmt.Println(err)
	}
	val,err := tqs.DeQue()
	fmt.Println(val,err)
	val,err = tqs.DeQue()
	fmt.Println(val,err)
	val,err = tqs.DeQue()
	fmt.Println(val,err)
	val,err = tqs.DeQue()
	fmt.Println(val,err)
	_,err = tqs.EnQue(5)
	if err != nil{
		fmt.Println(err)
	}
	_,err = tqs.EnQue(6)
	if err != nil{
		fmt.Println(err)
	}
	val,err = tqs.DeQue()
	fmt.Println(val,err)
	val,err = tqs.DeQue()
	fmt.Println(val,err)
	val,err = tqs.DeQue()
	fmt.Println(val,err)
	val,err = tqs.DeQue()
	fmt.Println(val,err)
}


func TestFindBsearch(t *testing.T){
	arr := []int{1,2,3,4,6,7,8,9,22,33}
	loc := alg.BResearch(arr,0,len(arr)-1,9)
	fmt.Println(loc)
}