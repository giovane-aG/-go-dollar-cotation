package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"cotacao-dolar/server/create_connection"
)

type CotationResponse struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type Cotation struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func GetDollarCotation(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context = context.Background()
	var apiURL string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
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
	var cotationResponse CotationResponse

	err = json.Unmarshal(parsedBody, &cotationResponse)

	ctx2 := context.Background()
	ctx2, cancel2 := context.WithTimeout(ctx2, 500*time.Millisecond)
	defer cancel2()

	newCotation := &Cotation{
		Code:       cotationResponse.Usdbrl.Code,
		Codein:     cotationResponse.Usdbrl.Codein,
		Name:       cotationResponse.Usdbrl.Name,
		High:       cotationResponse.Usdbrl.High,
		Low:        cotationResponse.Usdbrl.Low,
		VarBid:     cotationResponse.Usdbrl.VarBid,
		PctChange:  cotationResponse.Usdbrl.PctChange,
		Bid:        cotationResponse.Usdbrl.Bid,
		Ask:        cotationResponse.Usdbrl.Ask,
		Timestamp:  cotationResponse.Usdbrl.Timestamp,
		CreateDate: cotationResponse.Usdbrl.CreateDate,
	}

	db := create_connection.CreateConnection()
	db.WithContext(ctx2).Create(newCotation)

	apiResponse := fmt.Sprintf("dollar: %s", cotationResponse.Usdbrl.Bid)
	w.Write([]byte(apiResponse))
}

func main() {
	var multiplexer *http.ServeMux = http.NewServeMux()

	multiplexer.HandleFunc("/cotacao", GetDollarCotation)
	http.ListenAndServe(":8080", multiplexer)
}
