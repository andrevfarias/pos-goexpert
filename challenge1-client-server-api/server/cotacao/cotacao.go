package cotacao

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/valyala/fastjson"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func ObterCotacao() (*Cotacao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var p fastjson.Parser
	val, err := p.ParseBytes(body)
	if err != nil {
		return nil, err
	}
	cotJSON := val.Get("USDBRL").String()

	var data Cotacao

	if err = json.Unmarshal([]byte(cotJSON), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *Cotacao) SalvarCotacao() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	db, err := sql.Open("sqlite3", "file::./../db.sqlite?cache=shared")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO cotacoes (id, cotacao) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(uuid.New().String(), c.Bid)
	if err != nil {
		return err
	}
	return nil
}
