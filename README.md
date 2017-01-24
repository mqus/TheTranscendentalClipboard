# TheTranscendentalClipboard - Server

Goal: to write a very simple server which will 
synchronize all clipboard messages of clients.
Authentification: via a small user-chosen token (yes, isn't that secure)
Transport Security: tls over tcp

planned Clients: 
- Linux
- Windows 
- Android

supported Server-OS: anything where Go can compile and run

Status:
- [x] started
- [ ] fully implement a simple server



Possible Future Features:
- [ ] save clipboard messages
    - [ ] commandline client with direct access to the server
- [ ] make it more secure (mostly auth)
- [ ] save/transmit a history of messages
- [ ] be able to transmit more than just text
