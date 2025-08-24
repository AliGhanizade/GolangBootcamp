package week_1

import common "GolangBootcamp/common"

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
