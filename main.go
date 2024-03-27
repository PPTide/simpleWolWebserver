package main

import (
	_ "embed"
	"fmt"
	"github.com/mdlayher/wol"
	"net"
	"net/http"
	"os"
	"strings"
)

//go:embed index.html
var index string

func main() {
	wClient, err := wol.NewClient()
	if err != nil {
		panic(err)
	}

	var port = "8080"
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == "port" {
			port = pair[1]
		}
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			mac, err := net.ParseMAC(request.FormValue("mac"))
			if err != nil {
				fmt.Fprintf(writer, "Error Parsing Mac-Address: %v", err)
				return
			}
			err = wClient.Wake("255.255.255.255:9", mac)
			if err != nil {
				fmt.Fprintf(writer, "Error Waking: %v", err)
				return
			}
			fmt.Fprintf(writer, "Succesfully tried to wake %s", mac)
		} else {
			fmt.Fprintf(writer, index)
		}
	})

	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		panic(err)
	}
}
