package freshbooks

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
)

const (
	libraryVersion = "0.1"
	userAgent      = "go-freshbooks/" + libraryVersion
)

type Client struct {
	Org       string
	Key       string
	BaseURL   string
	UserAgent string
	client    http.Client

	Invoices *InvoicesService
}

func NewClient(org, key string) *Client {

	baseURL := fmt.Sprintf("https://%s.freshbooks.com/api/2.1/xml-in", org)

	c := &Client{BaseURL: baseURL, Org: org, Key: key, UserAgent: userAgent}

	c.Invoices = &InvoicesService{client: c}

	return c
}

// NewRequest creates an API request. All FreshBooks requests are POSTs.
func (c *Client) NewRequest(body interface{}) (*http.Request, error) {
	if body == nil {
		return nil, errors.New("NewRequest requires a non nil Request")
	}
	buf := new(bytes.Buffer)

	err := xml.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL, buf)

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
