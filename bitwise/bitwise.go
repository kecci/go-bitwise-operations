package bitwise

// And performs the bitwise AND operation between two numbers.
func And(a, b uint8) uint8 {
	return a & b
}

// Or performs the bitwise OR operation between two numbers.
func Or(a, b uint8) uint8 {
	return a | b
}

// Xor performs the bitwise XOR operation between two numbers.
func Xor(a, b uint8) uint8 {
	return a ^ b
}

// Not performs the bitwise NOT operation on a number.
func Not(a uint8) uint8 {
	return ^a
}
