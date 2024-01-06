package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

// simple-array-sum
func TestSimpleArraySum(t *testing.T) {
	ar := []int32{1, 2, 3, 4, 10, 11}
	if SimpleArraySum(ar) != 31 {
		t.Fail()
	}
	ar = []int32{1, 2, 3}
	if SimpleArraySum(ar) != 6 {
		t.Fail()
	}
}

func TestCompareTriplets(t *testing.T) {
	inA := []int32{5, 6, 7}
	inB := []int32{3, 6, 10}
	out := []int32{1, 1}
	test := CompareTriplets(inA, inB)
	for i, x := range out {
		if x != test[i] {
			t.Errorf("%v is not equal to %v", x, test[i])
		}
	}
}

// generated test case much better
func TestAVeryBigSum(t *testing.T) {
	type args struct {
		ar []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test",
			args: args{ar: []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}},
			want: 5000000015,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AVeryBigSum(tt.args.ar); got != tt.want {
				t.Errorf("AVeryBigSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// generated
func TestSolveMeFirst(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "test1",
			args: args{
				a: 2,
				b: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveMeFirst(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("SolveMeFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

// generated
func TestDiagonalDifference(t *testing.T) {
	type args struct {
		arr [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test1",
			args: args{
				arr: [][]int32{
					{11, 2, 4},
					{4, 5, 6},
					{10, 8, -12},
				},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiagonalDifference(tt.args.arr); got != tt.want {
				t.Errorf("DiagonalDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

// generated w/ modifications
// need to figure out how to monitor fmt.Println('') output from a function call

func TestPlusMinus(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{[]int32{-4, 3, -9, 0, 4, 1}},
			want: `0.500000
			0.333333
			0.166667`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := captureOutput(func() { PlusMinus(tt.args.arr) }); strings.EqualFold(got, tt.want) {
				t.Errorf("PlusMinus() = %v, want %v", got, tt.want)
			}
		})
	}
}

// stolen from https://stackoverflow.com/questions/26804642/how-to-test-a-functions-output-stdout-stderr-in-unit-tests
func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
