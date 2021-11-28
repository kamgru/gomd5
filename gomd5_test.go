package main

import (
	"fmt"
	"testing"
)

func Test_ComputeHash(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	tests := []testCase{
		{"What I cannot create, I do not understand.", "e28bee9858753d26209fb95415257f33"},
		{"The quick brown fox jumps over the lazy dog", "9e107d9d372bb6826bd81d3542a419d6"},
		{"123456", "e10adc3949ba59abbe56e057f20f883e"},
		{"abc", "900150983cd24fb0d6963f7d28e17f72"},
		{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus non imperdiet dui. " +
			"Cras egestas gravida ultricies. Praesent tellus est, sagittis at odio in, lacinia " +
			"imperdiet mauris. Aliquam fringilla maximus ultrices. Morbi maximus ut neque vel " +
			"commodo. Nunc arcu lorem, cursus at enim a, bibendum accumsan nibh. Maecenas eget " +
			"aliquam dui, in dictum turpis. Vestibulum vitae viverra nulla, et posuere risus. " +
			"Nullam in gravida leo. In rutrum quam et consequat aliquam. In ultrices, lacus luctus " +
			"mollis aliquet, neque nisl scelerisque orci, eget eleifend erat justo elementum risus.",
			"31d4a24afe8ad2098133cd38ec992519"},
	}

	for _, test := range tests {
		actual := fmt.Sprintf("%x", ComputeHashRfc([]byte(test.input)))
		if actual != test.expected {
			t.Errorf("expected")
		}
	}
}

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

		actual := encodeInputLength(test.input)

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
		{input: []byte{11, 0, 0, 0}, expected: 11},
		{input: []byte{202, 168, 0, 0}, expected: 43210},
		{input: []byte{97, 98, 99, 128}, expected: 2153996897},
	}

	for _, test := range tests {
		actual := byteArrayToUint32(test.input)

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
