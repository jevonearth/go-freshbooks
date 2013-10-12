package freshbooks_test

import (
	. "github.com/jevonearth/go-freshbooks/freshbooks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("TestInvoicesService", func() {

	var (
		// mux is the HTTP request multiplexer used with the test server.
		mux *http.ServeMux

		// client is the FreshBooks client being tested
		client *Client

		// server is a test HTTP server used to provide mock API responses
		server *httptest.Server

		//Example XML is taken directly from the FreshBooks API documentation
		// invoiceGetRequest  = `<?xml version="1.0" encoding="utf-8"?><request method="invoice.get"><invoice_id>344</invoice_id></request>`
		// invoiceGetResponse = `<?xml version="1.0"?><response xmlns="http://www.freshbooks.com/api/" status="ok"><invoice><invoice_id>344</invoice_id><client_id>3</client_id><contacts><contact><contact_id>0</contact_id></contact></contacts><number>FB00004</number><!-- Total invoice amount, taxes inc. (Read Only) --><amount>45.6</amount><!-- Outstanding amount on invoice from partial payment, etc. (Read Only) --><amount_outstanding>0</amount_outstanding><status>paid</status><date>2007-06-23</date><po_number></po_number><discount>0</discount><notes>Due upon receipt.</notes><terms>Payment due in 30 days.</terms><currency_code>CAD</currency_code><folder>active</folder><language>en</language><url deprecated="true">https://2ndsite.freshbooks.com/view/St2gThi6rA2t7RQ</url> <!-- (Read-only) --><auth_url deprecated="true">https://2ndsite.freshbooks.com/invoices/344</auth_url> <!-- (Read-only) --><links><client_view>https://2ndsite.freshbooks.com/view/St2gThi6rA2t7RQ</client_view> <!-- (Read-only) --><view>https://2ndsite.freshbooks.com/invoices/344</view> <!-- (Read-only) --><edit>https://2ndsite.freshbooks.com/invoices/344/edit</edit> <!-- (Read-only) --></links><return_uri>http://www.example.com/callback</return_uri> <!-- (Optional) --><updated>2009-08-12 00:00:00</updated>  <!-- (Read-only) --><recurring_id>15</recurring_id> <!-- (Read-only) --><organization>ABC Corp</organization><first_name>John</first_name><last_name>Doe</last_name><p_street1>123 Fake St.</p_street1><p_street2>Unit 555</p_street2><p_city>New York</p_city><p_state>New York</p_state><p_country>United States</p_country><p_code>553132</p_code><vat_name></vat_name><vat_number></vat_number><staff_id>1</staff_id><lines><line><line_id>1</line_id>  <!-- (Read Only) line id --><amount>40</amount><!-- Line amount, taxes/discount excluding. (Read Only) --><name>Yard work</name><description>Mowed the lawn</description><unit_cost>10</unit_cost><quantity>4</quantity><tax1_name>GST</tax1_name><tax2_name>PST</tax2_name><tax1_percent>5</tax1_percent><tax2_percent>8</tax2_percent><type>Item</type></line></lines></invoice></response>`
	)

	// setup a test HTTP server along with a freshbooks.client that is
	// configured to talk to that test server. Tests should register handlers on
	// mux which provide mock responses for the API method being tested
	BeforeEach(func() {
		// test server
		mux = http.NewServeMux()
		server = httptest.NewServer(mux)
		// freshbooks
		client = NewClient("", "123")
	})

	It("Fetchs an Invoice by id", func() {
		//TODO
		_, _, _ = client.Invoices.Get(344)

		Expect(nil).To(BeNil()) //no-op
		// Expect(invoice.Id).To(Equal(344))
		// Expect(response).To(Equal(1))

	})

})
