package freshbooks_test

import (
	. "github.com/jevonearth/go-freshbooks/freshbooks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"io/ioutil"
)

var _ = Describe("freshbooks.client", func() {

	var (
		testkey string
		testorg string
	)

	BeforeEach(func() {
		testorg = "testorg"
		testkey = "testkey!@#123"
	})

	It("Creating client ready for production use", func() {

		client := NewClient(testorg, testkey)
		url := fmt.Sprintf("https://%s.freshbooks.com/api/2.1/xml-in", testorg)

		Expect(client.BaseURL).To(Equal(url))
		Expect(client.Org).To(Equal(testorg))
		Expect(client.Key).To(Equal(testkey))
	})

	It("Returns an error when given a nil Request", func() {
		client := NewClient(testorg, testkey)
		_, err := client.NewRequest(nil)
		Expect(err.Error()).To(Equal("NewRequest requires a non nil Request"))
	})

	It("Creates a new POST request", func() {
		client := NewClient(testorg, testkey)
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
