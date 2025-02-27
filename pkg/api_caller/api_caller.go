package api_caller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func PostRequest(client *http.Client, url string, payload any) ([]byte, error) {
	return PostRequestWithHeaders(client, url, payload, nil)
}

func PostRequestWithHeaders(client *http.Client, url string, payload any, headers map[string]string) ([]byte, error) {

	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bytesPayload))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("Content-Type", "application/json")

	return sendRequest(client, req)
}

func sendRequest(client *http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func () {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("failed to close response body: %v", err)
		}
	} ()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Printf("received code: %d from: %v", resp.StatusCode, req.URL)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func GetRequest(client *http.Client, url string) ([]byte, error) {
	return GetRequestWithHeaders(client, url, nil)
}

func GetRequestWithHeaders(client *http.Client, url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return sendRequest(client, req)
}