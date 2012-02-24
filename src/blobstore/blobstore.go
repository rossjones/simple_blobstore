package main

import (
	"flag"
	"fmt"
	"github.com/kless/goconfig/config"
	"io"
	"log"
	"net/http"
)

func MetadataServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Metadata\n")
}

func DataServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Data\n")
}

var configFile = flag.String("c", "config.ini", "Location of the config file")

func main() {
	flag.Parse()

	c, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatal("Couldn't find the config file: ", err)
	}

	host, _ := c.String("network", "listen-ip")
	port, _ := c.Int("network", "listen-port")
	listen := fmt.Sprintf("%s:%d", host, port)

	http.HandleFunc("/metadata", MetadataServer)
	http.HandleFunc("/data", DataServer)

	log.Println("Server listening: ", listen)
	err = http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
