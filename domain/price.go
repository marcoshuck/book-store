package domain

type Price struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}
