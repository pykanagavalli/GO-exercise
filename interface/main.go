package main

import "fmt"

type myStructOne struct {
	Name string
}

type myStructtwo struct {
	Name string
}

func (obj myStructOne) myMethod() string {
	return obj.Name
}

func (obj myStructtwo) myMethod() string {
	return obj.Name
}

type Myinterface interface {
	myMethod() string
}

func main() {
	var Myinterface Myinterface
	
	Myinterface = myStructOne{Name: "hello"}
	fmt.Println(Myinterface.myMethod())

	Myinterface = myStructtwo{Name: "world"}
	fmt.Println(Myinterface.myMethod())
}
