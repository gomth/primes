package primes

import "testing"

// lets assume that slices are sorted
func cmp(left []int, right []int) bool {

	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func TestGenerateAtkinInvalid(t *testing.T) {
	actual, err := GenerateAtkin(-1)
	if err == nil {
		t.Errorf("GenerateInvalid failed.\nError expected")
	}

	if len(actual) != 0 {
		t.Errorf("GenerateInvalid failed.\nExpected: %v\nActual: %v", []int{}, actual)
	}
}

func TestGenerateAtkin5(t *testing.T) {
	actual, err := GenerateAtkin(5)
	if err != nil {
		t.Errorf("generate(5) failed.\nError: %v", err)
	}

	expected := []int{2, 3, 5}

	if !cmp(actual, expected) {
		t.Errorf("generate(5) failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGenerateAtkin4(t *testing.T) {
	actual, err := GenerateAtkin(4)
	if err != nil {
		t.Errorf("generate(4) failed.\nError: %v", err)
	}

	expected := []int{2, 3}

	if !cmp(actual, expected) {
		t.Errorf("generate(4) failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGenerateAtkin2(t *testing.T) {
	actual, err := GenerateAtkin(2)
	if err != nil {
		t.Errorf("generate(2) failed.\nError: %v", err)
	}

	expected := []int{2}

	if !cmp(actual, expected) {
		t.Errorf("generate(2) failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGenerateAtkin3(t *testing.T) {
	actual, err := GenerateAtkin(3)
	if err != nil {
		t.Errorf("generate(3) failed.\nError: %v", err)
	}

	expected := []int{2, 3}

	if !cmp(actual, expected) {
		t.Errorf("generate(3) failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGenerateAtkinInvalidRange(t *testing.T) {
	actual, err := GenerateAtkin(1)
	if err != nil {
		t.Errorf("generate(1) failed.\nError: %v", err)
	}

	if len(actual) != 0 {
		t.Errorf("generate(1) failed.\nExpected: %v\nActual: %v", []int{}, actual)
	}
}

func TestGenerateAtkinForRange(t *testing.T) {
	actual, err := GenerateAtkin(50)
	if err != nil {
		t.Errorf("Generate(0, 50) failed.\nError: %v", err)
	}

	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}

	if !cmp(actual, expected) {
		t.Errorf("Generate(0, 50) failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGenerateAtkinBoundsCheck(t *testing.T) {
	actual, err := GenerateAtkin(7)
	if err != nil {
		t.Errorf("generate(7) failed.\nError: %v", err)
	}

	expected := []int{2, 3, 5, 7}

	if !cmp(actual, expected) {
		t.Errorf("generate(7) failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}
