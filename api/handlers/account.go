package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guil95/go-card/app/domains/account"
	entities "github.com/guil95/go-card/app/entities/account"
	"github.com/guil95/go-card/app/vo/uuid"
)

//Index is the root api
func listAccounts(service *account.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accounts, err := service.ListAccounts()

		if err == entities.ErrorAccountNotFound {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(NewResponseError("Accounts not found"))
			return
		}

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(NewResponseError("Internal server error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accounts)
	})
}

//Index is the root api
func saveAccount(service *account.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Document             string  `json:"document_number" validate:"required"`
			AvailableCreditLimit float64 `json:"available_credit_limit" validate:"required,min=0.1"`
		}

		err := json.NewDecoder(r.Body).Decode(&payload)

		if err != nil {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(NewResponseError("Internal server error"))
			return
		}

		if !isValidRequest(payload) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(NewResponseError("Unprocessable entity"))
			return
		}

		account, err := service.FindAccountByDocument(payload.Document)

		if err != nil {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(NewResponseError("Internal server error"))
			return
		}

		if account != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(NewResponseError("Account exists"))
			return
		}

		accountSaved, err := service.CreateAccount(payload.Document, payload.AvailableCreditLimit)

		if err != nil {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(NewResponseError("Internal server error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(accountSaved)
	})
}

func findAccount(service *account.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := uuid.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(NewResponseError("Unprocessable entity"))
			return
		}

		account, err := service.FindAccountByID(id)

		if err == entities.ErrorAccountNotFound {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(NewResponseError("Account not found"))
			return
		}

		if err != nil {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(NewResponseError("Internal server error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(account)
	})
}

func MakeAccountHandler(r *mux.Router, service *account.Service) {
	r.Handle("/accounts", listAccounts(service)).Methods("GET", "OPTIONS").Name("listAccounts")
	r.Handle("/accounts/{id}", findAccount(service)).Methods("GET", "OPTIONS").Name("findAccount")
	r.Handle("/accounts", saveAccount(service)).Methods("POST", "OPTIONS").Name("saveAccount")
}
