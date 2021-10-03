package main

import "math"

func ComputeHash(input []byte) []byte {
	data := prepareInput(input)

	var A0, B0, C0, D0 uint32 = 0x67452301, 0xEFCDAB89, 0x98BADCFE, 0x10325476

	chunkCount := len(data) / 64

	for i := 0; i < chunkCount; i++ {
		chunk := data[i*64 : (i+1)*64]

		var A, B, C, D uint32 = A0, B0, C0, D0

		var words [16]uint32
		for j := 0; j < 16; j++ {
			words[j] = mergeInt(chunk[j*4 : (j+1)*4])
		}

		var k uint32
		for k = 0; k < 64; k++ {

			var F, g uint32
			if k <= 15 {
				F = (B & C) | (^B & D)
				g = k
			} else if k >= 16 && k <= 31 {
				F = (D & B) | (^D & C)
				g = (5*k + 1) % 16
			} else if k >= 32 && k <= 47 {
				F = B ^ C ^ D
				g = (3*k + 5) % 16
			} else {
				F = C ^ (B | (^D))
				g = (7 * k) % 16
			}

			K := uint32(math.Floor(math.Pow(2, 32) * math.Abs(math.Sin(float64(k+1)))))

			A += F + K + words[g]

			dTemp := D
			D = C
			C = B
			B += leftRotate(A, shifts[k])
			A = dTemp
		}

		A0 += A
		B0 += B
		C0 += C
		D0 += D

	}

	final := []byte{
		byte(A0),
		byte(A0 >> 8),
		byte(A0 >> 16),
		byte(A0 >> 24),
		byte(B0),
		byte(B0 >> 8),
		byte(B0 >> 16),
		byte(B0 >> 24),
		byte(C0),
		byte(C0 >> 8),
		byte(C0 >> 16),
		byte(C0 >> 24),
		byte(D0),
		byte(D0 >> 8),
		byte(D0 >> 16),
		byte(D0 >> 24),
	}

	return final
}

func calculatePaddingBytesCount(inputLength int) int {
	quotient := inputLength / 64
	return (64 * quotient) + 56 - inputLength
}

func encodeLength(length int64) [8]byte {

	var bytes [8]byte

	for i := 0; i < 8; i++ {
		bytes[i] = byte(length >> (8 * i))
	}

	return bytes
}

func getAppendBytes(input []byte) []byte {
	paddingBytesCount := calculatePaddingBytesCount(len(input))
	paddingBytes := make([]byte, paddingBytesCount)
	paddingBytes[0] = 128

	lengthBytes := encodeLength(int64(len(input) * 8))

	return append(paddingBytes, lengthBytes[:]...)
}

func prepareInput(input []byte) []byte {
	appendBytes := getAppendBytes(input)
	return append(input, appendBytes...)
}

var shifts = [64]uint32{
	7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
	5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
	4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
	6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
}

func leftRotate(value uint32, count uint32) uint32 {
	return (value << count) | (value >> (32 - count))
}

func mergeInt(bytes []byte) uint32 {

	a := uint32(bytes[0])
	b := uint32(bytes[1]) << 8
	c := uint32(bytes[2]) << 16
	d := uint32(bytes[3]) << 24

	result := a + b + c + d
	return result
}
