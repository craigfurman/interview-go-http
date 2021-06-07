package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/craigfurman/interview-go-http/api"
)

func main() {
	handler := api.New()
	if err := http.ListenAndServe("localhost:8080", handler); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
