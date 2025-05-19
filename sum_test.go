package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	c := Sum(10, 21)
	if c != 31 {
		t.Errorf("Expected 30, but got %d", c)
	}
	fmt.Println("sindhu")

}
