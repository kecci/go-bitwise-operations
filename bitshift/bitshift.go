package bitshift

import (
	"fmt"
)

// LeftShift perform a bitwise left shift
func LeftShift(num, shift int) int {
	return num << shift
}

// RightShift perform a bitwise right shift
func RightShift(num, shift int) int {
	return num >> shift
}

// LeftCircularShift perform bitwise left circular shift
// numBits is the number of bits in an integer. e.g., numBits := 8
func LeftCircularShift(num, shift, numBits int) (int, error) {
	if numBits < shift {
		return 0, fmt.Errorf("out of range shift(%d) of numBits(%d)", shift, numBits)
	}
	return (num << shift) | (num >> (numBits - shift)), nil
}

// RightCircularShift perform a bitwise right circular shift
// numBits is the number of bits in an integer. e.g., numBits := 8
func RightCircularShift(num, shift, numBits int) (int, error) {
	if numBits < shift {
		return 0, fmt.Errorf("out of range shift(%d) of numBits(%d)", shift, numBits)
	}
	return (num >> shift) | (num << (numBits - shift)), nil
}
