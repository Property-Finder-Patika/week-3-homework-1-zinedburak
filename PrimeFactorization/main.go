package main

import (
	"fmt"
	"math"
)

func main() {
	// A prime number
	fmt.Println(primeFactorization(127))
	// A number that is only divisible by only 2 of the same number
	fmt.Println(primeFactorization(49))

	fmt.Println(primeFactorization(1))

}

func primeFactorization(number int) []int {
	var primes []int
	// Checking the edge cases
	if number < 1 {
		return nil
	} else if number == 1 {
		primes = append(primes, 1)
		return primes
	}
	// Getting all the prime numbers until the number in case the given number is also prime
	primeNums := getAllPrimes(number)
	// creating an index to traverse over the prime numbers
	index := 0
	primeNum := primeNums[index]
	for number != 1 {
		// Checking if the number is divisible by the prime number if so divide the number and append the prime number
		// to the primes array
		if number%primeNum == 0 {
			number /= primeNum
			primes = append(primes, primeNum)

		} else {
			// if the number is not divisible by the current prime number go to the next prime number
			index++
			primeNum = primeNums[index]
		}
	}
	return primes
}

// Sieve of Eratosthenes
func getAllPrimes(limit int) []int {
	//Check for illegal limit
	if limit < 2 {
		return nil
	}

	// Creation of the Number Array Step 1
	var numberArray []int
	for i := 2; i <= limit; i++ {
		numberArray = append(numberArray, i)
	}

	primeNum := float64(numberArray[0])
	generalIndex := 0

	// Check in step 5
	for primeNum <= math.Sqrt(float64(limit)) {

		// Always starting the iteration over the array from the letter index that last prime was on if the last prime t
		//hat we have found is 3, and it is on index 2 we will iterate over the array [4...n]

		index := generalIndex + 1

		// Iteration of the array  Step 2
		for index < len(numberArray) {
			if numberArray[index]%int(primeNum) == 0 {
				// Elimination of the element if it is divisible by our prime number Step 3
				numberArray = remove(numberArray, index)
				index--
			}
			index++
		}
		// Updating the value of the next prime Step 5
		generalIndex++
		primeNum = float64(numberArray[generalIndex])
	}
	return numberArray
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
