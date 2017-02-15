package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

// Hello world bytes
var hwb = []byte("Hello world!")

func main() {
	var (
		// HTTP handler
		s srv
		// Listening address
		addr string
		// Listening port
		port string
	)

	// Assign addr variable to the addr flag
	flag.StringVar(&addr, "addr", "localhost", "listening host (eg --addr=localhost)")
	// Assign port variable to the port flag
	flag.StringVar(&port, "port", "80", "port to listen to (eg --port=80)")
	// Parse our flags before moving forward
	flag.Parse()

	// Set our location as a concatination of addr and port, with ":" as a port separator
	loc := addr + ":" + port

	// Post message to stdout notifying of listening location
	fmt.Println("Listening at " + loc)

	// Listen and serve using our srv struct as a handler
	if err := http.ListenAndServe(loc, &s); err != nil {
		// Post message to stderr notifying of error encountered while listening
		fmt.Fprintln(os.Stderr, "Error encountered!:", err)
	}
}

type srv struct{}

func (s *srv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Write our hello world bytes to the http.ResponseWriter
	w.Write(hwb)
}
