package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guil95/go-card/app/domains/transaction"
	entities "github.com/guil95/go-card/app/entities/transaction"
	"github.com/guil95/go-card/app/utils"
)

func makeTransaction(service *transaction.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			AccountID     utils.ID               `json:"account_id" validate:"required"`
			OperationType entities.OperationType `json:"operation_type_id" validate:"required,oneof=1 2 3 4"`
			Amount        float64                `json:"amount" validate:"required,min=0.1"`
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

		transaction, err := service.MakeTransaction(payload.AccountID, payload.Amount, payload.OperationType)

		if err != nil {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(NewResponseError("Internal server error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(transaction)
	})
}

func MakeTransactionHandler(r *mux.Router, service *transaction.Service) {
	r.Handle("/transactions", makeTransaction(service)).Methods("POST", "OPTIONS").Name("makeTransaction")
}
