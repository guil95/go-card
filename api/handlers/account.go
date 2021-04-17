package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guil95/go-card/app/domains/account"
	"github.com/guil95/go-card/app/infra"
)

//Index is the root api
func listAccounts(service *account.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accounts, err := service.ListAccounts()

		if err == infra.ErrorNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Accounts not found"))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accounts)
	})
}

func MakeAccountHandler(r *mux.Router, service *account.Service) {
	r.Handle("/accounts", listAccounts(service)).Methods("GET", "OPTIONS").Name("listAccounts")
}
