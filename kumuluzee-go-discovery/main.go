package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"

	"github.com/mc0239/kumuluzee-go-config/config"
	"github.com/mc0239/kumuluzee-go-discovery/discovery"
)

var disc discovery.Util

func main() {
	// initialize discovery
	configPath := path.Join(".", "config.yaml")

	disc = discovery.New(discovery.Options{
		Extension:  "consul",
		ConfigPath: configPath,
	})

	// register this service to registry
	// note: service parameters are read from configuration file
	_, err := disc.RegisterService(discovery.RegisterOptions{})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {
		// define parameters of the service we are looking for
		// and call DiscoverService
		serviceURL, err := disc.DiscoverService(discovery.DiscoverOptions{
			Value:       "test-service",
			Version:     "1.0.0",
			Environment: "dev",
			AccessType:  "direct",
		})
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
		} else {
			// prepare a struct for marshalling into json
			data := struct {
				Service string `json:"service"`
			}{
				serviceURL,
			}

			// generate json from data
			genjson, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(500)
			} else {
				// write generated json to ResponseWriter
				fmt.Fprint(w, string(genjson))
			}
		}
	})

	// initialize configuration
	conf := config.NewUtil(config.Options{
		Extension:  "consul",
		ConfigPath: configPath,
	})

	// perform service deregistration on received interrupt or terminate signals
	deregisterOnSignal()

	// get port number from configuration
	port, ok := conf.GetInt("kumuluzee.server.http.port")
	if !ok {
		log.Printf("Error reading port from configuration")
		port = 9000
	}

	// run server
	log.Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func deregisterOnSignal() {
	// catch interrupt or terminate signals and send them to sigs channel
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// function waits for received signal - and then performs service deregistration
	go func() {
		<-sigs
		if err := disc.DeregisterService(); err != nil {
			panic(err)
		}
		// besides deregistration, you could also do any other clean-up here.
		// Make sure to call os.Exit() with status number at the end.
		os.Exit(1)
	}()
}
