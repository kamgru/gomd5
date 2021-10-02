package main

import (
	"fmt"
	"math"
)

func calculatePaddingBytesCount(inputLength int) int {
	quotient := inputLength / 64
	return (64 * quotient) + 56 - inputLength
}

func encodeLength(length int64) [8]byte {

	var bytes [8]byte

	for i := 0; i < 8; i++ {
		bytes[i] = byte(length >> (8 * (7 - i)))
	}

	return bytes
}

func getBytes(input uint32) []byte {

	b := 16

	bytes := make([]byte, b)

	for i := 0; i < b; i++ {
		bytes[i] = byte(input >> (b * ((b - 1) - i)))
	}

	return bytes
}

func getAppendBytes(input []byte) []byte {
	paddingBytesCount := calculatePaddingBytesCount(len(input))
	paddingBytes := make([]byte, paddingBytesCount)
	paddingBytes[0] = 127

	lengthBytes := encodeLength(int64(len(input)*8))

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

func leftRotate(value uint32, count uint32) uint32{
	return (value << count) | (value >> (32 - count))
}

func hash(input []byte) []byte{
	data := prepareInput(input)

	A0 := uint32(0x01234567)
	B0 := uint32(0x89abcdef)
	C0 := uint32(0xfedcba98)
	D0 := uint32(0x76543210)


	chunkCount := len(data) / 64

	for i := 0; i < chunkCount; i++ {
		chunk := data[i*64:(i+1)*64]

		A := uint32(0x01234567)
		B := uint32(0x89abcdef)
		C := uint32(0xfedcba98)
		D := uint32(0x76543210)

		var words [16]uint32
		for j := 0; j < 16; j++ {
			 words[j] = mergeInt(chunk[j*4:(j + 1)*4])
		}

		var k uint32
		for k = 0; k < 64; k++ {

			var F, g uint32
			if k <= 15 {
				F = (B & C) | (^B & C)
				g = k
			} else if k <= 31 {
				F = (D & B) | (^D & C)
				g = (5*k + 1) % 16
			} else if k <= 47 {
				F = B ^ C ^ D
				g = (3 * k + 5) % 16
			} else {
				F = C ^ (B | (^D))
				g = (7 * k) % 16
			}

			K := uint32(math.Floor(math.Pow(2, 32) * math.Abs(math.Sin(float64(k + 1)))))

			F += A + K + words[g]

			A = D
			D = C
			C = B
 			B += leftRotate(F, shifts[k])
		}

		A0 += A
		B0 += B
		C0 += C
		D0 += D

	}

	return getBytes(A0 + B0 + C0 + D0)
}

func mergeInt(bytes []byte) uint32 {

	m := uint32(0)

	for i := 0; i < 4; i++ {
		m |= uint32(bytes[i]) << ((3 - i) * 8)
	}

	return m
}

func main() {

	input := "kamil grudzień"
	hashBytes := hash([]byte(input))

	fmt.Println(string(hashBytes))
}
