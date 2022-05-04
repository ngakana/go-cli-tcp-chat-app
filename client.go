package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
	alias string
	room *Room
	cmds chan<- Command
}

func (c *Client) readInput() {
	for {
		in, err := bufio.NewReader(c.conn).ReadString('\n')

		if err != nil {
			return
		}

		in =	strings.Trim(in, "\r\n")
		args := strings.Split(in," ")
		cmd := strings.TrimSpace(args[0])
		cmd = strings.ToLower(cmd)

		switch cmd {
		case "/alias":
			c.cmds<- Command{
				id: CMD_ALIAS,
				client: c,
				args: args,
			}
		case "/send":
			c.cmds<- Command{
				id: CMD_SEND,
				client: c,
				args: args,
			}
		case "/join":
			c.cmds<- Command{
				id: CMD_JOIN,
				client: c,
				args: args,
			}
		case "/leave":
			c.cmds<- Command{
				id: CMD_LEAVE,
				client: c,
				args: args,
			}
		case "/quit":
			c.cmds<- Command{
				id: CMD_QUIT,
				client: c,
				args: args,
			}
		default:
			c.err(fmt.Errorf("Unknown command: %s", cmd))
		}
	}
}

func (c *Client) err(err error) {
	c.conn.Write([]byte("ERR: " + err.Error() + "\n"))
}

func (c *Client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}