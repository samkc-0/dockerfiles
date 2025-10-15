package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerPage)
	srv := &http.Server{
		Addr:         ":8899",
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Println("starting server on port 8899")
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
