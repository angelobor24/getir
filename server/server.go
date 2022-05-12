package server

import (
	"encoding/json"
	"fmt"
	"getir/handlerMessage"
	"log"
	"net/http"
)

type Server interface {
	StartServer()
}

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

	fmt.Printf("Starting insurance pokemon service at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// handler for POST request to /trainer endpoint
func (serverImpl *ServerImpl) retrieveDB(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/database/retrieve" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		param1 := r.URL.Query().Get("test")
		fmt.Println(param1)
		//trainer := TrainerInfo{}
		//json.NewDecoder(r.Body).Decode(&trainer)
		err := serverImpl.service.takeFromDB()
		if err != nil {
			statusCode, errorMessage := handlerMessage.ToStatusCodeMessage(err)
			http.Error(w, errorMessage, statusCode)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		//json.NewEncoder(w).Encode()
		return
	}

	http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
}

func (serverImpl *ServerImpl) handleMemory(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/memory" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		trainer := TrainerInfo{}
		json.NewDecoder(r.Body).Decode(&trainer)
		err := serverImpl.service.takeFromMemory()
		if err != nil {
			statusCode, errorMessage := handlerMessage.ToStatusCodeMessage(err)
			http.Error(w, errorMessage, statusCode)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		//json.NewEncoder(w).Encode(resourceCreated)
		return
	}

	if r.Method == "POST" {
		trainer := TrainerInfo{}
		json.NewDecoder(r.Body).Decode(&trainer)
		err := serverImpl.service.insertInMemory()
		if err != nil {
			statusCode, errorMessage := handlerMessage.ToStatusCodeMessage(err)
			http.Error(w, errorMessage, statusCode)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		//json.NewEncoder(w).Encode(resourceCreated)
		return
	}

	http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
}
