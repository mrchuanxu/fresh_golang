package main

import (
	"fmt"
	"testing"
	"unsafe"
)


type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    float64
}

func NewEmployee(id int, firstName, lastName, position string, salary float64) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
	}
}

func (e *Employee) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
	fmt.Printf("Address is %x", unsafe.Pointer(e))
}	

func TestNewEmployeePointer(t *testing.T) {
	emp := NewEmployee(1, "John", "Doe", "Developer", 75000.0)
	empPtr := &emp

	if empPtr.ID != 1 {
		t.Errorf("Expected ID 1, got %d", empPtr.ID)
	}
	if empPtr.FirstName != "John" {
		t.Errorf("Expected FirstName 'John', got '%s'", empPtr.FirstName)
	}
	if empPtr.LastName != "Doe" {
		t.Errorf("Expected LastName 'Doe', got '%s'", empPtr.LastName)
	}
	if empPtr.Position != "Developer" {
		t.Errorf("Expected Position 'Developer', got '%s'", empPtr.Position)
	}
	if empPtr.Salary != 75000.0 {
		t.Errorf("Expected Salary 75000.0, got %f", empPtr.Salary)
	}

	t.Logf("Employee pointer: %+v", empPtr)
}

func TestEmployeeSalaryUpdate(t *testing.T) {
	emp := NewEmployee(2, "Alice", "Smith", "Manager", 90000.0)
	empPtr := &emp

	empPtr.Salary = 95000.0

	if empPtr.Salary != 95000.0 {
		t.Errorf("Expected updated Salary 95000.0, got %f", empPtr.Salary)
	}

	t.Logf("Updated Employee salary: %+v", empPtr)
	emp.UpdateSalary(100000.0)

}

type Programmer interface{
	WriteHelloWorld() string
	WriteName() string
}

type GoProgrammer struct {
	FirstName string	
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "Hello, World! from Go"
}	

func (g *GoProgrammer) WriteName() string {
	return g.FirstName
}


func TestClient(t *testing.T) {
	p := &GoProgrammer{FirstName: "John"}
	t.Logf("Hello, %s", p.WriteHelloWorld())
	t.Logf("Name: %s", p.WriteName())
}


type Pet struct {
	Name string
}

func (p *Pet) Speak() string {
	return "Woof!"
}

func (p *Pet) SpeakTo(target string) string {
	return fmt.Sprintf("%s says: Woof! to %s", p.Name, target)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() string {
	return "Bark!"
}
// go 多使用组合
func TestDog(t *testing.T) {
	dog := new(Dog)
	t.Log(dog.Speak())
	t.Log(dog.SpeakTo("John"))
}


func DoSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case float64:
		fmt.Printf("Float64: %f\n", v)
	default:
		fmt.Println("Unknown type")
	}
}

func TestDoSomething(t *testing.T) {
	DoSomething(42)
	DoSomething("Hello")
	DoSomething(3.14)
	DoSomething(true)
}


type Person interface {
	Speak() string
	Hand()	int
	Gender() string
}


type Yuan struct {
	Person
}

func (y *Yuan) Speak() string {
	return "Hello, I'm Yuan"
}

func (y *Yuan) Hand() int {
	return 2
}

func (y *Yuan) Gender() string {
	return "Male"
}

func TestYuan(t *testing.T) {
	y := &Yuan{}
	t.Logf("Yuan says: %s", y.Speak())
	t.Logf("Yuan has %d hands", y.Hand())
	t.Logf("Yuan is %s",y.Gender())
}


func TestPanicVxExit(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Recovered from panic: %v", r)
		}
	}()

	defer func() {
		fmt.Println("Defer function executed")
	}()
	panic("This is a panic")
}

// 依赖管理