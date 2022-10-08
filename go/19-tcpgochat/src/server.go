package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	rooms    map[string]*Room
	commands chan Command
}

func NewServer() *Server {
	return &Server{
		rooms:    make(map[string]*Room),
		commands: make(chan Command),
	}
}

func (this *Server) Run() {
	for cmd := range this.commands {
		switch cmd.id {
		case CMD_NICK:
			this.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			this.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			this.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			this.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			this.quit(cmd.client, cmd.args)
		}
	}
}

func (this *Server) newClient(conn net.Conn) {
	log.Printf("New cliente has connect: %s", conn.RemoteAddr().String())
	c := &Client{
		conn:     conn,
		nick:     "anonymus",
		commands: this.commands,
	}
	c.readInput()
}

func (this *Server) nick(c *Client, args []string) {
	if len(args) < 2 {
		c.msg("nick name is required. usage: /nick NAME")
		return
	}
	c.nick = args[1]
	c.msg(fmt.Sprintf("All rigth, I Will call you %s", c.nick))
}

func (this *Server) join(c *Client, args []string) {
	if len(args) < 2 {
		c.msg("room name is required. usage: /join ROOM_NAME")
		return
	}
	roomName := args[1]
	r, ok := this.rooms[roomName]
	if !ok {
		r = &Room{
			name:    roomName,
			members: make(map[net.Addr]*Client),
		}
		this.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c

	this.quitCurrenRoom(c)
	c.room = r
	r.Broadcast(c, fmt.Sprintf("%s has joined the room", c.nick))
	c.msg(fmt.Sprintf("Welcome to %s", r.name))
}

func (this *Server) quitCurrenRoom(c *Client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.Broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
	}
}

func (this *Server) listRooms(c *Client, args []string) {
	var rooms []string
	for name := range this.rooms {
		rooms = append(rooms, name)
	}
	c.msg(fmt.Sprintf("Available rooms are: %s", strings.Join(rooms, ", ")))
}

func (this *Server) msg(c *Client, args []string) {
	if len(args) < 2 {
		c.msg("message is required. usage: /msg MESSAGE")
		return
	}
	if c.room == nil {
		c.err(errors.New("You must Join the room first"))
		return
	}

	c.room.Broadcast(c, c.nick+": "+strings.Join(args[1:len(args)], " "))
}

func (this *Server) quit(c *Client, args []string) {
	log.Printf("Client has disconnect %s", c.conn.RemoteAddr().String())
	this.quitCurrenRoom(c)
	c.msg("Sad to see you go :(")
	c.conn.Close()
}
