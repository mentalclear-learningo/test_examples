package adder

import "testing"

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		// Use Fatalf to hard stop the test execution.
		t.Errorf("incorrect result: expected %d, got %d", 5, result)
	}
}
