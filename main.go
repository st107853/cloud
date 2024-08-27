package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var logger TransactionLogger

func main() {

	err := initializeTransactionLog()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", keyValueGetHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", keyValueDeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func keyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.WriteDelete(key)
	w.WriteHeader(http.StatusOK)
}

func keyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := Get(key)
	if errors.Is(err, ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.WriteGet(key)
	w.Write([]byte(value))
}

func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.WritePut(key, string(value))
	w.WriteHeader(http.StatusCreated)
}

func initializeTransactionLog() error {
	var err error

	logger, err = NewFileTransactionLogger("transaction.log")
	if err != nil {
		return fmt.Errorf("field to create event logger: %w", err)
	}

	events, errors := logger.ReadEvents()
	e, ok := Event{}, true

	for ok && err == nil {
		select {
		case err, ok = <-errors: // Retrieve any errors
		case e, ok = <-events:
			switch e.EventType {
			case EventDelete: // Got a DELETE event
				err = Delete(e.Key)
			case EventPut: // Got a PUT event
				err = Put(e.Key, e.Value)
			case EventGet: // Got a GET event
				_, err = Get(e.Key)
			}

		}
	}

	logger.Run()

	return err
}
