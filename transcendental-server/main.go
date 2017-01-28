package main

import "fmt"
import "net"
import "log"
import "time"

//import "github.com/mqus/transcendental/srv"

func main() {
	log.Println("starting up transcendental-server v0.1")
	fmt.Println(log.Prefix(), "Hi!")
	//ln, err := net.Listen("tcp", ":19192")
	addr, err := net.ResolveTCPAddr("tcp", ":19192")
	if err != nil {
		log.Fatal(err)
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal("Couldn't open port 19192! Details:", err)
	}
	defer ln.Close()
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Println("couldn't accept connection! Details:", err)
		}
		go handleNewConnection(conn)
	}

}
func handleNewConnection(conn *net.TCPConn) {
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(30 * time.Second)
	conn.SetNoDelay(true)
	AddClient(conn)
}
