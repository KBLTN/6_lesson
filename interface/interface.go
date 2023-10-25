package main

import "fmt"

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2
}

type interfaceHere interface {
	Sum() int
}

func main() {
	//var a interfaceHere
	//sh := structHere{N1: 1, N2: 2}
	//os := otherStruct{A: 2, B: 3}
	//
	//a = &sh
	////fmt.Println(a.Sum())
	//a = os
	////fmt.Println(a.Sum())
	var os *structHere
	var i interfaceHere
	i = os
	//fmt.Println(i.Sum())
	fmt.Printf("(%v, %T)\n", i, i)
}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B

}
