package oop_test

import (
	"fmt"
	"testing"

	"github.com/mrchuanxu/fresh_golang/golangreview/oop"
)


func TestDogMethods(t *testing.T) {
    dog := oop.NewDog("Buddy", 3, "Woof!", 20, 10)
	
	fmt.Println(dog.GetTransMasteredField())
}


func TestSendMessageInter(t *testing.T) {
    err := oop.SendMessage(&oop.Email{
		Recipient: "hello trans",
	},"This is a test message.")
	t.Error(err)
}