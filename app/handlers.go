package app

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"transactions-api/src/accounts"
	"transactions-api/src/transactions"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var a accounts.Account

	rb, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal(rb, &a)

	s, err := accounts.NewService()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := s.Create(a); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	var a accounts.Account

	aId, err := strconv.Atoi(getURLParam(r, "accountId"))
	a.Id = aId

	s, err := accounts.NewService()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	acc, err := s.Get(a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(acc)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var t transactions.Transaction

	rb, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal(rb, &t)

	s, err := transactions.NewService()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := s.Create(t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
