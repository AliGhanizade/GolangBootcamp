package week_5

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func handlerCSV(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	jsonData := csvToJson()
	if jsonData == nil {
		http.Error(w, "Error processing CSV", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func csvToJson() []byte {
	fileName := "../week_5/data.csv"

	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer csvFile.Close()

	csvrecord, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
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
		return nil
	}
	return jsonData

}

func saveInJsonFile(jsonData []byte) {
	jsonFile, err := os.OpenFile("../week_5/csvTojson.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	fmt.Println("CSV data successfully converted to JSON and saved to data.json")
	return
}
