package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// サーバーの起動
	err := http.ListenAndServe(
		":18080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", string(r.URL.Path[1:]))
		}),
	)
	if err != nil {
		fmt.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
