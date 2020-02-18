package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event Event) (string, error) {
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		err := errors.New("Error: WEBHOOK_URL environment variable must be set")
		log.Println(err)
		return "", err
	}

	message := CloudWatchMessage{}
	err := json.Unmarshal([]byte(event.Records[0].Sns.Message), &message)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Decide which color to use
	color := ""
	if message.NewStateValue == "ALARM" {
		color = "danger"
	}
	if message.NewStateValue == "OK" {
		color = "good"
	}

	text := message.NewStateValue + " - " + message.AlarmName

	description := "not set"

	// if message.AlarmDescription != nil {
	// 	description = message.AlarmDescription
	// }

	webhookContent := SlackWebhook{
		Text: text,
		Attachments: []SlackWebhookAttachment{
			SlackWebhookAttachment{
				Fallback: "AWS CloudWatch alarm event",
				Color:    color,
				Fields: []SlackWebhookField{
					SlackWebhookField{
						Title: "Name",
						Short: true,
						Value: message.AlarmName,
					},
					SlackWebhookField{
						Title: "Description",
						Short: false,
						Value: description,
					},
					SlackWebhookField{
						Title: "State change",
						Short: false,
						Value: message.OldStateValue + " -> " + message.NewStateValue,
					},
					SlackWebhookField{
						Title: "Reason for state change",
						Short: false,
						Value: message.NewStateReason,
					},
					SlackWebhookField{
						Title: "Timestamp",
						Short: false,
						Value: message.StateChangeTime,
					},
					SlackWebhookField{
						Title: "Region",
						Short: false,
						Value: message.Region,
					},
					SlackWebhookField{
						Title: "AWS Account ID",
						Short: false,
						Value: message.AWSAccountID,
					},
				},
			},
		},
	}

	body, err := json.Marshal(webhookContent)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// TODO webhook url as env var in Lambda
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Println("Slack Webhook response:")
	fmt.Println(string(body))

	return "success", nil
}

func main() {
	lambda.Start(HandleRequest)
}
