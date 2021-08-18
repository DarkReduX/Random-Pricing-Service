package main

import (
	"main/src/internal/data"
	"main/src/internal/data/repository"
	"main/src/service"
)

func main() {
	pr := map[string]data.SymbolPrice{
		"1": data.SymbolPrice{
			Uuid:   "321",
			Symbol: "apple",
		},
		"2": data.SymbolPrice{
			Uuid:   "321",
			Symbol: "lemon",
		},
		"3": data.SymbolPrice{
			Uuid:   "321",
			Symbol: "samsung",
		},
		"4": data.SymbolPrice{
			Uuid:   "321",
			Symbol: "bbbbbbbbbbbbbbbbbban",
		},
	}
	repository := repository.PriceRepository{}
	repository.ConnectRepository()
	randomPrice := service.RandomPrice{Repository: repository}
	randomPrice.UpdatePriceLoop(pr)
}
