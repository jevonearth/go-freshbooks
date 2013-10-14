package freshbooks

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	userAgent      = "go-freshbooks/" + libraryVersion
)

// Client manages communication with the FreshBooks API.
type Client struct {

	// Token is the unique authorization token assigned to your FreshBooks account.
	// Every request made the FreshBooks uses this token for HTTP basic authorization
	// The token is based on your freshBooks password. It your FreshBooks password
	// changes, so will your token.
	Key string

	// ServiceURL is the single point of entry to the FreshBooks API , it is
	// derived from your account URL
	ServiceURL *url.URL

	// User agent used when communication with the FreshBooks API.
	UserAgent string

	// HTTP client used to communicate with the API.
	client *http.Client

	// Services used for talking to different resources in the FreshBooks API.
	Invoices *InvoicesService
}

// NewClient Produces a new FreshBooks API client. Caller must provide the ServiceURL,
// and a Authorization Token.
func NewClient(serviceURL, key string) *Client {

	c := &Client{ServiceURL: serviceURL, Key: key, UserAgent: userAgent}

	c.Invoices = &InvoicesService{client: c}

	return c
}

// NewRequest creates an API request. All FreshBooks requests are POSTs.
func (c *Client) NewRequest(body interface{}) (*http.Request, error) {
	if body == nil {
		return nil, errors.New("newrequest requires a non nil request")
	}
	buf := new(bytes.Buffer)

	err := xml.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.ServiceURL, buf)

	req.Header.Add("User-Agent", c.UserAgent)
	req.SetBasicAuth(c.Key, "X")

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or retured as an error if
// and API error has occourred
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	_, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{}, nil
}

// Response is a FreshBooks API response. This wraps the standard http.Response
// returned from FreshBooks.
type Response struct {
	*http.Response
}

// Request represents the base FreshBooks API request body. This struct is used to
// compose specific resource requests by embedding.
type Request struct {
	XMLName   xml.Name `xml:"request"`
	Method    string   `xml:"method,attr"`
	Page      int      `xml:"page,omitempty"`
	PageCount int      `xml:"-"`
	PageSize  int      `xml:"per_page,omitempty"`
}
