package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/imankulov/pt-numbers/converter"
)

func main() {
	http.HandleFunc("/", convert)
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "9090"
	}
	err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func convert(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	io.WriteString(w, converter.Do(n))
}
