package handlers

import (
	"encoding/json"
	"main/internal/tools"
	"net/http"

	"github.com/go-chi/chi"
)

func CreateRunescapeAccount(w http.ResponseWriter, r *http.Request) {
	var account tools.RunescapeAccount
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad request."))
		return
	}

	tools.DB.Create(&account)
	w.WriteHeader(201)
}

func GetRunescapeAccount(w http.ResponseWriter, r *http.Request) {
	var account tools.RunescapeAccount
	tools.DB.First(&account)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(account)
}

func GetRunescapeAccountByName(w http.ResponseWriter, r *http.Request) {
	var account tools.RunescapeAccount
	var name string = chi.URLParam(r, "name")
	if name == "" {
		w.WriteHeader(400)
		w.Write([]byte("Bad request."))
		return
	}

	tools.DB.First(&account, tools.RunescapeAccount{Name: name})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(account)
}
