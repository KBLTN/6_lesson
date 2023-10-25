package main

import "fmt"

func main() {
	var a interface{} = 3.14

	//a = "jelly"
	//fmt.Println(a)
	//fmt.Printf("(%v, %T)\n\n", a, a)
	//a = 42
	//fmt.Println(a)
	//fmt.Printf("(%v, %T)\n\n", a, a)

	//s := a.(string)
	//fmt.Println(s)
	//s, ok := a.(string)
	//fmt.Println(s, ok)
	//f, ok := a.(float32)
	//fmt.Println(f, ok)

	switch a.(type) {
	case int:
		fmt.Println("a is int")
	case string:
		fmt.Println("a is string")
	case bool:
		fmt.Println("a is bool")
	default:
		fmt.Printf("unknown type %T\n", a)

	}
}
