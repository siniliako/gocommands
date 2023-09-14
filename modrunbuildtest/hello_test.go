// We need the same package of the function to test
package main

// This is a internet library to test in Golang
import (
	"testing"
)

func TestHello(t *testing.T) {
	actualString := hello()
	expectedString := "Hello, world"
	if actualString != expectedString {
		t.Errorf("Expected string(%s) is not the same as actual string= (%s)", expectedString, actualString)
	}
}
