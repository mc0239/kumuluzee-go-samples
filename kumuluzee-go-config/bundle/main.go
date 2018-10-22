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

type myConfig struct {
	StringProperty  string `config:"string-property,watch"`
	IntegerProperty int    `config:"integer-property,watch"`
	BooleanProperty bool   `config:"boolean-property,watch"`
	ObjectProperty  struct {
		SubProperty  string `config:"sub-property,watch"`
		SubProperty2 string `config:"sub-property-2"`
	} `config:"object-property"`
}

var conf myConfig

func main() {

	prefixKey := "rest-config"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Value    string `json:"value"`
			Subvalue string `json:"subvalue"`
		}{
			conf.StringProperty,
			conf.ObjectProperty.SubProperty,
		}

		genjson, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(500)
		} else {
			fmt.Fprint(w, string(genjson))
		}

	})

	configPath := path.Join(".", "config.yaml")

	opts := config.Options{
		Extension:  "consul",
		ConfigPath: configPath,
	}

	config.NewBundle(prefixKey, &conf, opts)

	util := config.NewUtil(opts)

	port, ok := util.GetInt("kumuluzee.server.http.port")
	if !ok {
		log.Printf("There was an error reading port from configuration")
		port = 9000
	}

	log.Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

}
