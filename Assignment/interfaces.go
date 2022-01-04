package main

import "fmt"

//defined the interface
type Names interface {
	GetName() string
}

//struct1
type Employee struct {
	Name string
}

//struct2
type Manager struct {
	Name string
}

func (name *Employee) GetName() string {
	return "The Employee Name is : " + name.Name
}

func (name *Manager) GetName() string {
	return "The Employee Name is : " + name.Name
}

func PrintDetail(name Names) {
	fmt.Println(name.GetName())

}

func main() {
	val := &Manager{Name: "dante"}
	val2 := &Employee{Name: "Nero"}
	PrintDetail(val)
	PrintDetail(val2)
}
