package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	influx "github.com/valtyr/shake/pkg/influx"
	logger "github.com/valtyr/shake/pkg/logger"
	queue "github.com/valtyr/shake/pkg/queue"
	websockets "github.com/valtyr/shake/pkg/websockets"
)

type Event struct {
	Type     string
	Data     string
	UserID   string
	SchoolID string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/event", receiveEvent).Methods("POST")

	clientManager := websockets.NewClientManager()
	go clientManager.Run()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websockets.ServeWs(clientManager, w, r)
	})

	queue.CreateQueue()

	go logger.Logger()
	go influx.InfluxPersistor()
	go websockets.WebsocketBroadcaster(clientManager)

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func receiveEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var e Event
	err := decoder.Decode(&e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	queue.PublishEvent(e)
	w.WriteHeader(http.StatusOK)
}
