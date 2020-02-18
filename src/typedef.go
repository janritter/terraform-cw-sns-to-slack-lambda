package main

import "time"

type SlackWebhook struct {
	Text        string                   `json:"text"`
	Attachments []SlackWebhookAttachment `json:"attachments"`
}

type SlackWebhookAttachment struct {
	Fallback   string              `json:"fallback"`
	Color      string              `json:"color"`
	Fields     []SlackWebhookField `json:"fields"`
	ImageURL   string              `json:"image_url"`
	ThumbURL   string              `json:"thumb_url"`
	Footer     string              `json:"footer"`
	FooterIcon string              `json:"footer_icon"`
	Ts         int                 `json:"ts"`
}

type SlackWebhookField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Event struct {
	Records []struct {
		EventVersion         string `json:"EventVersion"`
		EventSubscriptionArn string `json:"EventSubscriptionArn"`
		EventSource          string `json:"EventSource"`
		Sns                  struct {
			SignatureVersion  string    `json:"SignatureVersion"`
			Timestamp         time.Time `json:"Timestamp"`
			Signature         string    `json:"Signature"`
			SigningCertURL    string    `json:"SigningCertUrl"`
			MessageID         string    `json:"MessageId"`
			Message           string    `json:"Message"`
			MessageAttributes struct {
				Test struct {
					Type  string `json:"Type"`
					Value string `json:"Value"`
				} `json:"Test"`
				TestBinary struct {
					Type  string `json:"Type"`
					Value string `json:"Value"`
				} `json:"TestBinary"`
			} `json:"MessageAttributes"`
			Type           string `json:"Type"`
			UnsubscribeURL string `json:"UnsubscribeUrl"`
			TopicArn       string `json:"TopicArn"`
			Subject        string `json:"Subject"`
		} `json:"Sns"`
	} `json:"Records"`
}

type CloudWatchMessage struct {
	AlarmName        string      `json:"AlarmName"`
	AlarmDescription interface{} `json:"AlarmDescription"`
	AWSAccountID     string      `json:"AWSAccountId"`
	NewStateValue    string      `json:"NewStateValue"`
	NewStateReason   string      `json:"NewStateReason"`
	StateChangeTime  string      `json:"StateChangeTime"`
	Region           string      `json:"Region"`
	OldStateValue    string      `json:"OldStateValue"`
	Trigger          struct {
		MetricName    string      `json:"MetricName"`
		Namespace     string      `json:"Namespace"`
		StatisticType string      `json:"StatisticType"`
		Statistic     string      `json:"Statistic"`
		Unit          interface{} `json:"Unit"`
		Dimensions    []struct {
			Value string `json:"value"`
			Name  string `json:"name"`
		} `json:"Dimensions"`
		Period                           int     `json:"Period"`
		EvaluationPeriods                int     `json:"EvaluationPeriods"`
		ComparisonOperator               string  `json:"ComparisonOperator"`
		Threshold                        float64 `json:"Threshold"`
		TreatMissingData                 string  `json:"TreatMissingData"`
		EvaluateLowSampleCountPercentile string  `json:"EvaluateLowSampleCountPercentile"`
	} `json:"Trigger"`
}
