package main

import (
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func GetQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
	svc := sqs.New(sess)
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SendMsg(sess *session.Session, queueURL *string) error {
	svc := sqs.New(sess)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
			"Author": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("John Grisham"),
			},
			"WeeksOn": &sqs.MessageAttributeValue{
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
		MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 12/11/2016."),
		QueueUrl:    queueURL,
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	queue := flag.String("q", "", "The name of the queue")
	flag.Parse()

	if *queue == "" {
		log.Println("You must supply the name of a queue (-q QUEUE)")
		return
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	result, err := GetQueueURL(sess, queue)
	if err != nil {
		log.Println("Got an error getting the queue URL:")
		log.Println(err)
		return
	}

	queueURL := result.QueueUrl

	err = SendMsg(sess, queueURL)
	if err != nil {
		log.Println("Got an error sending the message:")
		log.Println(err)
		return
	}

	log.Println("Sent message to queue ")
}
