package data

type SymbolPrice struct {
	Uuid   string  `json:"uuid"`
	Symbol string  `json:"symbol"`
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
}
