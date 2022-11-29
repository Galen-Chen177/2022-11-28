package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "aaa"
	for item := range s {
		fmt.Printf("%T", item)
	}
}
