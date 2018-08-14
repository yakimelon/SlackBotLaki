package main

import (
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
	"log"
	"strings"
	"net/url"
)

var (
	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

type SlackMessage struct {
	Text string `json:"text"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var tw string
	var un string
	for _, value := range strings.Split(request.Body, "&") {
		param := strings.Split(value, "=")
		if param[0] == "trigger_word" {
			tw, _ = url.QueryUnescape(param[1])
		}
		if param[0] == "user_name" {
			un, _ = url.QueryUnescape(param[1])
		}
	}

	if un == "slackbot" {
		return events.APIGatewayProxyResponse {}, nil
	}

	var text string
	if tw == "おやすみ" {
		text = "おやすみなさいっ！せんぱい？(o'▽'o)ゝ🐸💕🐸"
	} else if tw == "疲れた" || tw == "つかれた" {
		text = "こんな時間まで開発しててかっこいいよっ！せんぱい？(o'▽'o)ゝ🐸💕🐸"
	} else if tw == "おはよう" {
		text = "せんぱいおはよーっ！今日も一日がんばろうね？(o'▽'o)ゝ🐸💕🐸"
	}

	j, err := json.Marshal(SlackMessage{Text: text})
	if err != nil {
		log.Print(err)
		return events.APIGatewayProxyResponse{Body: "エラー"}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(j),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
