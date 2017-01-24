package srv

import (
	"errors"
	"net"
)

import "encoding/json"

// A Room knows all the clients which share the same Clipboard.
type Room struct {
	name    string
	clients []Client
}

// A Client knows his connection(including all the neccessary en/decoders) and his position in the parent room.
// the position is important to not send the same clipboard message back (and inducing some ugly race conditions)
type Client struct {
	conn net.Conn
	dec  *json.Decoder
	enc  *json.Encoder
	//TODO (add number(dangerous, thinkitover) and space)
}

type pkg struct {
	msgtype string
	content []byte
}

var (
	// ErrConnClosed is thrown when the package size cannot be read/written or the package cannot be read/written entirely.
	ErrConnClosed = errors.New("client lost/closed the connection")
)

//MAYBE: change to a binary protocol by sending [type(1Byte)[,size(8Byte),msg(sizeByte)]] where type is one of {<close>, <join>, <msg>, <ping>}
func (c *Client) sendpkg(p pkg) (err error) {
	err = c.enc.Encode(p)
	return
}

func (c *Client) recvpkg() (p pkg, err error) {
	err = c.dec.Decode(&p)
	return
}

// AddClient adds a new Client to the system by encoding the first message with the room name and then
// assigning the Client to the room.
func AddClient(conn net.Conn) {
	c := Client{
		conn: conn,
		dec:  json.NewDecoder(conn),
		enc:  json.NewEncoder(conn),
	}
	//TODO
	c.recvpkg()
}
