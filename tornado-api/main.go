package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

type Notice struct {
	Origin  string `json:"origin"`
	Message string `json:"message"`
	Channel string `json:"channel"`
}

func get_queue_url() string {
	queue_url, exists := os.LookupEnv(`QUEUE_URL`)
	if !exists || !strings.HasPrefix("http", queue_url) {
		return "http://localhost:4566/000000000000/notices"
	}
	return queue_url
}

func get_queue_region() string {
	queue_region, exists := os.LookupEnv(`QUEUE_REGION`)
	if !exists {
		return "us-east-1"
	}
	return queue_region
}

func main() {

	queue_url := get_queue_url()
	queue_region := get_queue_region()
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Endpoint:    aws.String(queue_url),
			Region:      aws.String(queue_region),
			Credentials: credentials.AnonymousCredentials,
		},
	}))
	svc := sqs.New(sess)

	router := mux.NewRouter()

	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		var response Response
		response.Message = "I am a healthy one."
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return
		}
		w.Write(jsonResponse)

	}).Methods("GET")

	router.HandleFunc("/notice", func(w http.ResponseWriter, r *http.Request) {

		dec := json.NewDecoder(r.Body)
		var notice Notice
		dec.Decode(&notice)

		fmt.Printf("\nnotice received:\n%s\n", notice.Origin)

		result, err := svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
				"Message": {
					DataType:    aws.String("String"),
					StringValue: aws.String(notice.Message),
				},
				"Channel": {
					DataType:    aws.String("String"),
					StringValue: aws.String(notice.Channel),
				},
			},
			MessageBody: aws.String(notice.Origin),
			QueueUrl:    aws.String(queue_url),
		})
		if err != nil {
			fmt.Printf("\nfailed to send message to queue:\n%s\n", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Printf("\nmessage sent to queue\n")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonResponse, err := json.Marshal(result)
		if err != nil {
			return
		}
		w.Write(jsonResponse)

	}).Methods("POST")

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
