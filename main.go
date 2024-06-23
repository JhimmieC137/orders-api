package main

import (
	"fmt"
	"net/http"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goodbye cruel world!"))
}

func main() {
	server := &http.Server{
		Addr:    ":7000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to  server", err)
	}
}
