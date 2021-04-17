package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/guil95/go-card/app/infra/repositories"

	"github.com/guil95/go-card/api/handlers"
	"github.com/guil95/go-card/app/domains/account"

	"github.com/gorilla/mux"
)

//IndexApiHandler is index api
func indexApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go card!\n"))
}

func Run(db *sql.DB) {
	r := mux.NewRouter()
	r.HandleFunc("/", indexApiHandler)

	var accountRepository = repositories.NewAccountRepo(db)

	handlers.MakeAccountHandler(r, account.NewService(accountRepository))

	log.Fatal(http.ListenAndServe(":8000", r))
}
