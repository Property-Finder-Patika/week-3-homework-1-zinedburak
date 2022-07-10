package main

import (
	"fmt"
	"unicode"
)

func main() {
	slice := []int{4, 3, 2, 1, 0}
	var emptySlice []int
	fmt.Println(sliceRotation(slice, 8, 'r'))
	fmt.Println(sliceRotation(emptySlice, 10, 'l')) // with empty slice
	fmt.Println(sliceRotation(slice, 104, 'b'))     // With wrong rotation direction
	fmt.Println(sliceRotation(slice, 12, 'r'))      // Right Rotation
	fmt.Println(sliceRotation(slice, 12, 'l'))      // Left Rotation

}

func sliceRotation(slice []int, numberOfRotation int, direction rune) []int {
	// Empty Slice check
	if len(slice) == 0 {
		return nil
	}
	// Instead of rotating the number of rotation that is given by the user
	// We rotate the slice by the value of NoR % len(slice) because
	// Each time a slice is rotated len(slice) times we get the same slice that we start with
	numberOfRotation = numberOfRotation % len(slice)
	if unicode.ToUpper(direction) == 'L' {
		// Rotating slice to the left by one for NoR times that we calculated
		for i := 0; i < numberOfRotation; i++ {
			// The Rotation is done by getting the first element and adding it to the end of the slice
			slice = append(slice[1:], slice[0])
		}
	} else if unicode.ToUpper(direction) == 'R' {
		// Rotating slice to the left by one for NoR times that we calculated
		for i := 0; i < numberOfRotation; i++ {
			// The Rotation is done by getting the last element and adding it to the start of the slice
			slice = append([]int{slice[len(slice)-1]}, slice[0:len(slice)-1]...)
		}
	} else {
		return nil
	}
	return slice
}
