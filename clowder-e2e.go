package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	client "github.com/redhatinsights/app-common-go/pkg/api/v1"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Args: %v", os.Args)
	fmt.Fprintln(w, msg)
}

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("Hi")
	} else {
		mux := http.NewServeMux()
		port := *client.LoadedConfig.PublicPort
		mux.HandleFunc("/", helloHandler)
		mux.HandleFunc("/healthz", helloHandler)
		mux.HandleFunc("/api/puptoo/", helloHandler)

		address := fmt.Sprintf(":%d", port)

		server := http.Server{
			Addr:    address,
			Handler: mux,
		}

		outputChannel := make(chan error)

		go func() {
			fmt.Printf("Started serving base service at %s\n", address)
			err := server.ListenAndServe()
			if err != nil {
				outputChannel <- err
			}
		}()

		mux2 := http.NewServeMux()
		mux2.Handle(client.LoadedConfig.MetricsPath, promhttp.Handler())

		server2 := http.Server{
			Addr:    fmt.Sprintf(":%d", client.LoadedConfig.MetricsPort),
			Handler: mux2,
		}
		go func() {
			fmt.Printf("Started serving metrics service at :%d%s\n", client.LoadedConfig.MetricsPort, client.LoadedConfig.MetricsPath)
			err := server2.ListenAndServe()
			if err != nil {
				outputChannel <- err
			}
		}()

		err := <-outputChannel
		fmt.Printf("got err %s", err.Error())
	}
}
