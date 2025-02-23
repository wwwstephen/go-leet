package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		expect []int
	}{
		{
			name:   "Basic case",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			expect: []int{0, 1},
		},
		{
			name:   "Negative numbers",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			expect: []int{0, 2},
		},
		{
			name:   "Pair at the end",
			nums:   []int{3, 2, 4},
			target: 6,
			expect: []int{1, 2},
		},
		{
			name:   "No valid pair",
			nums:   []int{1, 2, 3},
			target: 7,
			expect: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSum(tt.nums, tt.target)
			if result != nil {
				sort.Ints(result)
			}

			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("expected %v, got %v", tt.expect, result)
			}
		})
	}
}
