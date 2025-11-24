package oop

import (
	"fmt"
)

func init(){
	fmt.Println("hello oop")
}

type Animal struct {
	Name               string `json:"name"`
	Age                int    `json:"age"`
	transMasteredField string `json:"-"`
}

type AnimalInterface interface {
	Sleep() string
	Eat() string
	Poop() string
}

func (a *Animal) Sleep() string {
	return a.Name + " is sleeping."
}

func (a *Animal) Eat() string {
	return a.Name + " is eating."
}

func (a *Animal) Poop() string {
	return a.Name + " is pooping."
}

func (a *Animal) GetTransMasteredField() string {
	return a.transMasteredField
}

// Dog struct embedding Animal
type Dog struct {
	Animal
	BarkSound          string `json:"bark_sound"`
	RuningSpeed        int    `json:"running_speed"`
	LoyaltyLevel       int    `json:"loyalty_level"`
	transMasteredField string `json:"-"`
}

func (d *Dog) Bark() string {
	return d.Name + " says " + d.BarkSound
}

func (d *Dog) Fetch() string {
	return d.Name + " is fetching the ball!"
}

func (d *Dog) Guard() string {
	return d.Name + " is guarding the house!"
}

func (d *Dog) Sleep() string {
	return d.Name + " the dog is sleeping soundly and longer."
}

type Cat struct {
	Animal
	MeowSound          string `json:"meow_sound"`
	ClimbingSkill      int    `json:"climbing_skill"`
	Independence       int    `json:"independence"`
	transMasteredField string `json:"-"`
}

func (c *Cat) Meow() string {
	return c.Name + " says " + c.MeowSound
}

func (c *Cat) Scratch() string {
	return c.Name + " is scratching the furniture!"
}

func (c *Cat) Purr() string {
	return c.Name + " is purring contentedly!"
}

func NewDog(name string, age int, barkSound string, runningSpeed int, loyaltyLevel int) *Dog {
	return &Dog{
		Animal: Animal{
			Name:               name,
			Age:                age,
			transMasteredField: "trans",
		},
		BarkSound:          barkSound,
		RuningSpeed:        runningSpeed,
		LoyaltyLevel:       loyaltyLevel,
		transMasteredField: "trans",
	}
}

func NewCat(name string, age int, meowSound string, climbingSkill int, independence int) *Cat {
	return &Cat{
		Animal: Animal{
			Name: name,
			Age:  age,
		},
		MeowSound:          meowSound,
		ClimbingSkill:      climbingSkill,
		Independence:       independence,
		transMasteredField: "trans",
	}
}

type MessageInterface interface {
	SendMessage(msg string) error
}

func SendMessage(mi MessageInterface, msg string) error {
	return mi.SendMessage(msg)
}

type Email struct {
	Recipient string
}

func (e *Email) SendMessage(msg string) error {
	// Simulate sending an email
	println("Sending email to", e.Recipient, "with message:", msg)
	return nil
}

func NewEmail(recipient string) *Email {
	return &Email{Recipient: recipient}
}