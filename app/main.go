package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	go func() {
		metricsMux := http.NewServeMux()
		metricsMux.HandleFunc("/metrtics", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Test")
			return
		})
		metricsServer := http.Server{
			Addr:    ":5000",
			Handler: metricsMux,
		}
		err := metricsServer.ListenAndServe()
		if err != nil {
			log.Fatal("Cannot start metrics server")
		}
	}()

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logs")
		return
	})
	apiServer := http.Server{
		Addr:    ":9102",
		Handler: apiMux,
	}
	err := apiServer.ListenAndServe()
	if err != nil {
		log.Fatal("Cannot start api server")
	}
}