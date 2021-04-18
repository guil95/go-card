package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/guil95/go-card/app/domains/account"
	"github.com/guil95/go-card/infra"
)

//Index is the root api
func listAccounts(service *account.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accounts, err := service.ListAccounts()

		if err == infra.ErrorNotFound {
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
			Document string `json:"document_number" validate:"required"`
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

		if account.Document != "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(NewResponseError("Account exists"))
			return
		}

		accountSaved, err := service.CreateAccount(payload.Document)

		if err != nil {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(NewResponseError("Internal server error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accountSaved)
	})
}

func isValidRequest(payload interface{}) bool {
	var validate *validator.Validate = validator.New()

	err := validate.Struct(payload)

	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

func MakeAccountHandler(r *mux.Router, service *account.Service) {
	r.Handle("/accounts", listAccounts(service)).Methods("GET", "OPTIONS").Name("listAccounts")
	r.Handle("/accounts", saveAccount(service)).Methods("POST", "OPTIONS").Name("saveAccount")

}
