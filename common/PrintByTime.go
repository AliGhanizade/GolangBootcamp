package common

import (
	"fmt"
	"time"
)

// uniq print by time
func PrintByTime(value any) {
	fmt.Printf("in [%v] => %v\n", time.Now().Format(time.TimeOnly), value)
}
