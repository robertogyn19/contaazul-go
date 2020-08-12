package contaazul

type Base struct {
	Name string `json:"name"`
}

type City Base
type State Base

type Address struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	ZipCode      string `json:"zip_code"`
	Neighborhood string `json:"neighborhood"`
	City         City   `json:"city"`
	State        State  `json:"state"`
}
