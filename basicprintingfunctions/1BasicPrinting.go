// It's needed to define package main and func main to create an executable
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "John", Age: 30}
	fmt.Printf("%v\n", person)
}
