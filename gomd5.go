package main

import "math"

func calculateSine(index uint32) uint32 {
	const multiplier = 4294967296
	sine := multiplier * math.Abs(math.Sin(float64(index)))
	return uint32(math.Floor(sine))
}

func auxiliaryF(B, C, D  uint32) uint32 {
	return  (B & C) | (^B & D)
}

func auxiliaryG(B, C, D  uint32) uint32 {
	return (D & B) | (^D & C)
}

func auxiliaryH(B, C, D  uint32) uint32 {
	return B ^ C ^ D
}

func auxiliaryI(B, C, D  uint32) uint32 {
	return C ^ (B | (^D))
}

func calculateRound(
	a,b,c,d uint32,
	k int,
	s,i uint32,
	words [16]uint32,
	aux func(b, c, d uint32) uint32) uint32 {
	a = b + leftRotate(a + aux(b, c, d) + words[k] + calculateSine(i), s)
	return a
}

func produceOutput(A, B, C, D uint32) []byte {
	result := []byte{
		byte(A),
		byte(A >> 8),
		byte(A >> 16),
		byte(A >> 24),
		byte(B),
		byte(B >> 8),
		byte(B >> 16),
		byte(B >> 24),
		byte(C),
		byte(C >> 8),
		byte(C >> 16),
		byte(C >> 24),
		byte(D),
		byte(D >> 8),
		byte(D >> 16),
		byte(D >> 24),
	}

	return result
}

