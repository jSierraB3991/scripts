package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Phrase struct {
	Id         int    `json:"id"`
	Text       string `json:"text"`
	Lang       string `json:"lang"`
	CategoryId int    `json:"category_id"`
	AuthorId   int    `json:"author_id"`
}

func main() {
	//https://www.positive-api.online/phrase/esp español
	//https://www.positive-api.online/phrase english
	req, err := http.NewRequest("GET", "https://www.positive-api.online/phrase", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, resp.Status)
		os.Exit(1)
	}
	decoder := json.NewDecoder(resp.Body)

	var result Phrase
	err = decoder.Decode(&result)
	if err != nil {
		fmt.Println("fortune-jd: Error al tratar de extraer la frase motivacional")
		os.Exit(1)
	}

	fmt.Println(result.Text)
}
