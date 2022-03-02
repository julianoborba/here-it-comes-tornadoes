package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "tornado/docs"
)

var SQS_CLIENT = sqs.New(session.Must(session.NewSessionWithOptions(session.Options{
	Config: aws.Config{
		Endpoint: aws.String(os.Getenv(`QUEUE_URL`)),
		Region:   aws.String(os.Getenv(`QUEUE_REGION`)),
	},
})))

// Response defines the model for generic response message
type Response struct {
	Message string `json:"message" example:"My generic user friendly message about response from some request"`
}

// Notice defines the model for a notice
type Notice struct {
	Subject string `json:"subject" example:"An EC2 instance is behaving in a manner indicating it is being used to perform a Denial of Service (DoS) attack using the TCP protocol."`
	Finding string `json:"finding" example:"ThreatPurpose:ResourceTypeAffected/ThreatFamilyName.DetectionMechanism!Artifact"`
	Channel string `json:"channel" example:"C05002EAE"`
}

// EnqueuedMessage defines the model for a AWS SQS enqueued message
type EnqueuedMessage struct {
	MD5OfMessageAttributes       *string `json:"attributes_md5" example:"e64461b4cb51a781f7d35414369a7bfc"`
	MD5OfMessageBody             *string `json:"body_md5" example:"f27eca4f499f59e0328f3f4ae35a4a1b"`
	MD5OfMessageSystemAttributes *string `json:"system_attributes_md5" example:"null"`
	MessageId                    *string `json:"id" example:"69069c03-8720-e75a-f386-3ca5b3d56801"`
	SequenceNumber               *string `json:"sequence" example:"null"`
}

// Health godoc
// @Summary Returns a indicator of health
// @Description Returns HTTP 200 upon and only upon a successfully completed request
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /health [get]
func health(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	var response Response
	response.Message = "I am a healthy one!"

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(jsonResponse)

}

// Notices godoc
// @Summary Enqueues a new finding notice
// @Description Parses JSON request body into a notice to be enqueued at AWS SQS for future consumption
// @Tags notices
// @Accept json
// @Produce json
// @Param notice body Notice true "Notice to enqueue"
// @Success 200 {object} EnqueuedMessage
// @Router /notices [post]
func notices(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	var notice Notice
	dec.Decode(&notice)

	result, err := SQS_CLIENT.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Finding": {
				DataType:    aws.String("String"),
				StringValue: aws.String(notice.Finding),
			},
			"Channel": {
				DataType:    aws.String("String"),
				StringValue: aws.String(notice.Channel),
			},
		},
		MessageBody: aws.String(notice.Subject),
		QueueUrl:    aws.String(os.Getenv(`QUEUE_URL`)),
	})

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var enqueued_message EnqueuedMessage
	enqueued_message.MD5OfMessageAttributes = result.MD5OfMessageAttributes
	enqueued_message.MD5OfMessageBody = result.MD5OfMessageBody
	enqueued_message.MD5OfMessageSystemAttributes = result.MD5OfMessageSystemAttributes
	enqueued_message.MessageId = result.MessageId
	enqueued_message.SequenceNumber = result.SequenceNumber

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(enqueued_message)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(jsonResponse)

}

// @title Tornado API
// @version 1.0.0
// @description API to enqueue findings with AWS SQS
// @host localhost:8080
// @BasePath /
func main() {

	router := mux.NewRouter()

	router.HandleFunc("/health", health).Methods("GET")

	router.HandleFunc("/notices", notices).Methods("POST")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
