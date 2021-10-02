package main

import (
	"testing"
)

func Test_calculatePaddingBytesCount(t *testing.T) {

	type testCase struct {
		inputLength   int
		expectedCount int
	}

	testCases := []testCase{
		{inputLength: 11, expectedCount: 45},
		{inputLength: 324, expectedCount: 52},
	}

	for _, test := range testCases {
		actualResult := calculatePaddingBytesCount(test.inputLength)
		if actualResult != test.expectedCount {
			t.Errorf("expected %d actual %d", test.expectedCount, actualResult)
		}
	}
}

func Test_encodeLength(t *testing.T) {

	type TestCase struct {
		input    int64
		expected [8]byte
	}

	tests := []TestCase{
		{input: 11, expected: [8]byte{0, 0, 0, 0, 0, 0, 0, 11}},
		{input: 43210, expected: [8]byte{0, 0, 0, 0, 0, 0, 168, 202}},
		{input: 98765432123456789, expected: [8]byte{1, 94, 226, 163, 33, 206, 125, 21}},
	}

	for _, test := range tests {

		actual := encodeLength(test.input)

		if actual != test.expected {
			t.Errorf("failed for input %d\nexpected %v\nactual %v", test.input, test.expected, actual)
		}

	}
}

func Test_mergeInt(t *testing.T) {
	type TestCase struct {
		input    []byte
		expected uint32
	}

	tests := []TestCase{
		{input: []byte{0, 0, 0, 11}, expected: 11},
		{input: []byte{0, 0, 168, 202}, expected: 43210},
	}

	for _, test := range tests {
		actual := mergeInt(test.input)

		if actual != test.expected {
			t.Errorf("expected %d actual %d", test.expected, actual)
		}
	}
}

func Test_getAppendBytes(t *testing.T) {

	type TestCase struct {
		input    []byte
		expected []byte
	}

	tests := []TestCase{
		{
			input: []byte("test case 1"),
			expected: []byte{
				127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 11,
			}},
	}

	for _, test := range tests {

		actual := getAppendBytes(test.input)

		if !areEquivalent(test.expected, actual) {
			t.Errorf("\n%v expected\n%v actual", test.expected, actual)
		}
	}
}

func areEquivalent(left []byte, right []byte) bool {
	if len(left) != len(right) {
		return false
	}

	for i, v := range left {
		if v != right[i] {
			return false
		}
	}

	return true
}
