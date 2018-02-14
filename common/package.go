package common

import (
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"
	"io"
)

//Pkg is the unit in which messages are transferred
//between the clients and the server.
type Pkg struct {
	ClientID int
	Type     string
	Content  []byte
}

type PkgConn struct {
	conn *net.TCPConn
	dec  *json.Decoder
	enc  *json.Encoder
	//closefn  func(*PkgConn)
	CanClose *sync.Cond
	IsClosed bool
}

func NewPkgConn(conn *net.TCPConn) *PkgConn {
	pc := &PkgConn{
		conn:     conn,
		dec:      json.NewDecoder(conn),
		enc:      json.NewEncoder(conn),
		IsClosed: false,
		CanClose: sync.NewCond(&sync.Mutex{}),
		//closefn:  closer,
	}

	return pc
}

//MAYBE: change to a binary protocol by sending
//[type(1Byte),clientID(4Byte),size(8Byte),msg(sizeByte)] where

func (c *PkgConn) SendPkg(p *Pkg) {
	if c.IsClosed {
		log.Println("conn is closed before send")
		return
	}

	c.conn.SetWriteDeadline(time.Now().Add(time.Minute))

	err := c.enc.Encode(p)
	if err != nil {
		if err != io.EOF {
			log.Println("conn-s is weird", err, "|", p)
		}
		c.Close()
	}

}

func (c *PkgConn) RecvPkg() (p *Pkg) {
	p = new(Pkg)
	if c.IsClosed {
		log.Println("conn is closed before rcv")
		return nil
	}
	err := c.dec.Decode(p)
	if err != nil {
		if err != io.EOF {
			log.Println("conn-r is weird", err, "|", p)
		}
		c.Close()
	}
	return
}

func (c *PkgConn) Close() {
	if !c.IsClosed {
		c.IsClosed = true
		//c.closefn(c)
		c.conn.Close()
		c.CanClose.Broadcast()
	}
}

func (c *PkgConn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
