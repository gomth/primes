package primes

import (
	"errors"
	"math"
)

// Generate - generate prime numbers in range [1..to] using "Sieve of Atkin" algo
func Generate(to int) ([]int, error) {

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
	sieve_lim := int(math.Sqrt(fto))

	if sieve_lim <= 0 || to <= 0 {
		return nil, errors.New("invalid range")
	}

	// init sieve
	is_prime := make([]bool, to+1, to+1) // all elements are false (bool default)

	is_prime[2] = true
	is_prime[3] = true
	primes_count := 2

	x2 := 0
	n := 0
	for i := 1; i <= sieve_lim; i++ {
		x2 += 2*i - 1
		y2 := 0
		for j := 1; j <= sieve_lim; j++ {
			y2 += 2*j - 1
			n = 4*x2 + y2

			if (n <= to) && (n%12 == 1 || n%12 == 5) {
				is_prime[n] = !is_prime[n]
				if is_prime[n] {
					primes_count++
				}
			}

			// n = 3 * x2 + y2;
			n -= x2
			if (n <= to) && (n%12 == 7) {
				is_prime[n] = !is_prime[n]
				if is_prime[n] {
					primes_count++
				}
			}

			// n = 3 * x2 - y2;
			n -= 2 * y2
			if (i > j) && (n <= to) && (n%12 == 11) {
				is_prime[n] = !is_prime[n]
				if is_prime[n] {
					primes_count++
				}
			}

		}
	}

	// Filter out squares of existing prime numbers [5, sqrt(to)].
	for i := 5; i <= sieve_lim; i++ {
		if is_prime[i] {
			n = i * i
			for j := n; j <= to; j += n {
				if is_prime[j] {
					primes_count--
				}
				is_prime[j] = false
			}
		}
	}

	res := make([]int, primes_count, primes_count)
	res[0] = 2
	res[1] = 3
	res[2] = 5
	idx := 3
	for i := 6; i <= to; i++ {
		if is_prime[i] {
			res[idx] = i
			idx++
		}
	}

	return res, nil
}
