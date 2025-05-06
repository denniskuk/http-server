package main

import (
	"fmt"
	"net/http"
)

func routes() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello from /hello")
    })

    mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Goodbye from /bye")
    })

    return mux
}
