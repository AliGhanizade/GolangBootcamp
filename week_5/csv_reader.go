package week_5

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV()  {
	fileName := "../week_5/data.csv"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	csv, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, record := range csv {
		fmt.Printf("Name: %v \tAge: %v \tcity: %v \n%v\n",record[0], record[1], record[2],record[3])
	}
}