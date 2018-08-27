package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/mc0239/kumuluzee-go-config/config"
)

var conf config.Util

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		stringProp, ok := conf.GetString("rest-config.string-property")
		if ok {
			fmt.Fprintf(w, "{value: %s}", stringProp)
		} else {
			w.WriteHeader(500)
		}
	})

	configPath := path.Join(".", "config.yaml")

	conf = config.NewUtil(config.Options{
		Extension:  "consul",
		ConfigPath: configPath,
	})

	log.Fatal(http.ListenAndServe(":9000", nil))

}
