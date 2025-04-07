package main

import "net/http"

import "log"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Whoa! What's Up?");
		d, err := ioutil.ReadAll(r.Body)
		w.Write([]byte("Whoa! What's Up?"))
	});
	
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World");
		w.Write([]byte("Hello World"))
	});

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World");
		w.Write([]byte("Goodbye World"))
	});

	http.ListenAndServe(":9090", nil)
}