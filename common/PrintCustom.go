package common

import (
	"fmt"
	"time"
)

// uniq print any item by time
func PrintByTime(value ...any) {
	fmt.Printf("in [%v] => %v\n", time.Now().Format(time.TimeOnly), value)
}

func PrintMapString(item map[string]string)  {
	for key, value := range item {
		fmt.Printf("key : %v\t=>\tvalue : %v \n",key,value)
	}
}
func PrintMapStringJustKey(item map[string]string)  {
	for key, _ := range item {
		fmt.Printf("key : %v\n",key)
	}
}