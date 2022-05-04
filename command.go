package main

type cmdID int

const (
	CMD_ALIAS cmdID = iota
	CMD_JOIN
	CMD_SEND
	CMD_LEAVE
	CMD_QUIT
)

type Command struct {
	id cmdID
	client *Client
	args []string
}