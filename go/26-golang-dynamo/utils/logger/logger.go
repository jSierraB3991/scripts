package logger

import "log"

func Panic(message string, err error) {
	if err != nil {
		log.Panic(message, err)
	}
}

func Info(message string, data interface{}) {
	log.Println(message, data)
}
