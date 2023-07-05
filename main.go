package main

import (
	"fmt"

	"github.com/kecci/go-bitwise-operations/bitshift"
	"github.com/kecci/go-bitwise-operations/bitwise"
)

func main() {
	// ===================================================================
	// 								BITWISE
	// ===================================================================

	a := uint8(12) // Binary: 00001100
	b := uint8(10) // Binary: 00001010

	// Bitwise AND
	fmt.Printf("[Bitwise] AND: %08b\n", bitwise.And(a, b))

	// Bitwise OR
	fmt.Printf("[Bitwise] OR:  %08b\n", bitwise.Or(a, b))

	// Bitwise XOR
	fmt.Printf("[Bitwise] XOR: %08b\n", bitwise.Xor(a, b))

	// Bitwise NOT
	fmt.Printf("[Bitwise] NOT: %08b\n", bitwise.Not(a))

	println("=============================================================")

	// ===================================================================
	// 								BITSHIFT
	// ===================================================================

	num := 1      // Binary: 1010
	shift := 31   // Shift Position: 31
	numBits := 31 // N-bit

	fmt.Printf("[Bitshift] Original number: %d\n", num)

	// Bitshift Left shift by 31 shifts
	fmt.Printf("[Bitshift] LeftShift number: %d (Shifts: %d)\n", bitshift.LeftShift(num, shift), shift)

	// Bitshift Right shift by 31 shifts
	fmt.Printf("[Bitshift] RightShift number: %d (Shifts: %d)\n", bitshift.RightShift(num, shift), shift)

	// Bitshift Left circular shift by 31 shifts and 32-bit
	leftCircular, err := bitshift.LeftCircularShift(num, shift, numBits)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Bitshift] LeftCircularShift number: %08b (Shifts: %d Bits: %d)\n", leftCircular, shift, numBits)

	// Bitshift Right circular shift by 31 shifts and 32-bit
	rightCircular, err := bitshift.RightCircularShift(num, shift, numBits)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Bitshift] RightCircularShift number: %08b (Shifts: %d Bits: %d)\n", rightCircular, shift, numBits)
}
