package sourcetrans_test

import (
	"testing"

	transpack "github.com/mrchuanxu/fresh_golang/easygolang/trans_pack"
	srcStd "github.com/mrchuanxu/fresh_golang/src/fmt"
)


func TestTransPack(t *testing.T) {
	t.Log(transpack.NewMath())
	srcStd.Println("hello world")
}