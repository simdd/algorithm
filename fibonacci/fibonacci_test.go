package fibonacci

import 	"testing"
import "fmt"

// TestFibonacci1 46.148s
func TestFibonacci1(t *testing.T) {
	fmt.Printf("%d \n", fibonacci1(50))
}

// TestFibonacci2 0.005s
func TestFibonacci2(t *testing.T) {
	fmt.Printf("%d \n", fibonacci2(50))
}
