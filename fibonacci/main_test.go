package problem

import "testing"

// TestFibonacci1 46.148s
func TestFibonacci1(t *testing.T) {
	ret := fibonacci1(10)

	if ret != 89 {
		t.Error("TestFibonacci1 ", ret)
	}
}

// TestFibonacci2 0.005s
func TestFibonacci2(t *testing.T) {
	ret := fibonacci2(10)

	if ret != 89 {
		t.Error("TestFibonacci1 ", ret)
	}
}
