package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	showAll := flag.Bool("a", false, "Mostrar archivos ocultos")
	longFormat := flag.Bool("l", false, "Usar el formato largo")
	flag.Parse()

	dir := "."
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		fmt.Printf("No se puede acceder a %s: No existe el directorio \n", dir)
		os.Exit(1)
	}

	for _, entry := range entries {
		name := entry.Name()

		if !*showAll && name[0] == '.' {
			continue
		}
		if *longFormat {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintln(os.Stderr, "error: ", err)
				continue
			}
			fmt.Printf("%s\t%8d\t%s\t%s\n",
				info.Mode().String(),
				info.Size(),
				info.ModTime().Format("Jan 02 15:04"),
				name)

		} else {
			fmt.Println(name)
		}
	}
}
