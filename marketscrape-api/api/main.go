package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/iunary/fakeuseragent"
)

type Request struct {
	ID string `json:"id"`
}

func constructRequest(id string) string {
	var body string

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(&url.URL{
				Scheme: "http",
				User:   url.UserPassword(os.Getenv("PROXY_USERNAME"), os.Getenv("PROXY_PASSWORD")),
				Host:   os.Getenv("PROXY_HOST"),
			}),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	payload := fmt.Sprintf(`
		&variables={
			"targetId":"%s"
		}
		&doc_id=7616889011758848`, id)

	req, err := http.NewRequest("POST", "https://www.facebook.com/api/graphql/", strings.NewReader(url.PathEscape(payload)))
	if err != nil {
		return "Error creating request"
	}

	req.Header.Set("User-Agent", fakeuseragent.RandomUserAgent())
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-CA,en-US;q=0.7,en;q=0.3")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-FB-Friendly-Name", "MarketplacePDPContainerQuery")
	req.Header.Set("Origin", "https://www.facebook.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.facebook.com/marketplace/?ref=app_tab")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Priority", "u=1")
	req.Header.Set("TE", "trailers")

	resp, err := client.Do(req)
	if err != nil {
		return "Error sending request"
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response"
	}
	body = string(bodyBytes)

	return body
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req Request
	var body string
	statusCode := 200

	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		body = "Error unmarshalling request"
		statusCode = 400
	}

	body = constructRequest(req.ID)
	if strings.HasPrefix(body, "Error") {
		statusCode = 500
	}

	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: statusCode,
	}, nil
}

func main() {
	lambda.Start(handler)
}
