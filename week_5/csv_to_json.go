package week_5

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func csvToJson(w http.ResponseWriter, r *http.Request) {
	fileName := "../week_5/data.csv"

	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer csvFile.Close()

	csvrecord, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var persons []map[string]string
	for _, record := range csvrecord {
		person := map[string]string{
			"name":        record[0],
			"age":         record[1],
			"city":        record[2],
			"phoneNumber": record[3],
		}
		persons = append(persons, person)
	}

	jsonData, err := json.Marshal(persons)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonFile, err := os.OpenFile("../week_5/data.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	fmt.Println("CSV data successfully converted to JSON and saved to data.json")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
