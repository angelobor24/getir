package server

import (
	"encoding/json"
	"fmt"
	"getir/handlerMessage"
	"getir/storage"
	"log"
	"net/http"
	"strings"
)

type Server interface {
	StartServer()
}

// use the service field for a dynamic behaviour of client responce handler
type ServerImpl struct {
	service Service
}

func NewServerImpl(service Service) Server {
	serverImpl := ServerImpl{service: service}
	return &serverImpl
}

func (serverImpl *ServerImpl) StartServer() {

	http.HandleFunc("/database/retrieve", serverImpl.retrieveDB) // GET
	http.HandleFunc("/memory", serverImpl.handleMemory)          // POST/GET

	fmt.Printf("Starting GETIR service at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func (serverImpl *ServerImpl) retrieveDB(w http.ResponseWriter, r *http.Request) {
	var totalData storage.RetrievedFromDB
	if r.URL.Path != "/database/retrieve" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		startDate := strings.Split(r.URL.Query().Get("startDate"), "-")
		endDate := strings.Split(r.URL.Query().Get("endDate"), "-")
		minCount := r.URL.Query().Get("minCount")
		maxCount := r.URL.Query().Get("maxCount")
		// start validation of the input
		if !ValidateDate(startDate) {
			http.Error(w, "Start Data not valid", http.StatusBadRequest)
			return
		}
		if !ValidateDate(endDate) {
			http.Error(w, "End Data not valid", http.StatusBadRequest)
			return
		}
		if !validateCount(minCount) {
			http.Error(w, "minCount Not valid", http.StatusBadRequest)
			return
		}
		if !validateCount(maxCount) {
			http.Error(w, "maxCount not valid", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		retrievedData, err := serverImpl.service.takeFromDB(startDate, endDate, minCount, maxCount)
		if err != nil {
			statusCode, errorMessage, code := handlerMessage.ToStatusCodeMessage(err)
			w.WriteHeader(statusCode)
			totalData.Code = code
			totalData.Message = errorMessage
			json.NewEncoder(w).Encode(totalData)
			return
		}
		totalData.List = retrievedData
		totalData.Message = "Success"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(totalData)
		return
	}
	http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
}

func (serverImpl *ServerImpl) handleMemory(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/memory" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		keyFilter := r.URL.Query().Get("key")
		value, err := serverImpl.service.takeFromMemory(keyFilter)
		if err != nil {
			statusCode, errorMessage, _ := handlerMessage.ToStatusCodeMessage(err)
			http.Error(w, errorMessage, statusCode)
			return
		}
		json.NewEncoder(w).Encode(value)
		return
	}

	if r.Method == "POST" {
		var dbValue storage.InsertInternalDB
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&dbValue)
		if err != nil {
			http.Error(w, "Error on input", http.StatusBadRequest)
			return
		}
		createdValue, err := serverImpl.service.insertInMemory(dbValue.Key, dbValue.Value)
		if err != nil {
			statusCode, errorMessage, _ := handlerMessage.ToStatusCodeMessage(err)
			http.Error(w, errorMessage, statusCode)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdValue)
		return
	}

	http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
}
