package week_1

import (
	common "GolangBootcamp/common"
	"fmt"
)

func SelectResult() {
	Numbers := randomNumberForSlice()
	common.PrintByTime(Numbers)
	var practice int
	fmt.Print("Write Number of practice 1-Average 2-Sum 3-Average&Sum 4-Maximum 5-Minimum  0-All :")
	fmt.Scan(&practice)

	switch practice {

	case 1:
		av := averageSlice(Numbers)
		common.PrintByTime(av)

	case 2:
		sum := sumSlice(Numbers)
		common.PrintByTime(sum)

	case 3:
		sum, average := sumAndAverageSlice(Numbers)
		common.PrintByTime(sum, average)

	case 4:
		max := maximumNumberInSlice(Numbers)
		common.PrintByTime(max)

	case 5:
		min := minimumNumberInSlice(Numbers)
		common.PrintByTime(min)

	case 0:
		AllWeek1()
	}
}

func AllWeek1() {
	var Numbers []int

	Numbers = randomNumberForSlice()
	common.PrintByTime(Numbers)

	Average := averageSlice(Numbers)
	common.PrintByTime(Average)

	Sum := sumSlice(Numbers)
	common.PrintByTime(Sum)

	min := minimumNumberInSlice(Numbers)
	max := maximumNumberInSlice(Numbers)
	common.PrintByTime(min, max)
	/*
		minBySlices := slices.Min(Numbers)
		common.PrintByTime(minBySlices)

		maxBySlices := slices.Max(Numbers)
		common.PrintByTime(maxBySlices)
	*/

}
