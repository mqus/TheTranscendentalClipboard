# transcendental

Goal: to write a very simple server which will 
synchronize all clipboard messages of clients.

Authentification: via a small user-chosen token (yes, isn't that secure)

Transport Security: Future:~~tls over tcp~~ ? Currently: the java-stl-implementation of AES/ECB/PKCS5Padding

planned Clients: 
- Linux
- Windows 
- Android
- (macOS, iOS) - not possible (yet) because i don't have any apple devices at hand

supported Server-OS: anything where Go can compile and run

Status:
- [x] started
- [x] fully implement a simple server
    - [x] make it accessible by providing commandline options
    - [ ] a bit of documentation (how to start, build, use)
- [x] a simple example client implemented (send/rcv)
    - [x] read from clipboard + detect changes 
    - [x] write to clipboard
- [ ] specify the client
- [x] implement a Simple Java client which supports lazy-transmit and more data flavours
- [ ] write tests

Possible Future Features:
- [x] be able to transmit more than just text
    - [x] Look up specifications on Windows, Android and X11 Clipboard Mechanisms with different MIME-types
- [_] make it more secure (mostly auth)
    - not yet really secure (no tests, no audit)
- [ ] save clipboard messages
    - [ ] commandline client with direct access to the server
- [ ] save/transmit a history of messages

Links:
- the official java client:
    https://github.com/mqus/transcendental-client


## build:
Dependencies:
 - go

To build the server, run
```sh
$ go get github.com/mqus/transcendental/transcendental-server
$ go install github.com/mqus/transcendental/transcendental-server
```

this will compile a binary and will place it in $GOPATH/bin
(typically $HOME/go/bin)

## run:
simply running it will launch a server, which
listens on incoming tcp-connections from all interfaces on port 19192.

```sh
$ $GOPATH/bin/transcendental-server
2018/02/14 21:24:19 starting up transcendental-server v0.2
2018/02/14 21:24:19 listening on [::]:19192
```

this can be changed by simply providing the listening address
as the first argument to the program, e.g.:
```sh
$GOPATH/bin/transcendental-server :1790
```
to listen on port 1790, or
```sh
$GOPATH/bin/transcendental-server localhost:0
```
to accept connections only from localhost and
get a random port assigned.

The running programm won't create or need any files.