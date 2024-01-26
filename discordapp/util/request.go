package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Status int
	Header http.Header
	Body   []byte
}

// MakeRequest sends the passed request using the passed client & unmarshals the response into unmarshalTo.
//
// If unmarshalTo is nil the response will not be unmarshaled.
func MakeRequest(req *http.Request, client *http.Client, unmarshalTo any) (*Response, error) {
	if req.Header == nil {
		req.Header = http.Header{}
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}
	response := &Response{
		Status: resp.StatusCode,
		Header: resp.Header,
		Body:   body,
	}
	if unmarshalTo == nil {
		return response, nil
	}
	err = json.Unmarshal(body, unmarshalTo)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %w", err)
	}
	return response, nil
}
