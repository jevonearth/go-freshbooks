package freshbooks_test

import (
	"encoding/xml"
	"fmt"
	. "github.com/jevonearth/go-freshbooks/freshbooks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var _ = Describe("freshbooks NewClient", func() {

	var (
		serviceURL = "serviceURL"
		testToken  = "!test!Token!@#123"
	)

	It("Should produce error when empty serviceURL provided", func() {
		_, err := NewClient("", testToken, http.DefaultClient)
		Expect(err.Error()).To(Equal("no serviceURL provided"))
	})

	It("Should produce error when token and http.Client are nil", func() {
		_, err := NewClient(serviceURL, "", nil)
		Expect(err.Error()).To(Equal("newclient requires either a valid authentication token or a http.Client capabale of handling authentication"))
	})

	It("Produces a client with Basic HTTP Auth configured", func() {
		httpClient := http.DefaultClient

		client, err := NewClient(serviceURL, testToken, httpClient)

		Expect(err).To(BeNil())

		Expect(client.Token).To(Equal(testToken))
		Expect(client.ServiceURL.String()).To(Equal(serviceURL))
	})
})

var _ = Describe("freshbooks NewRequest", func() {

	var client *Client

	BeforeEach(func() {
		client, _ = NewClient("serviceURL", "token", nil)
	})

	It("Should produce error when request body is nil", func() {
		_, err := client.NewRequest(nil)
		Expect(err.Error()).To(Equal("newrequest requires a valid API request"))
	})

	It("Produces a http.Request with a XML representation of v, and valid User-Agent header", func() {

		inBody := InvoiceRequestQuery{
			Request: Request{Method: "invoice.get"},
			ID:      1,
		}
		outBody := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<request method=\"invoice.get\"><invoice_id>1</invoice_id></request>"

		req, err := client.NewRequest(inBody)

		Expect(err).To(BeNil())

		//test that body was XML encoded
		body, _ := ioutil.ReadAll(req.Body)

		Expect(string(body)).To(Equal(outBody))
		Expect(req.Header.Get("User-Agent")).To(Equal(client.UserAgent))
	})

	It("Should produce an error when passed non encodeable struct", func() {
		type T struct {
			A map[int]interface{}
		}
		_, err := client.NewRequest(&T{})
		Expect(err.Error()).To(Equal("xml: unsupported type: map[int]interface {}"))
	})

	It("Produces request with basic digest headers when supplied with a token", func() {
		inBody := Request{Method: "invoice.get"}
		req, _ := client.NewRequest(inBody)
		Expect(req.Header).To(HaveKey("Authorization"))
	})

	It("Produces request with no basic digest headers when supplied empty token", func() {
		client, _ = NewClient("serviceURL", "", http.DefaultClient)
		inBody := Request{Method: "invoice.get"}
		req, _ := client.NewRequest(inBody)
		Expect(req.Header).NotTo(HaveKey("Authorization"))
	})
})

var _ = Describe("freshbooks Do http", func() {
	var (
		// mux is the HTTP request multiplexer used with the test server.
		mux *http.ServeMux

		// client is the FreshBooks client being tested
		client *Client

		// server is a test HTTP server used to provide mock API responses
		server *httptest.Server
	)

	// setup a test HTTP server along with a freshbooks.client that is
	// configured to talk to that test server. Tests should register handlers on
	// mux which provide mock responses for the API method being tested
	BeforeEach(func() {
		// test server
		mux = http.NewServeMux()
		server = httptest.NewServer(mux)
		// freshbooks

		client, _ = NewClient("abc", "123", http.DefaultClient)
		client.ServiceURL, _ = url.Parse(server.URL)
	})
	It("Produces a POST to the local http/testing server", func() {
		// inBody := Request{Method: "testpostmethod"}

		type foo struct {
			A string
		}

		//Set up testing server
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			Expect(r.Method).To(Equal("POST"))

			fmt.Fprint(w, xml.Header+"<foo><A>b</A></foo>")
		})

		//prep request
		inBody := &foo{"a"}
		req, _ := client.NewRequest(inBody)

		//prep instance of foo to recieve response
		outBody := new(foo)

		//make request
		_, err := client.Do(req, outBody)

		Expect(outBody.A).To(Equal("b"))
		Expect(err).To(BeNil())
	})

})
