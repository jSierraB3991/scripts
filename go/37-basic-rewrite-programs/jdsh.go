package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		cwd, _ := os.Getwd()
		fmt.Printf("%s $ ", cwd)
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fields := strings.Fields(line)
		cmd := fields[0]

		switch cmd {
		case "exit":
			os.Exit(0)
		}

	}
}
