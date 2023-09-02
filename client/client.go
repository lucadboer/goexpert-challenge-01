package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

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

func main() {
	cotation := getCotation()

	bid := cotation.Bid

	saveCotation(bid)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getCotation() *Cotation {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)

	checkError(err)

	res, err := http.DefaultClient.Do(req)

	checkError(err)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	checkError(err)

	var cotation Cotation

	err = json.Unmarshal(body, &cotation)

	checkError(err)

	return &cotation
}

func saveCotation(bid string) {
	file, err := os.Create("cotacao.txt")

	checkError(err)

	defer file.Close()

	_, err = file.WriteString("Dólar: " + bid)

	checkError(err)

	fmt.Println("Valor de bid salvo em cotacao.txt com sucesso.")
}
