package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "."
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		name := entry.Name()
		fmt.Println(name)
	}
}
