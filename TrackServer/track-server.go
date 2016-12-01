// track-server
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"./config"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url, error := base64.StdEncoding.DecodeString(strings.TrimPrefix(r.URL.Path, "/"))
		if error == nil {
			var redirectTo string = string(url)
			fmt.Printf("Redirect to: %s. Base64: %s\n", redirectTo, strings.TrimPrefix(r.URL.Path, "/"))
			http.Redirect(w, r, "http://"+redirectTo, http.StatusSeeOther)
		} else {
			fmt.Println(error)
		}
	})

	var tsconfig config.TrackServerConfig
	tsconfig, err := config.RecieveConfig()
	if err != nil {
		panic(1)
	} else {
		fmt.Printf("Listening for port %s\n", tsconfig.Port)
		http.ListenAndServe(tsconfig.Port, nil)
	}
}
