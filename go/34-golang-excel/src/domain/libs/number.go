package libs

import (
	"log"
	"strconv"
)

func GetUintForString(number string) uint {
	if number == "" {
		return 0
	}
	response, err := strconv.Atoi(number)
	if err != nil {
		log.Panic(err)
	}
	return uint(response)
}
