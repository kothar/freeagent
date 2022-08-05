package freeagent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const LiveEndpoint = "https://api.freeagent.com/v2"
const SandboxEndpoint = "https://api.sandbox.freeagent.com/v2"

type Client struct {
	HTTPClient  http.Client
	Endpoint    string
	AccessToken string
}

type errorsDTO struct {
	Errors struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	} `json:"errors"`
}

func (c *Client) get(path string, result any) error {
	parsedURL, err := url.Parse(c.Endpoint + path)
	if err != nil {
		return err
	}
	req := &http.Request{
		URL: parsedURL,
		Header: map[string][]string{
			"Authorization": {"Bearer " + c.AccessToken},
			"Accept":        {"application/json"},
		},
	}
	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	return c.parseResponse(response, result)
}

func (c *Client) post(path string, request any, response any) error {
	parsedURL, err := url.Parse(c.Endpoint + path)
	if err != nil {
		return err
	}

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	log.Println(string(body))

	req := &http.Request{
		URL:    parsedURL,
		Method: http.MethodPost,
		Header: map[string][]string{
			"Authorization": {"Bearer " + c.AccessToken},
			"Accept":        {"application/json"},
			"Content-Type":  {"application/json"},
		},
		Body: io.NopCloser(bytes.NewReader(body)),
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	return c.parseResponse(res, response)
}

func (c *Client) parseResponse(res *http.Response, response any) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode/100 != 2 {
		returnedError := &errorsDTO{}
		_ = json.Unmarshal(body, returnedError)
		return fmt.Errorf("%s: %s", res.Status, returnedError.Errors.Error.Message)
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}
	return nil
}
