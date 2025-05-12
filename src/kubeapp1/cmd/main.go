package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		slog.Error("Application finished with an error", "error", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	slog.Info("Received request to /hello endpoint")

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello World!")
}
