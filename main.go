package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	event "github.com/valtyr/shake/pkg/event"
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
	// go influx.InfluxPersistor()
	go websockets.WebsocketBroadcaster(clientManager)

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func receiveEvent(w http.ResponseWriter, r *http.Request) {
	b, readError := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if readError != nil {
		log.Fatal(readError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ev, eventError := event.DecodeJSON(b)
	if eventError != nil {
		log.Fatal(eventError)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queue.PublishEvent(ev)
	w.WriteHeader(http.StatusOK)
}
