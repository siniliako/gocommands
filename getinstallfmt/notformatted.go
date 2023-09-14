// It's needed to define package main and func main to create an executable
package main

import "fmt"

func hello() string { return "Hello, world" }
func main()         { fmt.Println(hello()) }
