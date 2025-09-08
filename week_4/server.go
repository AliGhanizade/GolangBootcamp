package week_4

import (
	"fmt"
	"net/http"
	"time"
)

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/test", test)
	mux.HandleFunc("/time", showTime)

	server1 := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Millisecond * 10,
		WriteTimeout: time.Microsecond * 10,
		Handler:      mux,
	}
	err := server1.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
func showTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "time is : %s", time.Now().Format(time.DateTime))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprint(w, "hello")

}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	fmt.Fprintf(w, "Test in %s method", r.Method)
}
