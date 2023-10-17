package main

import (
	"fmt"
	"testing"
)

func TestLargeNumber(t *testing.T) {
	testCases := []struct {
		hexInput       string
		blockSize      int
		operation      func(*LargeNumber)
		expectedResult string
	}{
		{
			hexInput:       "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4",
			blockSize:      64,
			operation:      func(ln *LargeNumber) { ln.INV() },
			expectedResult: "ae409f7beb52a8d95c3e413f670884e4ab004d878072ad758b3e280219b8f15b",
		},
		{
			hexInput:  "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4",
			blockSize: 64,
			operation: func(ln *LargeNumber) {
				ln2, _ := NewLargeNumber("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c", 64)
				ln.XOR(ln2)
			},
			expectedResult: "1182d8299c0ec40ca8bf3f49362e95e4ecedaf82bfd167988972412095b13db8",
		},
		{
			hexInput:  "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4",
			blockSize: 64,
			operation: func(ln *LargeNumber) {
				ln3, _ := NewLargeNumber("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c", 64)
				ln.OR(ln3)
			},
			expectedResult: "51bff8ad9cafd72eabffbfc9befffffffcffbffaffdd779afdf3d7fdf7f73fbc",
		},
		{
			hexInput:  "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4",
			blockSize: 64,
			operation: func(ln *LargeNumber) {
				ln3, _ := NewLargeNumber("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c", 64)
				ln.AND(ln3)
			},
			expectedResult: "403d208400a113220340808088d16a1b10121078400c1002748196dd62460204",
		},
		{
			hexInput:  "36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80",
			blockSize: 64,
			operation: func(ln *LargeNumber) {
				ln3, _ := NewLargeNumber("70983d692f648185febe6d6fa607630ae68649f7e6fc45b94680096c06e4fadb", 64)
				ln.ADD(ln3)
			},
			expectedResult: "a78865c13b14ae4d25e90771b54963ed2d68c0a64d4a8ba8c6f45ee0e9daa65c",
		},
		{
			hexInput:  "33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc",
			blockSize: 64,
			operation: func(ln *LargeNumber) {
				ln3, _ := NewLargeNumber("22e962951cb6cd2ce279ab0e2095825c141d48ef3ca9dabf253e38760b57fe03", 64)
				ln.SUB(ln3)
			},
			expectedResult: "10e570324e6ffdbd6b9c813dec968d9bad134bc0dbb061520934f4e59c2700b9",
		},
	}

	for _, testCase := range testCases {
		t.Run("Block Size "+string(rune(testCase.blockSize)), func(t *testing.T) {
			ln, err := NewLargeNumber(testCase.hexInput, testCase.blockSize)
			if err != nil {
				t.Fatalf("Error creating LargeNumber: %v", err)
			}

			testCase.operation(ln)
			result := ln.ToHexString()

			fmt.Println(result)
			fmt.Println(testCase.expectedResult)

			if result != testCase.expectedResult {
				t.Errorf("Expected: %s, Got: %s", testCase.expectedResult, result)
			}
		})
	}
}
