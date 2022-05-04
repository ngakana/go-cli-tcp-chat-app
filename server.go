package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	rooms map[string]*Room
	cmds chan Command
}

func newServer() *Server {
	return &Server {
		rooms: make(map[string]*Room),
		cmds: make(chan Command),
	}
}

func (s *Server) newClient(conn net.Conn) {
	log.Printf("A client has connected: %s", conn.RemoteAddr().String())

	c := &Client {
		conn: conn,
		alias: "anonym",
		cmds: s.cmds,
	}

	c.readInput()
}

func (s *Server) run() {
	for cmd := range s.cmds {
		switch cmd.id {
		case CMD_ALIAS:
			s.alias(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_LEAVE:
			s.leave(cmd.client)
		case CMD_SEND:
			s.send(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *Server) alias(c *Client, args []string) {
	curr := c.alias
	c.alias = args[1]
	c.msg(fmt.Sprintf("alias changed from %s to %s", curr, c.alias))
}

func (s *Server) join(c *Client, args []string) {
	rName := args[1]
	r, ok := s.rooms[rName]

	if ok {
		c.msg(fmt.Sprintf("room '%s' already exists. %d participants waiting...", rName, len(r.members)))
	} else {
		r = &Room{
			name: rName,
			members: make(map[net.Addr]*Client),
		}
		s.rooms[rName] = r
	}
	s.leave(c)
	c.room = r
	r.members[c.conn.RemoteAddr()] = c

	r.broadcast(c, fmt.Sprintf("%s has joined the chat", c.alias))
	c.msg(fmt.Sprintf("You can now see and write chat messages in room '%s'", r.name))
}

func (s *Server) send(c *Client, args []string) {
	if c.room == nil {
		c.err(errors.New("you must join a room to send messages"))
		return
	}
	msg := strings.Join(args[1:]," ")
	c.room.broadcast(c, fmt.Sprintf("%s: %s", c.alias, msg))
}

func (s *Server) leave(c *Client) {
	if c.room != nil {
		c.room.broadcast(c, fmt.Sprintf("%s has left the chat", c.alias))
		delete(c.room.members, c.conn.RemoteAddr())
	}
}

func (s *Server) quit(c *Client) {
	log.Printf("A client has disconnected: %s", c.conn.RemoteAddr().String())
	s.leave(c)
	c.msg("Closing connection...")
	c.msg("Connection successfully closed.")
	c.conn.Close()
}