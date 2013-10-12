package freshbooks

import (
	"encoding/xml"
	"time"
)

// InvoiceService handles communication with the issue related
// methods of the FreshBooks API.
//
// FreshBooks API docs: http://developers.freshbooks.com/docs/invoices/
type InvoicesService struct {
	client *Client
}

// Invoice represents a FreshBooks invoice
type Invoice struct {
	XMLName                 xml.Name   `xml:"invoice,omitempty"`
	Id                      *int       `xml:"invoice_id,omitempty"`
	ClientId                *int       `xml:"client_id,omitempty"`
	InvoiceNumber           *string    `xml:"number,omitempty"`
	Amount                  *float32   `xml:"amount,omitempty"`
	Currency                *string    `xml:"currency_code,omitempty"`
	Language                *string    `xml:"language,omitempty"`
	AmountOutstanding       *string    `xml:"amount_outstanding,omitempty"`
	Status                  *string    `xml:"status,omitempty"`
	Date                    *time.Time `xml:"date,omitempty"`
	Folder                  *string    `xml:"folder,omitempty"`
	CustomerReference       *string    `xml:"po_number,omitempty"`
	Discount                *string    `xml:"discount,omitempty"`
	Notes                   *string    `xml:"notes,omitempty"`
	Terms                   *string    `xml:"terms,omitempty"`
	RecordUrl               *string    `xml:"links>view,omitempty"`
	LastUpdated             *string    `xml:"updated,omitempty"`
	RecurringInvoiceProfile *string    `xml:"recurring_id,omitempty"`
	ClientName              *string    `xml:"organization,omitempty"`
	Lines                   []Line     `xml:"lines>line,omitempty"`
	//TODO - Add support for missing fields
	// `xml:"contacts"`
	// `xml:"return_uri"`
	// `xml:"first_name"`
	// `xml:"last_name"`
	// `xml:"p_street1"`
	// `xml:"p_street2"`
	// `xml:"p_city"`
	// `xml:"p_state"`
	// `xml:"p_country"`
	// `xml:"p_code"`
	// `xml:"vat_name"`
	// `xml:"vat_number"`
	// `xml:"staff_id"`
}

// Line represents a Invoice Line Item that is a child of a FreshBooks invoice
type Line struct {
	// XMLName     xml.Name `xml:"line"`
	LineId      *int     `xml:"line_id,omitempty"`
	Amount      *float32 `xml:"amount,omitempty"`
	Name        *string  `xml:"name,omitempty"`
	Description *string  `xml:"description,omitempty"`
	UnitCost    *float32 `xml:"unit_cost,omitempty"`
	Quantity    *float32 `xml:"quantity,omitempty"`
	Tax1Name    *string  `xml:"tax1_name,omitempty"`
	Tax2Name    *string  `xml:"tax2_name,omitempty"`
	Tax1Percent *int     `xml:"tax1_percent,omitempty"`
	Tax2Percent *int     `xml:"tax2_percent,omitempty"`
	Type        *string  `xml:"type,omitempty"`
}

// Get a single Invoice
//
// FreshBooks API Docs: http://developers.freshbooks.com/docs/invoices/#invoice.get
func (s *InvoicesService) Get(id int) (*Invoice, *Response, error) {

	var getInvoiceRequest struct {
		Request
		XMLName xml.Name `xml:"request"`
		Id      int      `xml:"invoice_id"`
	}
	getInvoiceRequest.Id = id
	getInvoiceRequest.Method = "invoice.get"

	req, err := s.client.NewRequest(&getInvoiceRequest)

	invResp := new(Invoice)
	resp, err := s.client.Do(req, invResp)
	if err != nil {
		return nil, resp, err
	}

	return invResp, resp, err
}