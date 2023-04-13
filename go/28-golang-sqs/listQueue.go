package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)

	result, err := svc.ListQueues(nil)
	if err != nil {
		log.Println(err.Error())
	}
	for i, url := range result.QueueUrls {
		log.Printf("%d: %s\n", i, *url)
	}
}
