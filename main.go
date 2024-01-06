package main

func main() {
	// do nothing
}

// simple-array-sum
func SimpleArraySum(ar []int32) int32 {
	var total int32
	for _, x := range ar {
		total += x
	}
	return total
}

// compare-the-triplets
func CompareTriplets(a []int32, b []int32) []int32 {
	score := []int32{0, 0}
	for i, x := range a {
		if x > b[i] {
			score[0]++
		}
		if x < b[i] {
			score[1]++
		}
	}
	return score
}

// a-very-big-sum
func AVeryBigSum(ar []int64) int64 {
	var total int64
	for _, x := range ar {
		total += x
	}
	return total
}

// solve-me-first
func SolveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func DiagonalDifference(arr [][]int32) int32 {
	n := len(arr)
	var ft int32 // forward total
	var bt int32 // backward total
	for f, b := 0, n-1; f < n; f++ {
		// forward logic
		ft += arr[f][f]
		bt += arr[f][b]
		b--
	}
	return AbsDiff(ft, bt)
}

// used for diagonal-difference
func AbsDiff(a int32, b int32) int32 {
	c := a - b
	if c < 0 {
		return c * -1
	}
	return c
}

// plus-minus [TODO: TESTING ]
func PlusMinus(arr []int32) {
	// Write your code here

}
