package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

// markPrimes marks non-prime numbers in the current segment
func markPrimes(segment []bool, primes []int, low, high int) {
	for _, prime := range primes {
		// Find the starting multiple of the prime in the segment
		start := math.Max(float64(prime*prime), float64(low+((prime-low%prime)%prime)))
		for j := int(start); j < high; j += prime {
			segment[j-low] = true
		}
	}
}

// simpleSieve generates all primes up to the square root of the target
func simpleSieve(limit int) []int {
	isPrime := make([]bool, limit+1)
	for i := 2; i*i <= limit; i++ {
		if !isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = true
			}
		}
	}

	// Collect all prime numbers
	primes := []int{}
	for i := 2; i <= limit; i++ {
		if !isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// findNthPrime uses the segmented sieve to find the nth prime number
func findNthPrime(n int) int {
	// Calculate an approximate upper bound for the nth prime using the prime number theorem
	upperBound := int(float64(n) * (math.Log(float64(n)) + math.Log(math.Log(float64(n)))))
	sqrtLimit := int(math.Sqrt(float64(upperBound)))
	primes := simpleSieve(sqrtLimit)

	segmentSize := 10_000_000 // Size of each segment
	count := len(primes)      // Initial count of primes from the simple sieve
	low := sqrtLimit + 1
	high := low + segmentSize

	var nthPrime int
	numWorkers := runtime.NumCPU()

	// Channel to collect segment ranges and signal termination
	segmentChan := make(chan [2]int, numWorkers)
	var wg sync.WaitGroup

	// Worker goroutines to process segments
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for segmentRange := range segmentChan {
				low, high := segmentRange[0], segmentRange[1]
				segment := make([]bool, high-low)
				markPrimes(segment, primes, low, high)

				for i := 0; i < len(segment) && count < n; i++ {
					if !segment[i] {
						count++
						if count == n {
							nthPrime = low + i
							return
						}
					}
				}
			}
		}()
	}

	// Generate segment ranges and feed them to workers
	for count < n {
		segmentChan <- [2]int{low, high}
		low, high = high, high+segmentSize
	}

	close(segmentChan) // Signal workers to stop
	wg.Wait()          // Wait for all workers to finish

	return nthPrime
}

func main() {
	billionthPrime := findNthPrime(1_000_000_000)
	fmt.Printf("The billionth prime number is: %d\n", billionthPrime)
	// The billionth prime number is: 25316211799
}
