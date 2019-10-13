package main

import "testing"

// Test for empty argument
func TestHello(t *testing.T) {
	emptyResult := hello("")
	if emptyResult != "Welcome Boss" {
		t.Errorf("hello(\"\") failed, expected %v, got %v", "Welcome Boss", emptyResult)
	}
	result := hello("testuser")
	if result != "Welcome testuser" {
		t.Errorf("hello(\"testuser\") failed, expected %v, got %v", "Welcome Testuser", result)
	}
}
