package ch6alg_test

import (
	"fmt"
	"testing"
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