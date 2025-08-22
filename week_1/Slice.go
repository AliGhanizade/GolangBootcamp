package week_1

import (
	"math/rand"
	"slices"
)

// create random int number (under 100) and append this to int slice
func randomNumberForSlice(Numbers []int) []int {
	var RandNumbers []int

	for i := 0; i < 10; i++ {
		RandNumbers = append(RandNumbers, rand.Intn(100))
	}

	Numbers = append(Numbers, RandNumbers...)
	slices.Sort(Numbers)
	return Numbers
}

// Sum & Average
func sumAndAverageSlice(Numbers []int) (int, float32) {
	var sum int
	var average float32

	for _, v := range Numbers {
		sum += v
	}
	average = float32(sum) / float32(len(Numbers))

	return sum, average
}

// sum int slice and return int
func averageSlice(Numbers []int) float32 {
	var sum int
	var average float32

	for _, v := range Numbers {
		sum += v
	}
	average = float32(sum) / float32(len(Numbers))

	return average

}

// sum int slice and return int
func sumSlice(Numbers []int) int {
	var sum int

	for _, v := range Numbers {
		sum += v
	}

	return sum

}

//this func for week2 but for slice and i read this in this file

func minimumNumberInSlice(Numbers []int) int {
	min := 100
	for _, v := range Numbers {
		if v < min {
			min = v
		}
	}
	return min
	// But we can use min:=slices.min(Numbers)
}

func maximumNumberInSlice(Numbers []int) int {
	max := 0
	for _, v := range Numbers {
		if max < v {
			max = v
		}
	}
	return max
	// But we can use min:=slices.max(Numbers)
}
