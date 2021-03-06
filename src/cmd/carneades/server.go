package main

import (
	"flag"
	"github.com/carneades/carneades-4/src/web"
	"log"
	"os"
	"path/filepath"
)

const helpServer = `
usage: carneades server [-p port-number] [-t template-directory]

Starts the Carneades web service.

The -p flag ("port") specifies the port number to use (default: 8080)

The -t flag ("templates") specifies the full path to the directory with
the HTML templates used by the service 
(default: $GOPATH/src/github.com/carneades/carneades-4/src/web/templates/)
`

func getGoPath() string {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		return os.Getenv("HOME") + "/go"
	} else {
		return goPath
	}
}

const defaultPort = "8080"

var defaultTemplatesDir = filepath.Join(getGoPath(), "/src/github.com/carneades/carneades-4/src/web/templates/")

func webServerCmd() {
	webFlags := flag.NewFlagSet("webFlags", flag.ContinueOnError)
	port := webFlags.String("p", defaultPort, "the port number of the HTTP Server")
	templatesDir := webFlags.String("t", defaultTemplatesDir, "the path to the templates directory")
	if err := webFlags.Parse(os.Args[2:]); err != nil {
		log.Fatal(err)
	}

	web.CarneadesServer(*port, *templatesDir)
}
