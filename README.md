# CLI chat app on top of TCP

## Usage

[Setup Go workspace](https://go.dev/doc/gopath_code) environment. Navigate to your Go workspace and clone this repo

```git
git clone https://github.com/ngakana/go-tcp-socket-cli-chat-app
```

```bash
cd ./go-tcp-socket-cli-chat-app
```

Start the server up by running command below

```go
go run *.go
```

This will setup a TCP socket on the server with IP=localhost and port=8080. Clients can communicate to one another in a 'chatroom' through this socket using telnet.

```bash
telnet localhost 8080
```

## Commands

<span style="color: orange">/alias</span> <span style="color: green"><i>name</i></span> [String] - change alias name
<br>
<span style="color: orange">/join</span> <span style="color: green"><i>room</i></span> [String] - join/create a chat room. Users can join one room at a time
<br>
<span style="color: orange">/send</span> <span style="color: green"><i>msg</i></span> [String]
<br>
<span style="color: orange">/leave</span> - leave current room. Pass if user has not joined a room
<br>
<span style="color: orange">/quit</span> - close client tcp socket connection
