package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loco-assessment/models"
	"github.com/loco-assessment/service"
	"net/http"
	"strconv"
)

type handler struct {
	svc service.Service
}

func NewHandler(svc service.Service) *handler {
	return &handler{svc: svc}
}

func (h *handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	w.Header().Set("Content-Type", "application/json") // Set the content type to json
	var response map[string]string

	vars := mux.Vars(r)
	id := vars["transaction_id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)

		response = map[string]string{
			"message": "Empty transaction id",
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	txnId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{
			"err": err.Error(),
		}

		json.NewEncoder(w).Encode(response)

		return

	}

	// Parse the body and call the service
	var txn models.Transaction

	err = json.NewDecoder(body).Decode(&txn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{
			"err": err.Error(),
		}

		json.NewEncoder(w).Encode(response)

		return
	}

	err = h.svc.CreateTransaction(txnId, txn.Amount, txn.TransactionType, txn.ParentTransactionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response = map[string]string{
			"err": err.Error(),
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)

	response = map[string]string{
		"message": "Transaction created successfully",
	}

	json.NewEncoder(w).Encode(response)

	return
}

func (h *handler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the content type to json
	var response map[string]string

	vars := mux.Vars(r)
	id := vars["transaction_id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{
			"message": "Empty transaction id",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	txnId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{
			"err": err.Error(),
		}

		json.NewEncoder(w).Encode(response)
		return

	}

	// Call the service
	txn, err := h.svc.GetTransaction(txnId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response = map[string]string{
			"err": err.Error(),
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(txn)
}

func (h *handler) GetTransactionSum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the content type to json
	var response map[string]interface{}

	vars := mux.Vars(r)
	id := vars["transaction_id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]interface{}{
			"message": "Empty transaction id",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	txnId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]interface{}{
			"err": err.Error(),
		}

		json.NewEncoder(w).Encode(response)
		return

	}

	// Call the service
	sum, err := h.svc.GetTransactionSum(txnId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Write the response
	w.WriteHeader(http.StatusOK)

	response = map[string]interface{}{
		"sum": sum,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetAllTransactionEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the content type to json
	var response map[string]string

	vars := mux.Vars(r)
	event := vars["transaction_event"]
	if event == "" {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{
			"message": "Empty transaction event",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// Call the service
	ids, err := h.svc.GetAllTransactionEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response = map[string]string{
			"err": err.Error(),
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ids)
}
