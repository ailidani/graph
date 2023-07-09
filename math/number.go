package math

import "math"

func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func Sign(x float64) float64 {
	switch {
	case x < 0:
		return -1.0
	case x > 0:
		return 1.0
	}
	return 0
}

func Divisors(n int) []int {
	divisors := make([]int, 0, 1)
	divisors = append(divisors, 1)
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			divisors = append(divisors, n/i)
		}
	}
	return divisors
}

func AmicableNumbers(n int) int {
	m := Sum(Divisors(n))
	if m != n && m > n && Sum(Divisors(m)) == n {
		return m
	}
	return 0
}

func PerfectNumber(n int) bool {
	return Sum(Divisors(n)) == n
}

// GCDEuclidean calculates GCD by Euclidian algorithm.
func GCDEuclidean(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

// GCDRemainder calculates GCD iteratively using remainder.
func GCDRemainder(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	}

	sqrt := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsCoprime(a, b int) bool {
	return GCDRemainder(a, b) == 1
}

// Get all prime factors of a given number n
func PrimeFactors(n int) []int {
	pfs := make([]int, 0)
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return pfs
}
