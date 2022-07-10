// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Fix the backing array problem
//
//  Ensure that changing the elements of the `mine` slice
//  does not change the elements of the `nums` slice.
//
//
// CURRENT OUTPUT (INCORRECT)
//
//  Mine         : [-50 -100 -150 25 30 50]
//  Original nums: [-50 -100 -150]
//
//
// EXPECTED OUTPUT
//
//  Mine         : [-50 -100 -150]
//  Original nums: [56 89 15]
//
// ---------------------------------------------------------

func main() {
	nums := []int{56, 89, 15, 25, 30, 50}

	// ----------------------------------------
	// breaks the connection:
	// mine and nums now have different backing arrays

	// verbose solution:
	// var mine []int
	// mine = append(mine, nums[:3]...)

	// better solution (almost the same thing):
	mine := append([]int(nil), nums[:3]...)
	// ----------------------------------------

	mine[0], mine[1], mine[2] = -50, -100, -150
	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}
