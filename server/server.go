package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func GetDollarCotation(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context = context.Background()
	var apiURL string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)

	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	parsedBody, err := io.ReadAll(response.Body)
	w.Write(parsedBody)
}

func main() {
	var multiplexer *http.ServeMux = http.NewServeMux()

	multiplexer.HandleFunc("/cotacao", GetDollarCotation)
	http.ListenAndServe(":8080", multiplexer)
}
