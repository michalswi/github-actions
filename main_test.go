package main

import (
	"testing"
)

// test Name function
func TestName(t *testing.T) {

	// test empty arg
	emptyResult := Name("")
	if emptyResult != "noname" {
		t.Errorf("func Name(\"\") failed, expected %v, got %v ", "noname", emptyResult)
	}

	// test valid arg
	result := Name("Mike")
	if result != "Hi Mike" {
		t.Errorf("func Name(\"Mike\") failed, expected %v, got %v ", "Hi Mike", result)
	}
}
