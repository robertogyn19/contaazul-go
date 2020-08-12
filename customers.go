package contaazul

import (
	"fmt"
	"net/http"
	"time"
)

type PersonType string
type RegistrationType string

const (
	Natural PersonType = "NATURAL"
	Legal   PersonType = "LEGAL"
)

const (
	NoContributor     RegistrationType = "NO_CONTRIBUTOR"
	Contributor       RegistrationType = "CONTRIBUTOR"
	ImmuneContributor RegistrationType = "IMMUNE_CONTRIBUTOR"
)

type Customer struct {
	ID                      string           `json:"id"`
	Name                    string           `json:"name"`
	CompanyName             string           `json:"company_name"`
	Email                   string           `json:"email"`
	BusinessPhone           string           `json:"business_phone"`
	MobilePhone             string           `json:"mobile_phone"`
	PersonType              PersonType       `json:"person_type"`
	Document                string           `json:"document"`
	IdentityDocument        string           `json:"identity_document"`
	StateRegistrationNumber string           `json:"state_registration_number"`
	StateRegistrationType   RegistrationType `json:"state_registration_type"`
	CityRegistrationNumber  string           `json:"city_registration_number"`
	DateOfBirth             string           `json:"date_of_birth"`
	Notes                   string           `json:"notes"`
	CreatedAt               time.Time        `json:"created_at"`
	Address                 Address          `json:"address"`
}

// ListCustomers list customers
func (cli *Client) ListCustomers(listOpts ListCustomerOptions) ([]Customer, error) {
	customers := make([]Customer, 0)

	url := fmt.Sprintf("%s/v1/customers", cli.baseURL)
	url = listOpts.AddOptions(url)

	rp := RequestParams{
		cli:              cli,
		method:           http.MethodGet,
		url:              url,
		target:           &customers,
		expectedHttpCode: []int{http.StatusOK},
	}

	err := executeApiRequest(rp)
	return customers, err
}