func ComputeHashRfc(input []byte) []byte {
	data := prepareInput(input)

		A := uint32(0x67452301)
		B := uint32(0xEFCDAB89)
		C := uint32(0x98BADCFE)
		D := uint32(0x10325476)

	chunks := splitInput(data)
	for _, chunk := range chunks {

		words := produceWords(chunk)

		AA, BB, CC, DD := A, B, C, D

		/* Round 1. */
		/* Let [abcd k s i] denote the operation
		   a = b + ((a + F(b,c,d) + X[k] + T[i]) <<< s). */
		/* Do the following 16 operations.
		[ABCD  0  7  1]  [DABC  1 12  2]  [CDAB  2 17  3]  [BCDA  3 22  4]
		[ABCD  4  7  5]  [DABC  5 12  6]  [CDAB  6 17  7]  [BCDA  7 22  8]
		[ABCD  8  7  9]  [DABC  9 12 10]  [CDAB 10 17 11]  [BCDA 11 22 12]
		[ABCD 12  7 13]  [DABC 13 12 14]  [CDAB 14 17 15]  [BCDA 15 22 16]
		*/

		AA = calculateRound(AA, BB, CC, DD, 0, 7, 1, words, auxiliaryF)
		DD = calculateRound(DD, AA, BB, CC, 1, 12, 2, words, auxiliaryF)
		CC = calculateRound(CC, DD, AA, BB, 2, 17, 3, words, auxiliaryF)
		BB = calculateRound(BB, CC, DD, AA, 3, 22, 4, words, auxiliaryF)

		AA = calculateRound(AA, BB, CC, DD, 4, 7, 5, words, auxiliaryF)
		DD = calculateRound(DD, AA, BB, CC, 5, 12, 6, words, auxiliaryF)
		CC = calculateRound(CC, DD, AA, BB, 6, 17, 7, words, auxiliaryF)
		BB = calculateRound(BB, CC, DD, AA, 7, 22, 8, words, auxiliaryF)

		AA = calculateRound(AA, BB, CC, DD, 8, 7, 9, words, auxiliaryF)
		DD = calculateRound(DD, AA, BB, CC, 9, 12, 10, words, auxiliaryF)
		CC = calculateRound(CC, DD, AA, BB, 10, 17, 11, words, auxiliaryF)
		BB = calculateRound(BB, CC, DD, AA, 11, 22, 12, words, auxiliaryF)

		AA = calculateRound(AA, BB, CC, DD, 12, 7, 13, words, auxiliaryF)
		DD = calculateRound(DD, AA, BB, CC, 13, 12, 14, words, auxiliaryF)
		CC = calculateRound(CC, DD, AA, BB, 14, 17, 15, words, auxiliaryF)
		BB = calculateRound(BB, CC, DD, AA, 15, 22, 16, words, auxiliaryF)

		/* Round 2. */
		/* Let [abcd k s i] denote the operation
		   a = b + ((a + G(b,c,d) + X[k] + T[i]) <<< s). */
		/* Do the following 16 operations.
		[ABCD  1  5 17]  [DABC  6  9 18]  [CDAB 11 14 19]  [BCDA  0 20 20]
		[ABCD  5  5 21]  [DABC 10  9 22]  [CDAB 15 14 23]  [BCDA  4 20 24]
		[ABCD  9  5 25]  [DABC 14  9 26]  [CDAB  3 14 27]  [BCDA  8 20 28]
		[ABCD 13  5 29]  [DABC  2  9 30]  [CDAB  7 14 31]  [BCDA 12 20 32]
		 */

		AA = calculateRound(AA, BB, CC, DD, 1, 5, 17, words, auxiliaryG)
		DD = calculateRound(DD, AA, BB, CC, 6, 9, 18, words, auxiliaryG)
		CC = calculateRound(CC, DD, AA, BB, 11, 14, 19, words, auxiliaryG)
		BB = calculateRound(BB, CC, DD, AA, 0, 20, 20, words, auxiliaryG)

		AA = calculateRound(AA, BB, CC, DD, 5, 5, 21, words, auxiliaryG)
		DD = calculateRound(DD, AA, BB, CC, 10, 9, 22, words, auxiliaryG)
		CC = calculateRound(CC, DD, AA, BB, 15, 14, 23, words, auxiliaryG)
		BB = calculateRound(BB, CC, DD, AA, 4, 20, 24, words, auxiliaryG)

		AA = calculateRound(AA, BB, CC, DD, 9, 5, 25, words, auxiliaryG)
		DD = calculateRound(DD, AA, BB, CC, 14, 9, 26, words, auxiliaryG)
		CC = calculateRound(CC, DD, AA, BB, 3, 14, 27, words, auxiliaryG)
		BB = calculateRound(BB, CC, DD, AA, 8, 20, 28, words, auxiliaryG)

		AA = calculateRound(AA, BB, CC, DD, 13, 5, 29, words, auxiliaryG)
		DD = calculateRound(DD, AA, BB, CC, 2, 9, 30, words, auxiliaryG)
		CC = calculateRound(CC, DD, AA, BB, 7, 14, 31, words, auxiliaryG)
		BB = calculateRound(BB, CC, DD, AA, 12, 20, 32, words, auxiliaryG)

		/* Round 3. */
		/* Let [abcd k s t] denote the operation
		   a = b + ((a + H(b,c,d) + X[k] + T[i]) <<< s). */
		/* Do the following 16 operations.
		[ABCD  5  4 33]  [DABC  8 11 34]  [CDAB 11 16 35]  [BCDA 14 23 36]
		[ABCD  1  4 37]  [DABC  4 11 38]  [CDAB  7 16 39]  [BCDA 10 23 40]
		[ABCD 13  4 41]  [DABC  0 11 42]  [CDAB  3 16 43]  [BCDA  6 23 44]
		[ABCD  9  4 45]  [DABC 12 11 46]  [CDAB 15 16 47]  [BCDA  2 23 48]
		 */

		AA = calculateRound(AA, BB, CC, DD, 5, 4, 33, words, auxiliaryH)
		DD = calculateRound(DD, AA, BB, CC, 8, 11, 34, words, auxiliaryH)
		CC = calculateRound(CC, DD, AA, BB, 11, 16, 35, words, auxiliaryH)
		BB = calculateRound(BB, CC, DD, AA, 14, 23, 36, words, auxiliaryH)

		AA = calculateRound(AA, BB, CC, DD, 1, 4, 37, words, auxiliaryH)
		DD = calculateRound(DD, AA, BB, CC, 4, 11, 38, words, auxiliaryH)
		CC = calculateRound(CC, DD, AA, BB, 7, 16, 39, words, auxiliaryH)
		BB = calculateRound(BB, CC, DD, AA, 10, 23, 40, words, auxiliaryH)

		AA = calculateRound(AA, BB, CC, DD, 13, 4, 41, words, auxiliaryH)
		DD = calculateRound(DD, AA, BB, CC, 0, 11, 42, words, auxiliaryH)
		CC = calculateRound(CC, DD, AA, BB, 3, 16, 43, words, auxiliaryH)
		BB = calculateRound(BB, CC, DD, AA, 6, 23, 44, words, auxiliaryH)

		AA = calculateRound(AA, BB, CC, DD, 9, 4, 45, words, auxiliaryH)
		DD = calculateRound(DD, AA, BB, CC, 12, 11, 46, words, auxiliaryH)
		CC = calculateRound(CC, DD, AA, BB, 15, 16, 47, words, auxiliaryH)
		BB = calculateRound(BB, CC, DD, AA, 2, 23, 48, words, auxiliaryH)

		/* Round 4. */
		/* Let [abcd k s t] denote the operation
		   a = b + ((a + I(b,c,d) + X[k] + T[i]) <<< s). */
		/* Do the following 16 operations.
		[ABCD  0  6 49]  [DABC  7 10 50]  [CDAB 14 15 51]  [BCDA  5 21 52]
		[ABCD 12  6 53]  [DABC  3 10 54]  [CDAB 10 15 55]  [BCDA  1 21 56]
		[ABCD  8  6 57]  [DABC 15 10 58]  [CDAB  6 15 59]  [BCDA 13 21 60]
		[ABCD  4  6 61]  [DABC 11 10 62]  [CDAB  2 15 63]  [BCDA  9 21 64]
		 */

		AA = calculateRound(AA, BB, CC, DD, 0, 6, 49, words, auxiliaryI)
		DD = calculateRound(DD, AA, BB, CC, 7, 10, 50, words, auxiliaryI)
		CC = calculateRound(CC, DD, AA, BB, 14, 15, 51, words, auxiliaryI)
		BB = calculateRound(BB, CC, DD, AA, 5, 21, 52, words, auxiliaryI)

		AA = calculateRound(AA, BB, CC, DD, 12, 6, 53, words, auxiliaryI)
		DD = calculateRound(DD, AA, BB, CC, 3, 10, 54, words, auxiliaryI)
		CC = calculateRound(CC, DD, AA, BB, 10, 15, 55, words, auxiliaryI)
		BB = calculateRound(BB, CC, DD, AA, 1, 21, 56, words, auxiliaryI)

		AA = calculateRound(AA, BB, CC, DD, 8, 6, 57, words, auxiliaryI)
		DD = calculateRound(DD, AA, BB, CC, 15, 10, 58, words, auxiliaryI)
		CC = calculateRound(CC, DD, AA, BB, 6, 15, 59, words, auxiliaryI)
		BB = calculateRound(BB, CC, DD, AA, 13, 21, 60, words, auxiliaryI)

		AA = calculateRound(AA, BB, CC, DD, 4, 6, 61, words, auxiliaryI)
		DD = calculateRound(DD, AA, BB, CC, 11, 10, 62, words, auxiliaryI)
		CC = calculateRound(CC, DD, AA, BB, 2, 15, 63, words, auxiliaryI)
		BB = calculateRound(BB, CC, DD, AA, 9, 21, 64, words, auxiliaryI)

		A += AA
		B += BB
		C += CC
		D += DD
	}

	return produceOutput(A, B, C, D)
}

