package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jsierrab3991/scripts/basic-gosql/gosql"
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
func printTableMetadata(tables []gosql.TableMetadata) {
	for _, table := range tables {
		fmt.Printf("\n Table: %s\n", table.Name)

		// Calcular ancho de cada columna
		nameWidth := len("Name")
		typeWidth := len("Type")
		notNullWidth := len("Not Null")

		for _, column := range table.Columns {
			if len(column.Name) > nameWidth {
				nameWidth = len(column.Name)
			}

			typeName := fmt.Sprintf("%v", column.Type)
			if len(typeName) > typeWidth {
				typeWidth = len(typeName)
			}

			notNull := fmt.Sprintf("%t", column.NotNull)
			if len(notNull) > notNullWidth {
				notNullWidth = len(notNull)
			}
		}

		// Separador
		fmt.Print(" ")
		fmt.Print(strings.Repeat("-", nameWidth+2))
		fmt.Print("+")
		fmt.Print(strings.Repeat("-", typeWidth+2))
		fmt.Print("+")
		fmt.Print(strings.Repeat("-", notNullWidth+2))
		fmt.Println()

		// Header
		fmt.Printf(
			" %-*s | %-*s | %-*s\n",
			nameWidth, "Name",
			typeWidth, "Type",
			notNullWidth, "Not Null",
		)

		// Separador
		fmt.Print(" ")
		fmt.Print(strings.Repeat("-", nameWidth+2))
		fmt.Print("+")
		fmt.Print(strings.Repeat("-", typeWidth+2))
		fmt.Print("+")
		fmt.Println(strings.Repeat("-", notNullWidth+2))

		// Filas
		for _, column := range table.Columns {
			fmt.Printf(
				" %-*s | %-*v | %-*t\n",
				nameWidth, column.Name,
				typeWidth, column.Type,
				notNullWidth, column.NotNull,
			)
		}
	}
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
		case "exit", "quit", "\\q":
			fmt.Println("Bye")
			os.Exit(0)
		case "help":
			printHelp()
			continue
		case "\\l":
			tables := mb.GetTables()
			if len(tables) <= 0 {
				fmt.Println("No tables found.")
				continue
			}
			printTableMetadata(tables)
			fmt.Println("ok")
		}
		text = strings.ReplaceAll(text, "\n", "")

		parser := gosql.Parser{}

		ast, err := parser.Parse(text)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, stmt := range ast.Statements {
			switch stmt.Kind {
			case gosql.CreateTableKind:
				err = mb.CreateTable(ast.Statements[0].CreateTableStatement)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println("ok")
			case gosql.InsertKind:
				err = mb.Insert(stmt.InsertStatement)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println("ok")
			case gosql.SelectKind:
				results, err := mb.Select(stmt.SelectStatement)
				if err != nil {
					fmt.Println(err)
					continue
				}

				if len(results.Rows) == 0 {
					fmt.Println("No rows found.")
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
			default:
				fmt.Printf("Unsupported statement: %v\n", stmt.Kind)
			}
		}
	}
}
