package week_1

import common "GolangBootcamp/common"

func AllWeek1() {
	var Numbers []int

	Numbers = RandomNumberForSlice(Numbers)
	common.PrintByTime(Numbers)

	Average := AverageSlice(Numbers)
	common.PrintByTime(Average)

	Sum := SumSlice(Numbers)
	common.PrintByTime(Sum)

	min := MinimumNumberInSlice(Numbers)
	max := MaximumNumberInSlice(Numbers)
	common.PrintByTime(min,max)
	/*
		minBySlices := slices.Min(Numbers)
		common.PrintByTime(minBySlices)

		maxBySlices := slices.Max(Numbers)
		common.PrintByTime(maxBySlices)
	*/

}
