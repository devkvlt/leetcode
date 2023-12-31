package main

import (
	"slices"
	"testing"
)

var tests = []struct {
	name string
	nums []int
	want []int
}{
	{"test 1", []int{2, 3, 5}, []int{4, 3, 5}},
	{"test 2", []int{1, 4, 6, 8, 10}, []int{24, 15, 13, 15, 21}},
}

func TestGetSumAbsoluteDifferences(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSumAbsoluteDifferences(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Errorf("getSumAbsoluteDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}
