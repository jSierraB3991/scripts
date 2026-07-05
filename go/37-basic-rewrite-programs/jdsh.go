package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func printHelp() {
	fmt.Println(`Comandos disponibles:
	cd [dir]    - Cambiar el directorio (sin argumento va al home)
	exit	    - Salir de la shell
	help	    - Mostrar la ayuda
	`)
}

func handleCd(args []string) {
	dir := ""
	if len(args) == 0 {
		home, _ := os.UserHomeDir()
		dir = home
	} else {
		dir = args[0]
	}

	if err := os.Chdir(dir); err != nil {
		fmt.Printf("cd: No existe el directorio %s\n", dir)
		fmt.Fprintln(os.Stderr, err)
	}
}

func runExternal(cmd string, args []string) {
	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		fmt.Printf("jdsh: %s instrucción no encontrada...\n", cmd)
		fmt.Fprintln(os.Stderr, err)
	}
}

func main() {
	colorCwd := "\033[33m"
	resetColor := "\033[0m"
	reader := bufio.NewReader(os.Stdin)

	for {
		cwd, _ := os.Getwd()
		fmt.Printf("%s %s %s > ", colorCwd, cwd, resetColor)
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fields := strings.Fields(line)
		if len(fields) <= 0 {
			continue
		}

		cmd := fields[0]
		args := fields[1:]

		switch cmd {
		case "exit":
			os.Exit(0)
		case "cd":
			handleCd(args)
		case "help":
			printHelp()
		default:
			runExternal(cmd, args)
		}
		fmt.Println()
	}
}
