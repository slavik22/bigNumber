package main

import "encoding/hex"

type LargeNumber struct {
	blocks    []uint64
	blockSize int
}

func NewLargeNumber(hexString string, blockSize int) (*LargeNumber, error) {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	blockSizeBytes := blockSize / 8
	numBlocks := len(hexBytes) / blockSizeBytes
	blocks := make([]uint64, numBlocks)
	for i := 0; i < numBlocks; i++ {
		blockStart := i * blockSizeBytes
		blockEnd := blockStart + blockSizeBytes
		block := uint64(0)
		for j := blockStart; j < blockEnd; j++ {
			block <<= 8
			block |= uint64(hexBytes[j])
		}
		blocks[i] = block
	}

	return &LargeNumber{blocks: blocks, blockSize: blockSize}, nil
}

func (ln *LargeNumber) ToHexString() string {
	hexString := ""
	for _, block := range ln.blocks {
		blockBytes := make([]byte, ln.blockSize/8)
		for i := ln.blockSize/8 - 1; i >= 0; i-- {
			blockBytes[i] = byte(block & 0xFF)
			block >>= 8
		}
		hexString += hex.EncodeToString(blockBytes)
	}

	return hexString
}

func (ln *LargeNumber) INV() {
	for i := range ln.blocks {
		ln.blocks[i] = ^ln.blocks[i]
	}
}

func (ln *LargeNumber) XOR(other *LargeNumber) {
	if len(ln.blocks) != len(other.blocks) {
		return
	}

	for i := range ln.blocks {
		ln.blocks[i] ^= other.blocks[i]
	}
}

func (ln *LargeNumber) OR(other *LargeNumber) {
	if len(ln.blocks) != len(other.blocks) {
		return
	}

	for i := range ln.blocks {
		ln.blocks[i] |= other.blocks[i]
	}
}

func (ln *LargeNumber) AND(other *LargeNumber) {
	if len(ln.blocks) != len(other.blocks) {
		return
	}

	for i := range ln.blocks {
		ln.blocks[i] &= other.blocks[i]
	}
}

func (ln *LargeNumber) shiftR(n int) {
	shiftSize := uint(n % ln.blockSize)
	carry := uint64(0)

	for i := len(ln.blocks) - 1; i >= 0; i-- {
		value := ln.blocks[i]
		shifted := (value >> shiftSize) | (carry << (uint(ln.blockSize) - shiftSize))
		ln.blocks[i] = shifted
		carry = value << (uint(ln.blockSize) - shiftSize)
	}
}

func (ln *LargeNumber) shiftL(n int) {
	shiftSize := uint(n % ln.blockSize)
	carry := uint64(0)

	for i := 0; i < len(ln.blocks); i++ {
		value := ln.blocks[i]
		shifted := (value << shiftSize) | carry
		ln.blocks[i] = shifted
		carry = value >> (uint(ln.blockSize) - shiftSize)
	}
}

func (ln *LargeNumber) ADD(other *LargeNumber) {
	if len(ln.blocks) != len(other.blocks) {
		return
	}

	var carry uint64
	for i := range ln.blocks {
		sum := ln.blocks[i] + other.blocks[i] + carry
		if sum < ln.blocks[i] || sum < other.blocks[i] {
			carry = 1
		} else {
			carry = 0
		}
		ln.blocks[i] = sum
	}
}

func (ln *LargeNumber) SUB(other *LargeNumber) {
	if len(ln.blocks) != len(other.blocks) {
		return
	}

	var borrow uint64
	for i := range ln.blocks {
		diff := ln.blocks[i] - other.blocks[i] - borrow
		if diff > ln.blocks[i] {
			borrow = 1
		} else {
			borrow = 0
		}
		ln.blocks[i] = diff
	}
}

func (ln *LargeNumber) MOD(divisor *LargeNumber) {
	if len(ln.blocks) != len(divisor.blocks) {
		return
	}

	for i := range ln.blocks {
		ln.blocks[i] %= divisor.blocks[i]
	}
}
