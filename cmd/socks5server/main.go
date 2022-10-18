package main

import (
	"flag"
	"github.com/things-go/go-socks5"
)

var flagInterface = flag.String("i", "", "interface to listen on")
var flagPort = flag.String("p", "10800", "port to listen on")

func main() {
	flag.Parse()
	// Create a SOCKS5 server
	server := socks5.NewServer()

	listenAddr := *flagInterface + ":" + *flagPort
	println("socks server Listening on", listenAddr)
	// Create SOCKS5 proxy on flagInterface port flagPort
	if err := server.ListenAndServe("tcp", listenAddr); err != nil {
		panic(err)
	}
}
