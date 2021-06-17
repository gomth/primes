package primes

import (
	"errors"
)

// GenerateEratosthenes - simple way to get list of prime numbers.
// Complexity: O(n*log(log n)/2) -> /2 becasue we  skip even numbers
// Memory: n/2 -> allocating memory only for odd flags
// Improvement: We dont check even numbers at all (except 2)
func GenerateEratosthenes(to int) ([]int, error) {

	if to <= 0 {
		return nil, errors.New("Invalid input range")
	}

	if to == 1 {
		return []int{}, nil
	}

	if to == 2 {
		return []int{2}, nil
	}

	flagsBufferLen := to / 2
	if !isEven(to) {
		flagsBufferLen++
	}
	isPrime := make([]bool, flagsBufferLen)
	for i := range isPrime {
		isPrime[i] = true // thanks go
	}

	res := make([]int, 1, len(isPrime)) // at least one elem - 3
	res[0] = 2                          // 2 already "found"

	for i := 3; i <= to; i += 2 {
		if !isPrime[i/2-1] { // /2-1 because we allocated memory only for odd flags. And [0] flag points n==3 etc.
			continue // skip already non-prime number
		}
		res = append(res, i)
		for j := i + i; j <= to; j += i {
			if isEven(j) {
				continue // all even numbers are not primes except 2
			}
			isPrime[j/2-1] = false
		}
	}

	return res, nil
}

func isEven(n int) bool {
	return n%2 == 0
}
