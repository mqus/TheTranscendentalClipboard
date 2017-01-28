package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mqus/transcendental/common"

	"time"
)

var server, room string
var c *common.PkgConn
var maxRetryTime = time.Second * 30

func main() {
	log.Println("starting up transcendental-client v0.1")
	room = "Raum1"
	server = "localhost:19192"
	if len(os.Args) > 1 {
		room = os.Args[1]
	}
	if len(os.Args) > 2 {
		server = os.Args[2]
	}

	log.Println("connect to :", server, "|", room)
	err := connectToServer()
	if err != nil {
		reconnect()
	}
	send := make(chan []byte, 2)
	recv := make(chan []byte)

	go handleSends(send)
	go handleRecv(recv)
	//TODO
	//just now there is just some test code
	if len(os.Args) > 3 {
		go inputter(send)
	}
	outputter(recv)

}

func connectToServer() error {
	addr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		log.Println("couldn't resolve adress", server, err)
		return err
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("couldn't dial adress", server, err)
		return err
	}
	c = common.NewPkgConn(conn)
	c.SendPkg(&common.Pkg{Type: "Hello", Content: []byte(room)})
	if c.IsClosed {
		log.Println("failed to send hello pkg")
		return errors.New("ClosedConnection")
	}
	return nil
}

func handleSends(send chan []byte) {
	for {
		fmt.Println(" -sendPrepare")
		data := <-send
		pkg := &common.Pkg{Type: "Text", Content: data}
		log.Print("sending:", pkg, "...")
		fmt.Println(" -sendStart")
		c.SendPkg(pkg)
		fmt.Println(" -sendEnd")
		if c.IsClosed {
			fmt.Println(" -sendClosed")
			send <- data
			log.Println("pushed back send because of closed connection")
			time.Sleep(100 * time.Millisecond)
		}
		log.Println("sent!")
	}
}

func handleRecv(recv chan<- []byte) {
	for {
		fmt.Println("-recvStart")
		pkg := c.RecvPkg()
		fmt.Println("-recvEnd")
		if c.IsClosed {
			fmt.Println("-recvReconnect")
			reconnect()
			fmt.Println("-recvReconnectDone")
		} else {
			fmt.Println("-recvPkg")
			if pkg.Type == "Text" {
				recv <- pkg.Content
			}
		}
	}
}

func reconnect() {
	err := errors.New("empty")
	for i := time.Second; err != nil; {
		log.Println("Wait for", i, "and try again")
		time.Sleep(i)
		err = connectToServer()
		i *= 2
		if i > maxRetryTime {
			i = maxRetryTime
		}
	}
	log.Println("Connected!")

}

func outputter(recv <-chan []byte) {
	for {
		fmt.Println("---received: ", time.Now(), string(<-recv))
	}
}

func inputter(send chan<- []byte) {
	switch os.Args[3] {
	case "1":
		send <- []byte("passwort")
		time.Sleep(5 * time.Second)
		send <- []byte("pass2wort")
		send <- []byte("passasdwort")
		time.Sleep(13 * time.Second)
		send <- []byte("send<-[]byte(\"passwort\")")
	case "2":
		send <- []byte("daswarwohl nix")
		time.Sleep(7 * time.Second)
		send <- []byte("versuch2")
		c.Close()
		send <- []byte("versuchsmalgeschlossen")
		time.Sleep(10 * time.Second)
		send <- []byte("wartmakurz")

	}

}
