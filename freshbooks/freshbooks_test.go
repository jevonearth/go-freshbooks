package freshbooks_test

import (
	. "github.com/jevonearth/go-freshbooks/freshbooks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("freshbooks.client", func() {

	var (
		serviceURL string
		testkey    string
	)

	BeforeEach(func() {
		serviceURL = "serviceURL"
		testkey = "testkey!@#123"
	})

	It("Creating client ready for production use", func() {

		client := NewClient(serviceURL, testkey)

		Expect(client.ServiceURL).To(Equal(serviceURL))
		Expect(client.Key).To(Equal(testkey))
	})

	It("Returns an error when given a nil Request", func() {
		client := NewClient(serviceURL, testkey)
		_, err := client.NewRequest(nil)
		Expect(err.Error()).To(Equal("newrequest requires a non nil request"))
	})

	It("Creates a new POST request", func() {
		client := NewClient(serviceURL, testkey)
		r := Request{
			Method: "FOO",
		}
		req, err := client.NewRequest(r)

		Expect(err).To(BeNil())
		Expect(req.UserAgent()).To(Equal(client.UserAgent))

		body, _ := ioutil.ReadAll(req.Body)
		Expect(string(body)).To(Equal(`<request method="FOO"></request>`))

	})

	It("Sends an API request and returns the API response", func() {

	})

	// It("Creating client for testing use", func() {
	// 	client := NewClient(nil, testkey)

	// 	Expect(client.Org).To(Equal(nil))
	// 	Expect(client.Key).To(Equal(testkey))
	// })

})
