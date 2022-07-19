package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the Ip",
	Long:  "Trace the Ip",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {

			for _, ip := range args {
				fmt.Println(green(fmt.Sprintf("------SHOW INFO FOR IP %s-------", ip)))
				showData(ip)
			}
		} else {
			fmt.Println("Please Provide IP to trace")
		}
	},
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unabl to Unmarshal the response")
	}

	fmt.Println(blue(fmt.Sprintf("Ip: %s", ip)))
	fmt.Printf("Host Name: %s\n", data.Hostname)
	fmt.Printf("Any Cast: %v\n", data.AnyCast)
	fmt.Printf("City: %s\n", data.City)
	fmt.Printf("Region: %s\n", data.Region)
	fmt.Printf("Country: %s\n", data.Country)
	fmt.Printf("Location: %s\n", data.Location)
	fmt.Printf("Organization: %s\n", data.Organization)
	fmt.Printf("Postal Code: %s\n", data.PostalCode)
	fmt.Printf("Time Zone: %s\n", data.TimeZone)
	fmt.Printf("\n")
}

const (
	ColorDefault = "\x1b[39m"
	ColorBlue    = "\x1b[94m"
	ColorGreen   = "\x1b[32m"
)

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

type Ip struct {
	Ip           string `json:"ip"`
	City         string `json:"city"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	Location     string `json:"loc"`
	PostalCode   string `json:"postal"`
	TimeZone     string `json:"timezone"`
	Hostname     string `json:"hostname"`
	AnyCast      bool   `json:"anycast"`
	Organization string `json:"org"`
	Readme       string `json:"readme"`
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Unabl to get the response")
		os.Exit(1)
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Unabl to get the response")
		os.Exit(1)
	}
	return responseByte
}

func init() {
	rootCmd.AddCommand(traceCmd)
}
