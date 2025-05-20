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
	d := multiply(20, 30)
	if d != 600 {
		t.Errorf("Expected 600, but got %d", d)
		e := divide(10, 3)
		if e != 3 {
			t.Errorf("Expected 3,but got %d", e)
		}
	}

}
