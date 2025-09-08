package week_5

import (
	"GolangBootcamp/common"
	"encoding/json"
	"net/http"
)

func convertStructToJSON(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to convert structs to JSON
	persons := createStructs()
	jsonData, err := json.Marshal(persons)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func createStructs() []common.Person {
	persons := []common.Person{
		{ID: 1, Name: "Ali", Age: 17, PhoneNumber: "989052315520", Email: "mraliGhanizade@gmail.com"},
		{ID: 2, Name: "Reza", Age: 18, PhoneNumber: "989052315521", Email: ""},
		{ID: 3, Name: "Sara", Age: 19, PhoneNumber: "989052315522", Email: ""},
		{ID: 4, Name: "Ali", Age: 17, PhoneNumber: "989052315520", Email: "mraliGhanizade@gmail.com"},
		{ID: 5, Name: "Reza", Age: 18, PhoneNumber: "989052315521", Email: ""},
		{ID: 6, Name: "Sara", Age: 19, PhoneNumber: "989052315522", Email: ""}, 
		{ID: 7, Name: "Ali", Age: 17, PhoneNumber: "989052315520", Email: "mraliGhanizade@gmail.com"},
		{ID: 8, Name: "Reza", Age: 18, PhoneNumber: "989052315521", Email: ""},
		{ID: 9, Name: "Sara", Age: 19, PhoneNumber: "989052315522", Email: ""}, 
		{ID: 10, Name: "Ali", Age: 17, PhoneNumber: "989052315520", Email: "mraliGhanizade@gmail.com"},
		{ID: 11, Name: "Reza", Age: 18, PhoneNumber: "989052315521", Email: ""},
		{ID: 12, Name: "Sara", Age: 19, PhoneNumber: "989052315522", Email: ""},
	}
	return persons
}
