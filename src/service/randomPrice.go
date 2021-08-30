package service

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"main/src/internal/data"
	"main/src/internal/repository"
	"math/rand"
	"time"
)

type RandomPrice struct {
	Repository *repository.PriceRepository
}

func (s *RandomPrice) UpdatePriceLoop(prices map[string]data.SymbolPrice) error {
	ticker := time.NewTicker(time.Second * 10)

	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			for _, price := range prices {
				priceBytes, err := json.Marshal(price)
				ctx, _ := context.WithTimeout(context.Background(), time.Second*1)

				price.Uuid = time.Now().Unix()
				price.Ask = float32(rand.Intn(10000-9000+9000)) + rand.Float32()
				price.Bid = float32(rand.Intn((int(price.Ask+10000) - int(price.Ask-5000)) + int(price.Ask-5000)))
				priceBytes, err := json.Marshal(&price)
				if err != nil {
					return err
				}
				price.Ask = float64(rand.Int()) + rand.Float64()
				price.Bid = float64(rand.Intn((math.MaxInt32 - int(price.Ask)) + int(price.Ask)))
				s.Repository.Client.RPush(ctx, queueKey, priceBytes)
			}
		case <-quit:
			ticker.Stop()

			return nil
		}
	}
}
