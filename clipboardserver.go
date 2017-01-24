package main

import "fmt"
import "net"
import "log"
import this "github.com/mqus/TheTranscendentalClipboard-Server/srv"

func main() {
	fmt.Println(log.Prefix(), "Hi!")
	ln, err := net.Listen("tcp", ":19192")
	if err != nil {
		log.Fatal("Couldn't open port 19192! Details:", err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("couldn't accept connection! Details:", err)
		}
		go handleNewConnection(conn)
	}

}
func handleNewConnection(conn net.Conn) {
	this.AddClient(conn)
	defer conn.Close()
}
