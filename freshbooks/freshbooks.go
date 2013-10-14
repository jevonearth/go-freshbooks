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

	// ServiceURL is the single point of entry to the FreshBooks API , it is
	// derived from your account URL
	ServiceURL *url.URL

	// Token is the unique authorization token assigned to your FreshBooks account.
	// Every request made the FreshBooks uses this token for HTTP basic authorization
	// The token is based on your freshBooks password. It your FreshBooks password
	// changes, so will your token.
	Token string

	// User agent used when communication with the FreshBooks API.
	UserAgent string

	// HTTP client used to communicate with the API.
	client *http.Client

	// Services used for talking to different resources in the FreshBooks API.
	Invoices *InvoicesService
}

// NewClient Produces a new FreshBooks API client. Caller must either provide a
// Autorization Token or a http.Client that will preform OAuth 1a authentication
// for you (such as that provided by https://github.com/kurrik/oauth1a)
// If a AutorizationToken and a nil client is provided, a http.DefaultClient
// will be used.
func NewClient(serviceURL, token string, httpClient *http.Client) (*Client, error) {

	if serviceURL == "" {
		return nil, errors.New("no serviceURL provided")
	}
	url, _ := url.Parse(serviceURL)

	if token == "" && httpClient == nil {
		return nil, errors.New("newclient requires either a valid authentication token or a http.Client capabale of handling authentication")
	}

	c := &Client{client: httpClient, ServiceURL: url, Token: token, UserAgent: userAgent}

	c.Invoices = &InvoicesService{client: c}

	return c, nil
}

// NewRequest creates an API request. All FreshBooks requests are POSTs.
func (c *Client) NewRequest(body interface{}) (*http.Request, error) {
	if body == nil {
		return nil, errors.New("newrequest requires a valid API request")
	}

	xmlBody, err := xml.Marshal(body)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader([]byte(xml.Header + string(xmlBody)))

	req, err := http.NewRequest("POST", c.ServiceURL.String(), reader)

	if c.Token != "" {
		req.SetBasicAuth(c.Token, "X")
	}
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or retured as an error if
// and API error has occourred
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// so that the caller can inspect it.
		return response, err
	}

	if v != nil {
		err = xml.NewDecoder(resp.Body).Decode(v)
	}

	return response, err
}

func CheckResponse(r *http.Response) error {
	//TODO
	return nil
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

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}

// Float is a helper routine that allocates a new float32 value
// to store v ane returns a pointer to it.
func Float(v float32) *float32 {
	p := new(float32)
	*p = v
	return p
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	p := new(string)
	*p = v
	return p
}
