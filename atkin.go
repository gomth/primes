package primes

import (
	"errors"
	"math"
)

// GenerateAtkin - generate prime numbers in range [1..to] using "Sieve of Atkin" algo
func GenerateAtkin(to int) ([]int, error) {

	if to == 1 {
		return []int{}, nil
	}

	if to == 2 {
		return []int{2}, nil
	}

	if to == 3 {
		return []int{2, 3}, nil
	}

	fto := float64(to)
	sieveLim := int(math.Sqrt(fto))

	if sieveLim <= 0 || to <= 0 {
		return nil, errors.New("invalid range")
	}

	// init sieve
	isPrime := make([]bool, to+1, to+1) // all elements are false (bool default)

	isPrime[2] = true
	isPrime[3] = true
	primesCount := 2

	x2 := 0
	n := 0
	for i := 1; i <= sieveLim; i++ {
		x2 += 2*i - 1
		y2 := 0
		for j := 1; j <= sieveLim; j++ {
			y2 += 2*j - 1
			n = 4*x2 + y2

			if (n <= to) && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
				if isPrime[n] {
					primesCount++
				}
			}

			// n = 3 * x2 + y2;
			n -= x2
			if (n <= to) && (n%12 == 7) {
				isPrime[n] = !isPrime[n]
				if isPrime[n] {
					primesCount++
				}
			}

			// n = 3 * x2 - y2;
			n -= 2 * y2
			if (i > j) && (n <= to) && (n%12 == 11) {
				isPrime[n] = !isPrime[n]
				if isPrime[n] {
					primesCount++
				}
			}

		}
	}

	// Filter out squares of existing prime numbers [5, sqrt(to)].
	for i := 5; i <= sieveLim; i++ {
		if isPrime[i] {
			n = i * i
			for j := n; j <= to; j += n {
				if isPrime[j] {
					primesCount--
				}
				isPrime[j] = false
			}
		}
	}

	res := make([]int, primesCount, primesCount)
	res[0] = 2
	res[1] = 3
	if primesCount >= 3 {
		res[2] = 5
	}
	idx := 3
	for i := 6; i <= to; i++ {
		if isPrime[i] {
			res[idx] = i
			idx++
		}
	}

	return res, nil
}
