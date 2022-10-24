package main

import (
	"flag"
	"time"

	"github.com/things-go/go-socks5"
)

var flagInterface = flag.String("i", "", "interface to listen on")
var flagPort = flag.String("p", "10800", "port to listen on")
var flagWaitSeconds = flag.Int("w", 10, "wait seconds before listen on port")

var Version = "0.0.2"

func main() {
	println("socks5server version:", Version)
	flag.Parse()
	time.Sleep(time.Duration(*flagWaitSeconds) * time.Second)
	// Create a SOCKS5 server
	server := socks5.NewServer()

	listenAddr := *flagInterface + ":" + *flagPort
	println("socks server Listening on", listenAddr)
	// Create SOCKS5 proxy on flagInterface port flagPort
	if err := server.ListenAndServe("tcp", listenAddr); err != nil {
		panic(err)
	}
}
