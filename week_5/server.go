package week_5

import (
	"fmt"
	"net/http"

)

func Run() {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the server!\nAvailable routes: /struct, /csv")
	})

	r.HandleFunc("/struct", convertStructToJSON)
	r.HandleFunc("/csv", handlerCSV)

	fmt.Println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

}
