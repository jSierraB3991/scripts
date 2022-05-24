package main 

import (
    "github.com/gen2brain/beeep"
)

func main() {
    err := beeep.Notify("Title", "Message body", "anime/information.png")
    if err != nil {
        panic(err)
    }
    err = beeep.Alert("Title", "Message body", "anime/warning.png")
    if err != nil {
        panic(err)
    }
}
