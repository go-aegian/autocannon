package pivot

import (
	"log"
	"testing"
)

var testCases = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{2, 4, 5, 1, 2, 3},
		expected: 2,
	},
	{
		input:    []int{1, 1, 1, 1, 1, 1},
		expected: -1,
	},
	{
		input:    []int{1, 1, 1, 1, 1},
		expected: 2,
	},
	{
		input:    []int{11, 4, 5, 1, 2, 3},
		expected: 1,
	},
	{
		input:    []int{11, 4, 11},
		expected: 1,
	},
	{
		input:    []int{11, 4},
		expected: -1,
	},
	{
		input:    []int{11},
		expected: -1,
	},
	{
		input:    []int{},
		expected: -1,
	},
}

func TestPivotPoint(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run("", func(t *testing.T) {

			log.Printf("Finding pivot for %v", tc.input)

			result := FindPivotPoint(tc.input, true)

			if result == tc.expected {
				log.Printf("PASS - Expected %d and got %d\n\n", tc.expected, result)
			} else {
				log.Printf("FAIL - Expected %d got %d\n\n", tc.expected, result)
				t.Fail()
			}
		})
	}
}
