package main

import "net"

type Room struct {
	name string
	members map[net.Addr]*Client
}

func (r *Room) broadcast(c *Client, msg string) {
	if _, ok := r.members[c.conn.RemoteAddr()]; ok {
		for addr, m := range r.members {
			if addr != c.conn.RemoteAddr() {
				m.msg(msg)
			}
		}
	}
}