package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/eatonphil/gosql"
)

func printHelp() {
	fmt.Println(`
=========================================
             GOPSQL HELP
=========================================

Comandos
---------
help              Muestra esta ayuda.
exit              Sale del programa.
quit              Sale del programa.
\q                Sale del programa.

SQL soportado
-------------
CREATE TABLE      Crea una tabla.
INSERT INTO       Inserta una fila.
SELECT *          Consulta todas las columnas.
SELECT col1,...   Consulta columnas específicas.
WHERE             Filtra resultados.

Tipos soportados
----------------
INT
TEXT

Ejemplos de sentencias
----------------------
CREATE TABLE users (...);
INSERT INTO users VALUES (...);
SELECT * FROM users;
SELECT id, name FROM users;
SELECT * FROM users WHERE id = 1;
`)
}
func main() {
	mb := gosql.NewMemoryBackend()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to gopsql")
	for {
		fmt.Print("# ")
		text, err := reader.ReadString('\n')
		textHelp := strings.TrimSpace(text)
		switch strings.ToLower(textHelp) {
		case "exit", "quit":
			fmt.Println("Bye")
			os.Exit(0)
		case "help":
			printHelp()
			continue
		}
		text = strings.Replace(text, "\n", "", -1)

		parser := gosql.Parser{}

		ast, err := parser.Parse(text)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		for _, stmt := range ast.Statements {
			switch stmt.Kind {
			case gosql.CreateTableKind:
				err = mb.CreateTable(ast.Statements[0].CreateTableStatement)
				if err != nil {
					fmt.Println(err)
					continue
				}
			case gosql.InsertKind:
				err = mb.Insert(stmt.InsertStatement)
				if err != nil {
					fmt.Println(err)
					continue
				}
			case gosql.SelectKind:
				results, err := mb.Select(stmt.SelectStatement)
				if err != nil {
					log.Println(err)
					continue
				}
				for _, col := range results.Columns {
					fmt.Printf("| %s ", col.Name)
				}
				fmt.Println("|")
				for i := 0; i < 20; i++ {
					fmt.Printf("=")
				}
				fmt.Println()

				for _, result := range results.Rows {
					fmt.Printf("|")

					for i, cell := range result {
						typ := results.Columns[i].Type
						s := ""
						switch typ {
						case gosql.IntType:
							s = fmt.Sprintf("%d", cell.AsInt())
						case gosql.TextType:
							s = *cell.AsText()
						}

						fmt.Printf(" %s | ", s)
					}

					fmt.Println()
				}

			}
		}
		fmt.Println("ok")
	}
}
