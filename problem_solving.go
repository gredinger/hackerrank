package hackerrank

import (
	"fmt"
	"math"
)

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

// plus-minus
func PlusMinus(arr []int32) {
	var pP, nN, nZ float64
	ar := float64(len(arr))
	for _, x := range arr {
		switch {
		case x > 0:
			pP++
		case x < 0:
			nN++
		default:
			nZ++
		}
	}
	fmt.Printf("%.6f\n%.6f\n%.6f", pP/ar, nN/ar, nZ/ar)
}

// staircase
func Starcase(n int32) {
	for i := 1; i <= int(n); i++ {
		for x := 0; x < int(n)-i; x++ {
			fmt.Print(" ")
		}
		for y := 0; y < i; y++ {
			fmt.Print("#")
		}
		if i != int(n) {
			fmt.Println()
		}
	}
}

// climbing-the-leaderboard
func ClimbingLeaderboard(ranked []int32, player []int32) []int32 {
	// sorted rank;
	sRanked := removeDuplicate(ranked)
	var playerPos []int32
	hwp := len(player) / 2
	for pc, p := range player {
		hs := false
		if pc > hwp { // start at begining of list
			for i, s := range sRanked {
				if p >= s {
					playerPos = append(playerPos, int32(i+1))
					hs = true
					break // found best score, leave
				}
			}
		} else { // start at end of list
			var pr int32
			for i := len(sRanked) - 1; i >= 0; i-- {
				if p >= sRanked[i] {
					pr = int32(i + 1)
					hs = true
				}
				if p <= sRanked[i] {
					break // not gonna find a better score
				}
			}
			if pr > 0 {
				playerPos = append(playerPos, pr)
			}
		}
		if !hs {
			playerPos = append(playerPos, int32(len(sRanked)+1))
		}
	}
	return playerPos
}

// stolen from https://stackoverflow.com/questions/66643946/how-to-remove-duplicates-strings-or-int-from-slice-in-go
func removeDuplicate[T int32](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// birthday cake candles
func BirthdayCakeCandles(candles []int32) int32 {
	lc := 0
	ln := int32(0)
	for _, x := range candles {
		if x > ln {
			lc = 0
			ln = x
		}
		if x == ln {
			lc++
		}
	}
	return int32(lc)
}

// mini-max-sum
func MiniMaxSum(arr []int32) {
	highest := int32(0)
	lowest := int32(math.Pow(2, 30))
	total := int64(0)
	for _, x := range arr {
		if x > highest {
			highest = x
		}
		if x < lowest {
			lowest = x
		}
		total += int64(x)
	}
	fmt.Printf("%v %v", total-int64(highest), total-int64(lowest))

}

// grading
func GradingStudents(grades []int32) (rGrades []int32) {
	for _, x := range grades {
		dR := x % 5 //delta round value
		if dR >= 3 && x >= 38 {
			rGrades = append(rGrades, (5-dR)+x)
			continue
		}
		rGrades = append(rGrades, x)
	}
	return rGrades
}
