package main

import (
	"log"
	"fmt"

	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"strings"
	"time"

	"github.com/dmitryk-dk/raspberry/server/arduino"
)

const (
	staticDir = "./build/"
	listen    = ":3000"
)

type Weather struct {
	Temperature string `json:"temperature"`
	Humidity    string `json:"humidity"`
}

func dependenciesHandler() http.Handler {
	return http.StripPrefix("/", http.FileServer(http.Dir(staticDir)))
}

func main() {
	// get static files
	depHandler := dependenciesHandler()

	// handle static files
	http.Handle("/", depHandler)

	// handle get request
	http.HandleFunc("/temperature", getData)

	prepareShutdown()
	fmt.Printf("Running server on port: %s\n Type Ctr-c to shutdown server.\n", listen)
	http.ListenAndServe(listen, nil)
}

func prepareShutdown() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Got signal %d", <-sig)
		os.Exit(0)
	}()
}

func getData(w http.ResponseWriter, r *http.Request) {
	ard := new(arduino.Arduino)
	if r.Method == "GET" {

		if err := ard.Connect(); err != nil {
			log.Fatalf("can't connect to Arduino: %s: %s", arduino.ArduinoSerialName, err)
		}

		n, err := ard.SendCommand()
		if err != nil {
			log.Fatalf("has problem when sending command to arduino: %s: %s", arduino.ArduinoSerialName, err)
		}
		time.Sleep(1 * time.Second)
		if n > 0 {
			_, buf, err := ard.GetData()
			if err != nil {
				log.Fatalf("can't get data from arduino: %s: %s", arduino.ArduinoSerialName, err)
			}

			weather := Weather{
				Temperature: string(buf[:5]) + " *C",
				Humidity:    strings.TrimSuffix(string(buf[6:12]), "\r") + " %",
			}

			jsonData, err := json.Marshal(weather)

			if err != nil {
				log.Println("Marshal error:", err)
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, "Wrong data", http.StatusInternalServerError)
			}

			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		} else {
			log.Println("Can't get data because 0 bytes received", err)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "Can't get data because 0 bytes received", http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Used wrong method", http.StatusInternalServerError)
	}
}
