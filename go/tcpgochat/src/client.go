package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn     net.Conn
	nick     string
	room     *Room
	commands chan<- Command
}

func (this *Client) readInput() {
	for {
		msg, err := bufio.NewReader(this.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nick":
			this.commands <- Command{
				id:     CMD_NICK,
				client: this,
				args:   args,
			}
		case "/join":
			this.commands <- Command{
				id:     CMD_JOIN,
				client: this,
				args:   args,
			}
		case "/rooms":
			this.commands <- Command{
				id:     CMD_ROOMS,
				client: this,
				args:   args,
			}
		case "/msg":
			this.commands <- Command{
				id:     CMD_MSG,
				client: this,
				args:   args,
			}
		case "/quit":
			this.commands <- Command{
				id:     CMD_QUIT,
				client: this,
				args:   args,
			}
		default:
			this.err(fmt.Errorf("Unknown Command %s", cmd))
		}
	}
}

func (this *Client) err(err error) {
	this.conn.Write([]byte("ERR: " + err.Error() + "\n"))
}
func (this *Client) msg(message string) {
	this.conn.Write([]byte("> " + message + "\n"))
}
