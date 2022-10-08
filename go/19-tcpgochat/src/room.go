package main

import "net"

type Room struct {
	name    string
	members map[net.Addr]*Client
}

func (this *Room) Broadcast(sender *Client, msg string) {
	for addr, m := range this.members {
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}
