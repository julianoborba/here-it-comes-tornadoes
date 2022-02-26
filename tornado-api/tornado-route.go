package main

import (
	"encoding/json"
	"net/http"

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
	Message string `json:"message"`
	Channel string `json:"channel"`
}

func main() {

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

		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Config: aws.Config{
				Region:      aws.String("us-east-1"),
				Credentials: credentials.AnonymousCredentials,
				Endpoint:    aws.String("http://localhost:4566/000000000000/tornados"),
			},
		}))
		svc := sqs.New(sess)
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
			MessageBody: aws.String("Notices from a screamming guy."),
			QueueUrl:    aws.String("http://localhost:4566/000000000000/tornados"),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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
