package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.PraseForm(); err != nil {
		fmt.fprintf(w, "ParseFrom() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.formValue("name")
	address := r.formValue("address")

	fmt.Fprintf(w, "Name= %\n", name)
	fmt.Fprintf(w, "Address= %\n", address)


}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting a new server at port 8080 \n")
	if err := http.ListenAndServe("; 8080", nil): err != nil {
		log.Fatal(err)
	}
}