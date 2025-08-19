package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	var Numbers []int

	Numbers = RandomNumberForSlice(Numbers)
	PrintByTime(Numbers)

	Average := AverageSlice(Numbers)
	PrintByTime(Average)

	Sum := SumSlice(Numbers)
	PrintByTime(Sum)
}

//create random int number (under 100) and append this to int slice  		
func RandomNumberForSlice(Numbers []int)  []int {
	var RandNumbers []int

	for i := 0; i < 10; i++ {
		RandNumbers = append(RandNumbers, rand.Intn(100))
	}

	Numbers = append(Numbers,RandNumbers...)
	return Numbers
}

//Sum & Average
func SumAndAverageSlice(Numbers []int)(int, float32)  {
	var sum int
	var average float32

	for _, v := range Numbers {
	sum += v
	}
	average = float32(sum) / float32(len(Numbers))

	return sum, average
}

//sum int slice and return int
func AverageSlice(Numbers []int) float32{
	var sum int
	var average float32

	for _, v := range Numbers {
	sum += v
	}
	average = float32(sum) / float32(len(Numbers))

	return average

}

// sum int slice and return int
func SumSlice(Numbers []int) int {
	var sum int

	for _, v := range Numbers {
		sum+= v
	}

	return sum
	
}

//uniq print by time
func PrintByTime(value any){
	fmt.Printf("in [%v] => %v\n",time.Now().Format(time.TimeOnly),value)
}