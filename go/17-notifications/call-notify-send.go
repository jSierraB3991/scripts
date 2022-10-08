package main

import (
	"os"
	"os/exec"
)

func main() {
	send, err := exec.LookPath("notify-send")
	if err != nil {
		os.Exit(1)
	}
	c := exec.Command(send, "Title", "Message body", "-i", "dialog-information")
	c.Run()
	c = exec.Command(send, "Title", "Message body", "-i", "dialog-error", "-u", "CRITICAL")
	c.Run()
}
