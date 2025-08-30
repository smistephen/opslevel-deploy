package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	env := os.Getenv("ENV")

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			http.Redirect(w, req, "https://"+req.Host+req.URL.Path, http.StatusMovedPermanently)
		})

		log.Println("starting server on port :80")
		if err := http.ListenAndServe(":80", nil); err != nil {
			log.Fatal(err)
		}
	}()

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello! I am running in %s\n", env)
	})

	log.Println("starting server on port :443")
	if err := http.ListenAndServeTLS(":443","crtfile.crt","keyfile.key", nil); err != nil {
		log.Fatal(err)
	}

}
