// It's needed to define package main and func main to create an executable
package main

import "fmt"

func main() {
	value := 42
	pointer := &value
	fmt.Printf("Value: %d, Pointer: %p\n", value, pointer)
}
