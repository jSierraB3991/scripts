package main 

import (
    "os"
    "os/exec"
    "github.com/gen2brain/beeep"
)

func main() {
    err := beeep.Notify("Title", "Message body", "dialog-information")
    if err != nil {
        panic(err)
    }
    err = beeep.Alert("Title", "Message body", "dialog-error")
    if err != nil {
        panic(err)
    }

    send, err := exec.LookPath("notify-send")
    if err != nil {
        os.Exit(1)
    }
    c := exec.Command(send, "Title", "Message body", "-i", "dialog-information")
    c.Run()
    c = exec.Command(send, "Title", "Message body", "-i", "dialog-error")
    c.Run()
}
