package primes

import (
	"testing"
)

func TestGenerateEratInvalid(t *testing.T) {
	actual, err := GenerateEratosthenes(-1)
	if err == nil {
		t.Errorf("GenerateInvalid failed.\nError expected")
	}

	if len(actual) != 0 {
		t.Errorf("GenerateInvalid failed.\nExpected: %v\nActual: %v", []int{}, actual)
	}
}

func TestGenerateErat5(t *testing.T) {
	actual, err := GenerateEratosthenes(5)
	if err != nil {
		t.Errorf("generate(5) failed.\nError: %v", err)
	}

	expected := []int{2, 3, 5}

	if !cmp(actual, expected) {
		t.Errorf("generate(5) failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGenerateEratEven(t *testing.T) {
	actual, err := GenerateEratosthenes(8)
	if err != nil {
		t.Errorf("generate(8) failed.\nError: %v", err)
	}

	expected := []int{2, 3, 5, 7}

	if !cmp(actual, expected) {
		t.Errorf("generate(8) failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGenerateEratForRange(t *testing.T) {
	actual, err := GenerateEratosthenes(50)
	if err != nil {
		t.Errorf("Generate(0, 50) failed.\nError: %v", err)
	}

	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}

	if !cmp(actual, expected) {
		t.Errorf("Generate(0, 50) failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}

// TODO: use cpu and memory profile output. https://blog.golang.org/pprof
func BenchmarkGenerateErat100(b *testing.B) {
	for i := 1; i < b.N; i++ {
		_, err := GenerateEratosthenes(i)
		if err != nil {
			b.Fatalf("Generate(0, 1000) failed.\nError: %v", err)
		}
	}
}
