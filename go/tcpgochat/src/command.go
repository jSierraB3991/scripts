package main

type commandId int

const (
	CMD_NICK commandId = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type Command struct {
	id     commandId
	client *Client
	args   []string
}
