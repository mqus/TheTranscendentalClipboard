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
- [x] a simple example client implemented (send/rcv)
    - [ ] read from clipboard + detect changes 
    - [ ] write to clipboard
- [ ] specify the client
- [ ] implement a Simple Java client which supports lazy-transmit and more data flavours
- [ ] write tests

Possible Future Features:
- [ ] be able to transmit more than just text
    - [ ] Look up specifications on Windows, Android and X11 Clipboard Mechanisms with different MIME-types
- [ ] make it more secure (mostly auth)
- [ ] save clipboard messages
    - [ ] commandline client with direct access to the server
- [ ] save/transmit a history of messages

Links:
- possible library (for simple text passing):
    https://github.com/atotto/clipboard
- the official java client:
    https://github.com/mqus/transcendental-client
    