package sourcetrans_test

import (
	"fmt"
	"testing"

	transpack "github.com/mrchuanxu/fresh_golang/easygolang/trans_pack"
)


func TestTransPack(t *testing.T) { 
	t.Log(transpack.NewMath())
	fmt.Println("hello")
}