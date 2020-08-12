package contaazul

import (
	"fmt"
	"net/url"
	"time"

	"github.com/fatih/structs"
)

// ListCustomerOptions list customer options
type ListCustomerOptions struct {
	Search      string `json:"search,omitempty"        structs:"search,omitempty"`
	Name        string `json:"name,omitempty"          structs:"name,omitempty"`
	CompanyName string `json:"company_name,omitempty"  structs:"company_name,omitempty"`
	Document    string `json:"document,omitempty"      structs:"document,omitempty"`
	Page        int    `json:"page,omitempty"          structs:"page,omitempty"`
	Size        int    `json:"size,omitempty"          structs:"size,omitempty"`
}

func (lco ListCustomerOptions) AddOptions(u string) string {
	return structToQueryValues(lco, u)
}

// ListSaleOptions list sale options
type ListSaleOptions struct {
	EmissionStart time.Time  `json:"emission_start,omitempty"  structs:"emission_start,omitempty"`
	EmissionEnd   time.Time  `json:"emission_end,omitempty"    structs:"emission_end,omitempty"`
	Status        SaleStatus `json:"status,omitempty"          structs:"status,omitempty"`
	CustomerId    string     `json:"customer_id,omitempty"     structs:"customer_id,omitempty"`
	Page          int        `json:"page,omitempty"            structs:"page,omitempty"`
	Size          int        `json:"size,omitempty"            structs:"size,omitempty"`
}

func (lso ListSaleOptions) AddOptions(u string) string {
	return structToQueryValues(lso, u)
}

func structToQueryValues(value interface{}, u string) string {
	url2, _ := url.Parse(u)
	q := url2.Query()

	m := structs.Map(value)

	for k, v := range m {
		if v != nil {
			t, isTime := v.(time.Time)

			if isTime {
				if !t.IsZero() {
					q.Add(k, fmt.Sprintf("%v", t.Format(time.RFC3339)))
				}
			} else {
				q.Add(k, fmt.Sprintf("%v", v))
			}
		}
	}

	url2.RawQuery = q.Encode()
	return url2.String()
}
