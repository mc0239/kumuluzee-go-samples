package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/mc0239/kumuluzee-go-config/config"
)

type myConfig struct {
	StringProperty  string `config:"string-property"`
	IntegerProperty int    `config:"integer-property"`
	BooleanProperty bool   `config:"boolean-property"`
	ObjectProperty  struct {
		SubProperty  string `config:"sub-property"`
		SubProperty2 string `config:"sub-property-2"`
	} `config:"object-property"`
}

var conf myConfig

func main() {

	prefixKey := "rest-config"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{ \"value\": \"%s\", \"subvalue\": \"%s\"}", conf.StringProperty, conf.ObjectProperty.SubProperty)
	})

	configPath := path.Join(".", "config.yaml")

	config.NewBundle(prefixKey, &conf, config.Options{
		Extension:  "consul",
		ConfigPath: configPath,
	})

	log.Fatal(http.ListenAndServe(":9000", nil))

}
