package main

import "fmt"

type First struct {
	Name string
}

type Second struct {
	name string
}
type third struct {
	First
	Second
}

func (obj First) myfunction1() string {
	return obj.Name
}

func (obj Second) myfunction2() string {
	return obj.name
}
func main() {
	b1 := third{
		First{
			Name: "inheritance successfully",
		},
		Second{
			name: "created",
		},
	}
	fmt.Println(b1.myfunction1(), b1.myfunction2())
}
