package router

import (
	"net/http"

	"github.com/HasnathJami/go-myWallet-mux-mongoDb/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {w.Write([]byte("Welcome To MyWallet!"))})
	router.HandleFunc("/login", controllers.UserLogin).Methods("POST")
	router.HandleFunc("/register", controllers.UserRegistration).Methods("POST")
        router.HandleFunc("/credit", controllers.CreditMoney).Methods("POST")
	router.HandleFunc("/debit", controllers.DebitMoney).Methods("POST")
	router.HandleFunc("/report", controllers.UserReport).Methods("GET")

	return router
}
