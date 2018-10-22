package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/mc0239/kumuluzee-go-config/config"
)

var conf config.Util

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		stringProp, ok := conf.GetString("rest-config.string-property")
		if ok {
			// prepare a struct for marshalling into json
			data := struct {
				Value string `json:"value"`
			}{
				stringProp,
			}

			// generate json from data
			genjson, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(500)
			} else {
				// write generated json to ResponseWriter
				fmt.Fprint(w, string(genjson))
			}
		} else {
			w.WriteHeader(500)
		}

	})

	// initialize configuration
	configPath := path.Join(".", "config.yaml")

	conf = config.NewUtil(config.Options{
		Extension:  "consul",
		ConfigPath: configPath,
	})

	// get port number from configuration aswell
	port, ok := conf.GetInt("kumuluzee.server.http.port")
	if !ok {
		log.Printf("Error reading port from configuration")
		port = 9000
	}

	// run server
	log.Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

}