func ComputeHash(input []byte) []byte {
	data := prepareInput(input)

	computationBuffer := Buffer{
		0x67452301,
		0xEFCDAB89,
		0x98BADCFE,
		0x10325476,
	}

	chunks := splitInput(data)
	for _, chunk := range chunks{
		var A, B, C, D = computationBuffer.WordA, computationBuffer.WordB, computationBuffer.WordC, computationBuffer.WordD

		words := produceWords(chunk)

		var iterationIndex uint32
		for iterationIndex = 0; iterationIndex < 64; iterationIndex++ {

			var F, g uint32
			if iterationIndex <= 15 {
				F = (B & C) | (^B & D)
				g = iterationIndex
			} else if iterationIndex <= 31 {
				F = (D & B) | (^D & C)
				g = (5*iterationIndex + 1) % 16
			} else if iterationIndex <= 47 {
				F = B ^ C ^ D
				g = (3*iterationIndex + 5) % 16
			} else {
				F = C ^ (B | (^D))
				g = (7 * iterationIndex) % 16
			}

			K := calculateSine(iterationIndex + 1)

			A += F + K + words[g]

			dTemp := D
			D = C
			C = B
			B += leftRotate(A, shifts[iterationIndex])
			A = dTemp
		}

		computationBuffer.WordA += A
		computationBuffer.WordB += B
		computationBuffer.WordC += C
		computationBuffer.WordD += D

	}

	return computationBuffer.toByteArray()
}


func splitInput(data []byte) [][]byte {
	chunkCount := len(data) / 64

	var chunks = make([][]byte, chunkCount)

	for index := 0; index < chunkCount; index += 1{
		chunks[index] =	data[index * 64 : (index + 1) * 64]
	}

	return chunks
}

func produceWords(chunk []byte) [16]uint32 {
	var words [16]uint32
	for j := 0; j < 16; j++ {
		words[j] = byteArrayToUint32(chunk[j*4 : (j+1)*4])
	}

	return words
}



func calculatePaddingBytesCount(inputLength int) int {
	quotient := inputLength / 64
	return (64 * quotient) + 56 - inputLength
}

func encodeInputLength(length int64) [8]byte {
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

	lengthBytes := encodeInputLength(int64(len(input) * 8))

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

func byteArrayToUint32(bytes []byte) uint32 {

	a := uint32(bytes[0])
	b := uint32(bytes[1]) << 8
	c := uint32(bytes[2]) << 16
	d := uint32(bytes[3]) << 24

	result := a + b + c + d
	return result
}
