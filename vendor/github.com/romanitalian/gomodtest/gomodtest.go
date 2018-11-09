package gomodtest

import "fmt"

// Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s! It is great Day!", name)
}
