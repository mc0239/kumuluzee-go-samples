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
			data := struct {
				Value string `json:"value"`
			}{
				stringProp,
			}

			genjson, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(500)
			}

			fmt.Fprint(w, string(genjson))
		} else {
			w.WriteHeader(500)
		}

	})

	configPath := path.Join(".", "config.yaml")

	conf = config.NewUtil(config.Options{
		Extension:  "consul",
		ConfigPath: configPath,
	})

	port, ok := conf.GetInt("kumuluzee.server.http.port")
	if !ok {
		log.Printf("There was an error reading port from configuration")
		port = 9000
	}

	log.Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

}
