package main

import (
	"net"
	"time"
	"os"
	"log"
	"strings"
	"fmt"
)

func main() {

	addressString := parseArgs(os.Args[1:])

	log.Println("starting up transcendental-server v0.3")

	addr, err := net.ResolveTCPAddr("tcp", addressString)
	if err != nil {
		log.Fatal("Couldn't resolve address! Details:", err)
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal("Couldn't open port ", addr.Port, "! Details:", err)
	}
	log.Println("listening on", ln.Addr())
	defer ln.Close()
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Println("couldn't accept connection! Details:", err)
		}
		go handleNewConnection(conn)
	}

}

func parseArgs(args []string) string {
	addressString := ":19192"
	for i := range args {
		if strings.HasPrefix(args[i], "--") {
			switch args[i] {
			case "-h":
				fallthrough
			case "--help":
				fmt.Println("transcendental-server v0.3, a server application for synchronizing clipboards.\n" +
					"Usage: transcendental-server [server:port] [flags]\n" +
					"\t[server:port]\t The adress on which the server will listen on incoming connections (default: ':19192'" +
					"\nFlags:\n" +
					"\t--help\t-h\t print this dialog.")
				os.Exit(0)
			}
			continue
		} else {
			addressString = args[i]
		}

	}
	return addressString
}
func handleNewConnection(conn *net.TCPConn) {
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(30 * time.Second)
	conn.SetNoDelay(true)
	AddClient(conn)
}
