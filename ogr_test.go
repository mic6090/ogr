package ogr

import (
	"fmt"
	"testing"
)

func equal(a, b []uint16) bool {
	if len(b) != len(a) {
		return false
	}
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func TestOgrv2(t *testing.T) {
	for _, testCase := range testCases {
		ret := Ogrv2(testCase.length)
		if !equal(ret, testCase.expected) {
			t.Fatalf("Ogrv2: length %d, expected %v, received %v\n", testCase.length, testCase.expected, ret)
		} else {
			t.Logf("PASS: Ogrv2 length %d", testCase.length)
		}
	}
}

func TestOgrv3(t *testing.T) {
	for _, testCase := range testCases {
		ret := Ogrv3(testCase.length)
		if !equal(ret, testCase.expected) {
			t.Fatalf("Ogrv3: length %d, expected %v, received %v\n", testCase.length, testCase.expected, ret)
		} else {
			t.Logf("PASS: Ogrv3 length %d", testCase.length)
		}
	}
}

func BenchmarkOgrv2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ogrv2(9)
	}
}

func BenchmarkOgrv3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ogrv3(9)
	}
}

var testCases = []struct {
	description string
	length      int
	expected    []uint16
}{
	{
		description: "2 level OGR",
		length:      2,
		expected:    []uint16{1},
	},
	{
		description: "3 level OGR",
		length:      3,
		expected:    []uint16{1, 2},
	},
	{
		description: "4 level OGR",
		length:      4,
		expected:    []uint16{1, 3, 2},
	},
	{
		description: "5 level OGR",
		length:      5,
		expected:    []uint16{1, 3, 5, 2},
	},
	{
		description: "6 level OGR",
		length:      6,
		expected:    []uint16{1, 3, 6, 2, 5},
	},
	{
		description: "7 level OGR",
		length:      7,
		expected:    []uint16{1, 3, 6, 8, 5, 2},
	},
	{
		description: "8 level OGR",
		length:      8,
		expected:    []uint16{1, 3, 5, 6, 7, 10, 2},
	},
	{
		description: "9 level OGR",
		length:      9,
		expected:    []uint16{1, 4, 7, 13, 2, 8, 6, 3},
	},
	{
		description: "10 level OGR",
		length:      10,
		expected:    []uint16{1, 5, 4, 13, 3, 8, 7, 12, 2},
	},
}
