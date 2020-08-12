package contaazul

import (
	"fmt"
	"net/http"
	"time"
)

// SaleStatus represent sales status
type SaleStatus string

const (
	Pending   SaleStatus = "PENDING"
	Committed SaleStatus = "COMMITTED"
)

// Sale sale
type Sale struct {
	ID           string     `json:"id"`
	Number       int        `json:"number"`
	Emission     time.Time  `json:"emission"`
	Status       SaleStatus `json:"status"`
	Scheduled    bool       `json:"scheduled"`
	Customer     Customer   `json:"customer"` // TODO Criar um tipo simplificado de Customer?
	Discount     Discount   `json:"discount"`
	Payment      Payment    `json:"payment"`
	Notes        string     `json:"notes"`
	ShippingCost float64    `json:"shipping_cost"`
	Total        float64    `json:"total"`
	Seller       Seller     `json:"seller"`
}

func (cli *Client) ListSales(listOpts ListSaleOptions) ([]Sale, error) {
	sales := make([]Sale, 0)

	url := fmt.Sprintf("%s/v1/sales", cli.baseURL)
	url = listOpts.AddOptions(url)

	rp := RequestParams{
		cli:              cli,
		method:           http.MethodGet,
		url:              url,
		target:           &sales,
		expectedHttpCode: []int{http.StatusOK},
	}

	err := executeApiRequest(rp)
	return sales, err
}
