package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/loco-assessment/datastore/inmemory"
	"github.com/loco-assessment/handler"
	"github.com/loco-assessment/service"
	"net/http"
)

func main() {
	db := inmemory.NewInMemoryDatastore()
	svc := service.NewService(db)
	h := handler.NewHandler(svc)

	router := mux.NewRouter()

	router.HandleFunc("/transactionservice/{transaction_id}", h.CreateTransaction).Methods("PUT")
	router.HandleFunc("/transactionservice/transaction/{transaction_id}", h.GetTransaction).Methods("GET")
	router.HandleFunc("/transactionservice/types/{transaction_event}", h.GetAllTransactionEvent).Methods("GET")
	router.HandleFunc("/transactionservice/sum/{transaction_id}", h.GetTransactionSum).Methods("GET")

	// Start the server
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", router)

}
