package contaazul

type PaymentType string

const (
	Cash  PaymentType = "CASH"
	Times PaymentType = "TIMES"
)

type Payment struct {
	Type PaymentType `json:"type"`
	// TODO mapear installments
}
