package try_test

import (
	"strings"
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

func TestArrayInit(t *testing.T){
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr3 := [...]int{1,2,3,4,5}
	arr1[1] = 5
	t.Log(arr[1],arr[2])
	t.Log(arr1,arr3)
}


func TestArrayTravel(t *testing.T){
	arr3:=[...]int{1,3,4,5}
	for i:=0;i<len(arr3);i++{
		t.Log(arr3[i])
	}
	arr := arr3[:1]
	t.Log(arr)
}


func TestSliceTravel(t *testing.T){
	var s0 []int
	t.Log(len(s0),cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0),cap(s0))

	s1 := []int{1,2,3,4}
	t.Log(len(s1),cap(s1))

	s2 := make([]int,3,5)
	t.Log(len(s2),cap(s2))
}


func TestSliceGrowing(t *testing.T){
	s := []int{}
	for i := 0;i<10;i++{
		s = append(s, i)
		t.Log(s,len(s),cap(s))
	}
}

func TestSliceShare(t *testing.T){
	nums := []int{0,1,2,3,4,5,56,6,7}
	s1 := nums[3:5]
	s2 := nums[4:6]
	t.Log(len(s1),cap(s1),s1[0])
	t.Log(len(s2),cap(s2))
	alph := make([]string,3,5)
	t.Log(alph[0])
}


func TestInitMap(t *testing.T){
	m1 := map[int]int{1:2,2:3,3:6}
	t.Log(m1[2])
	for k,v := range m1{
		t.Log(k,v)
	}
}

func TestStringFn(t *testing.T){
	s:="A,B,C"
	parts := strings.Split(s,",")
	for _,part := range parts{
		t.Log(part)
	}
	t.Log(strings.Join(parts,"_"))
}