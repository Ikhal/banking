package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func Start()  {
	router := mux.NewRouter()
	//define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Post request received")
}

func getCustomer(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}