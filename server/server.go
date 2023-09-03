package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"context"

	_ "github.com/mattn/go-sqlite3"
)

type Response struct {
	USDBRL Cotation `json:"USDBRL"`
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

func checkError(err error) {
	if err != nil {
		print(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", getCotationHandler)

	fmt.Println("Servidor iniciado na porta 8080...")
	http.ListenAndServe(":8080", mux)
}

func getCotationHandler(w http.ResponseWriter, r *http.Request) {
	cotation := getCotation()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotation)
}

func saveCotationToDB(cotation *Cotation) {
	db, err := sql.Open("sqlite3", "cotation.db")

	checkError(err)

	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS cotation (
		code TEXT PRIMARY KEY,
		codein TEXT,
		name TEXT,
		high TEXT,
		low TEXT,
		varBid TEXT,
		pctChange TEXT,
		bid TEXT,
		ask TEXT,
		timestamp TEXT,
		create_date TEXT
)`)
	checkError(err)

	stmt, err := db.Prepare("INSERT OR REPLACE INTO cotation (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	checkError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err = stmt.ExecContext(ctx,
		cotation.Code, cotation.Codein, cotation.Name, cotation.High, cotation.Low, cotation.VarBid, cotation.PctChange, cotation.Bid, cotation.Ask, cotation.Timestamp, cotation.CreateDate)

	checkError(err)
}

func getCotation() *Cotation {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	checkError(err)

	res, err := http.DefaultClient.Do(req)
	checkError(err)

	defer res.Body.Close()

	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	checkError(err)

	saveCotationToDB(&response.USDBRL)

	return &response.USDBRL
}
