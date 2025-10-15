package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerPage)
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Println("starting server on port " + port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlerPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	page := `<html>
	<head><title>INDEX</title></head>
	<body>
	<p> Hello from Docker! I'm a Go server. </p></body>
	</html>
`
	w.Write([]byte(page))
}
