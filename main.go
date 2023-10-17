package main

import (
	"fmt"
)

func main() {
	// Example usage:
	hexInput := "1A4B2F8C34"
	blockSize := 64 // Choose the block size (8, 16, 32, or 64)
	ln, err := NewLargeNumber(hexInput, blockSize)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Original Hexadecimal Input: %s\n", hexInput)
	fmt.Printf("Converted LargeNumber: %s\n", ln.ToHexString())

	// Perform operations on LargeNumber
	ln.INV()
	fmt.Printf("Bitwise Inversion: %s\n", ln.ToHexString())

	ln2, _ := NewLargeNumber("0123456789ABCDEF", blockSize)
	ln.XOR(ln2)
	fmt.Printf("Bitwise XOR with 0123456789ABCDEF: %s\n", ln.ToHexString())

	ln3, _ := NewLargeNumber("FEDCBA9876543210", blockSize)
	ln.OR(ln3)
	fmt.Printf("Bitwise OR with FEDCBA9876543210: %s\n", ln.ToHexString())
}
