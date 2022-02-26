package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gorilla/mux"
)

type Request struct {
	Notices []Notice `json:"notices"`
}

type Response struct {
	Notices []Notice `json:"notices"`
}

type Notice struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Channel string `json:"channel"`
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "API is up and running")

	}).Methods("GET")

	router.HandleFunc("/notice", func(w http.ResponseWriter, r *http.Request) {

		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Config: aws.Config{
				Region:      aws.String("us-east-1"),
				Credentials: credentials.AnonymousCredentials,
				Endpoint:    aws.String("http://localhost:4566/000000000000/tornados"),
			},
		}))

		svc := sqs.New(sess)

		dec := json.NewDecoder(r.Body)
		var notice Notice
		dec.Decode(&notice)

		_, err := svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
				"Id": {
					DataType:    aws.String("String"),
					StringValue: aws.String(notice.Id),
				},
				"Message": {
					DataType:    aws.String("String"),
					StringValue: aws.String(notice.Message),
				},
				"Channel": {
					DataType:    aws.String("String"),
					StringValue: aws.String(notice.Channel),
				},
			},
			MessageBody: aws.String("Notices from a screamming guy."),
			QueueUrl:    aws.String("http://localhost:4566/000000000000/tornados"),
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Got error")
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "All good")

	}).Methods("POST")

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
