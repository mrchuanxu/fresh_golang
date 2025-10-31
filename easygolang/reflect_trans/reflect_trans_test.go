package reflecttrans_test

import (
	"fmt"
	"reflect"
	"testing"
)

var x int = 100

func TestReflect(t *testing.T) {

	reflectType := reflect.TypeOf(x)
	reflectValue := reflect.ValueOf(x)

	fmt.Println(reflectType)
	fmt.Println(reflectValue)
	originalValue := reflectValue.Interface().(int)
	fmt.Println(originalValue)
}


func TestRange(t *testing.T){
	arr := []int{1,2,3,4,5}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	s := "hello"

	for i,v := range s{
		fmt.Println(i,v)
	}
}
