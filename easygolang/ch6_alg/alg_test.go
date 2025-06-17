package ch6alg_test

import (
	"fmt"
	"testing"

	"github.com/mrchuanxu/vito_infra/alg"
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
	arr := []int{4,1,2,0,3,6,5,7,8,9}
	alg.QuickSort(arr,0,len(arr)-1)
	

	for i:= 0;i<10;i++{
        t.Logf("%d",arr[i])
	}
	nums := []int{4,5,6,7,0,1,2}
	fmt.Println(alg.SearchNums(nums,0))
	// for i:= 0;i<10;i++{
    //     t.Logf("%x",unsafe.Pointer(&arr[i]))
	// }
	// t.Logf("%d",unsafe.Sizeof(1))
}


