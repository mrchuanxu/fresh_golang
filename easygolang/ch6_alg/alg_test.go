package ch6alg_test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestAlgTwo(t *testing.T) {
	i :=1
	for i < 10 {
		i = i * 2
		fmt.Println(i)
	}


	arr := [100]int{}
    for i:= 0;i<100;i++{
        arr[i] = i*i
    }

	for i:= 100-1;i>=0;i--{
		fmt.Println(arr[i])
	}	
}


func TestArray(t *testing.T){
	arr := [10]int{0,1,2,3,4,5,6,7,8,9}
	for i:= 0;i<10;i++{
        t.Logf("%x",unsafe.Pointer(&arr[i]))
	}
	t.Logf("%d",unsafe.Sizeof(1))
}