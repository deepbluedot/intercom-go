package intercom

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/deepbluedot/intercom-go/interfaces"
	"github.com/google/go-querystring/query"
)

type IntercomHTTPClient struct {
	*http.Client
	BaseURI     *string
	AppID       string
	AccessToken string
}

func NewIntercomHTTPClient(appID, accessToken string, baseURI, clientVersion *string) IntercomHTTPClient {
	return IntercomHTTPClient{
		Client:      &http.Client{},
		AppID:       appID,
		AccessToken: accessToken,
		BaseURI:     baseURI,
	}
}

func (c IntercomHTTPClient) Get(url string, queryParams interface{}) ([]byte, error) {
	// Setup request
	req, _ := http.NewRequest("GET", *c.BaseURI+url, nil)
	req.Header.Add("Authorization", "Bearer "+c.AccessToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Intercom-Version", clientVersion)
	addQueryParams(req, queryParams)

	// Do request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	data, err := c.readAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, c.parseResponseError(data, resp.StatusCode)
	}
	return data, err
}

func addQueryParams(req *http.Request, params interface{}) {
	v, _ := query.Values(params)
	req.URL.RawQuery = v.Encode()
}

func (c IntercomHTTPClient) Patch(url string, body interface{}) ([]byte, error) {
	return c.postOrPatch("PATCH", url, body)
}

func (c IntercomHTTPClient) Post(url string, body interface{}) ([]byte, error) {
	return c.postOrPatch("POST", url, body)
}

func (c IntercomHTTPClient) postOrPatch(method, url string, body interface{}) ([]byte, error) {
	// Marshal our body
	buffer := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return nil, err
	}

	// Setup request
	req, err := http.NewRequest(method, *c.BaseURI+url, buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.AccessToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Intercom-Version", clientVersion)

	// Do request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	data, err := c.readAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, c.parseResponseError(data, resp.StatusCode)
	}
	return data, err
}

func (c IntercomHTTPClient) Delete(url string, queryParams interface{}) ([]byte, error) {
	// Setup request
	req, _ := http.NewRequest("DELETE", *c.BaseURI+url, nil)
	req.Header.Add("Authorization", "Bearer "+c.AccessToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Intercom-Version", clientVersion)

	// Do request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	data, err := c.readAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, c.parseResponseError(data, resp.StatusCode)
	}
	return data, err
}

func (c IntercomHTTPClient) parseResponseError(data []byte, statusCode int) IntercomError {
	errorList := interfaces.HTTPErrorList{}
	err := json.Unmarshal(data, &errorList)
	if err != nil {
		return interfaces.NewUnknownHTTPError(statusCode)
	}
	if len(errorList.Errors) == 0 {
		return interfaces.NewUnknownHTTPError(statusCode)
	}
	httpError := errorList.Errors[0]
	httpError.StatusCode = statusCode
	return httpError // only care about the first
}

func (c IntercomHTTPClient) readAll(body io.Reader) ([]byte, error) {

	b, err := io.ReadAll(body)

	return b, err
}
