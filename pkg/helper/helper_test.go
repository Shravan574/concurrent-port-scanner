package helper

import "testing"

func TestGreet(t *testing.T) {
	expected := "Hello, Alice!"
	actual := Greet("Alice")
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
