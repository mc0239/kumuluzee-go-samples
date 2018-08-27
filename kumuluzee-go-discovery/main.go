package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/mc0239/kumuluzee-go-discovery/discovery"
)

var disc discovery.Util

func main() {

	configPath := path.Join(".", "config.yaml")

	disc = discovery.New(discovery.Options{
		Extension:  "consul",
		ConfigPath: configPath,
	})

	_, err := disc.RegisterService(discovery.RegisterOptions{})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {
		service, err := disc.DiscoverService(discovery.DiscoverOptions{
			Value:       "test-service",
			Version:     "1.0.0",
			Environment: "dev",
		})
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
		} else {
			fmt.Fprintf(w, "{ \"service\": \"%s:%d\" }", service.Address, service.Port)
		}
	})

	log.Fatal(http.ListenAndServe(":9000", nil))

}
