package ch8_test

import (
	"fmt"
	"reflect"
	"testing"
)


func TestTypeVal(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f)) // int64
	t.Log(reflect.ValueOf(f)) // 10	
	t.Log(reflect.ValueOf(f).Type()) // 10
}

func CheckType(v interface{}){
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32,reflect.Float64:
		fmt.Println("Float")
	case reflect.Int,reflect.Int32,reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("unknown")
	}
}

func TestBasicType(t *testing.T){
	var f float32 = 43.33
	CheckType(f)

}