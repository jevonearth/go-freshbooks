package freshbooks_test

import (
	. "github.com/jevonearth/go-freshbooks/freshbooks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
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
