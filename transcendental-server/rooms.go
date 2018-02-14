package main

import (
	"errors"
	"log"
	"net"
	"sync"

	"github.com/mqus/transcendental/common"
)

// A Room knows all the clients which share the same Clipboard.
type Room struct {
	name    string
	clients map[int]*Client
	maxid   int
	mutex   sync.RWMutex
}

// A Client knows his connection(including all the neccessary en/decoders) and his position in the parent room.
// the position is important to not send the same clipboard message back (and inducing some ugly race conditions)
type Client struct {
	conn *common.PkgConn
	id   int
	room *Room
}

var (
	// ErrConnClosed is thrown when the package size cannot be read/written or the package cannot be read/written entirely.
	ErrConnClosed = errors.New("client lost/closed the connection")
	rooms         = make(map[string]*Room)
	roomsMutex    sync.RWMutex
)

// AddClient adds a new Client to the system by encoding the first message with the room name and then
// assigning the Client to the room.
func AddClient(conn *net.TCPConn) {
	tc := common.NewPkgConn(conn)
	//conn.SetReadDeadline(time.Now().Add(time.Minute))
	pkg := tc.RecvPkg()
	if tc.IsClosed {
		log.Println("Client was closed before the room was assigned... suspicious...")
		return
	}
	//	log.Println("type", pkg.Type, "|content", string(pkg.Content))
	if pkg.Type != "Hello" {
		tc.Close()
		log.Println("Client sent the wrong first message:", pkg)
		return
	}
	roomname := string(pkg.Content)
	room := getRoom(roomname)
	c := Client{
		conn: tc,
	}

	//add Client to Room and the room to the client
	room.mutex.Lock()
	room.maxid++
	c.id = room.maxid

	room.clients[c.id] = &c
	c.room = room
	log.Printf("Added a new Client(cid:%d,%s)\tto room %s,\tsize:%d\n",
		c.id, c.Addr(), string(pkg.Content), len(room.clients))
	//log.Println("#clients in room, cid:", len(room.clients), c.id)

	room.mutex.Unlock()
	go waitForClosing(&c)

	c.recvLoop()
}

func getRoom(name string) (room *Room) {
	roomsMutex.RLock()
	room = rooms[name]
	roomsMutex.RUnlock()

	//if room is not there already, add it
	if room == nil {
		room = &Room{
			name:    name,
			clients: make(map[int]*Client),
			maxid:   0,
		}

		roomsMutex.Lock()
		rooms[name] = room
		roomsMutex.Unlock()

	}
	return room
}

func (c *Client) recvLoop() {
	for {
		pkg := c.conn.RecvPkg()
		if c.conn.IsClosed {
			return
		}
		log.Println("got pkg", c.id, pkg)
		switch pkg.Type {
		case "Text":
			fallthrough
		case "Copy":
			//safely read all currently connected clients
			c.room.mutex.RLock()
			clients := c.room.clients
			c.room.mutex.RUnlock()

			//Write FromID to the pkg
			pkg.ClientID = c.id

			//relay package to all other clients
			for toID, client := range clients {
				if toID != c.id {
					client.conn.SendPkg(pkg)
					log.Println("sentsomthg")
				}
			}

		case "Request":
			fallthrough
		case "Data":
			fallthrough
		case "Reject":
			//safely get client to send package to
			c.room.mutex.RLock()
			to, ok := c.room.clients[pkg.ClientID]
			c.room.mutex.RUnlock()

			if !ok {
				//if the requested Client is not there anymore, respond
				//with the same data and ClientID set to zero
				pkg.ClientID = 0
				to = c
			} else {
				//Write FromID to the pkg
				pkg.ClientID = c.id
			}

			to.conn.SendPkg(pkg)
		}
	}
}

func (c *Client) Addr() string {
	return c.conn.RemoteAddr().String()
}

func waitForClosing(c *Client) {
	//Wait till PkgConn has closed the connection and then close from this end.
	c.conn.CanClose.L.Lock()
	c.conn.CanClose.Wait()
	if c.room != nil {
		//delete client from room
		c.room.mutex.Lock()
		delete(c.room.clients, c.id)
		c.room.mutex.Unlock()
		log.Println("closed connection of cid", c.id, "(", c.Addr(), ")")
		//MAYBE: implement some kind of GC
		/* IS NOT SAFE TO DO, THEREFORE COMMENTED OUT
		c.room.mutex.RLock()

		//if the room is now empty, delete the room.
		if len(c.room.clients) == 0 and c.room.{
			roomsMutex.Lock()
			delete(rooms, c.room.name)
			////////////////22222
			roomsMutex.Lock()
		}
		c.room.mutex.RUnlock()*/
	}
	c.conn.CanClose.L.Unlock()
}
