package try_test

import (
	"testing"
)

func TestFirstTry(t *testing.T) {
	b := "c"
	t.Logf("%s",b)
}

// go 不支持隐式类型数据转换

func TestImplicit(t *testing.T){
	// var a int32 = 1
	// var b int64
	// b = a
}

// go 不支持指针运算
func TestPoint(t *testing.T){
	//a:=1
	//aPtr := &a
	//aPtr = aPtr + 1
	
}

// go 运算符
func TestCompareArray(t *testing.T){
	a := [...]int{1,2,3,4}
	b := [...]int{1,2,4,3}
	d := [...]int{1,2,3,4}
	t.Log(a==b,a==d)
}

func TestSwitch(t *testing.T){
	for i:=0;i<5;i++{
		switch i{
		case 0,2:
			t.Log("Even")
		case 1,3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
	// switch os:=runtime.GOOS; os{
	// 	case "d"
	// }
}