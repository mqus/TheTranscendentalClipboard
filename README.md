# transcendental

Goal: to write a very simple server which will 
synchronize all clipboard messages of clients.
Authentification: via a small user-chosen token (yes, isn't that secure)
Transport Security: Future:tls over tcp Currently: base64 ;)

planned Clients: 
- Linux
- Windows 
- Android

supported Server-OS: anything where Go can compile and run

Status:
- [x] started
- [x] fully implement a simple server
    - [ ] make it accessible by providing commandline options
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
    