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
				ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

				price.Uuid = time.Now().Unix()
				price.Ask = float32(rand.Intn(10000-9000+9000)) + rand.Float32()
				price.Bid = float32(rand.Intn((int(price.Ask+10000) - int(price.Ask-5000)) + int(price.Ask-5000)))
				priceBytes, err := json.Marshal(&price)
				if err != nil {
					return err
				}
				err = s.Repository.Client.RPush(ctx, "queueKey", priceBytes).Err()
				log.WithFields(log.Fields{
					"symbol": price.Symbol,
					"ask":    price.Ask,
					"bid":    price.Bid,
					"id":     price.Uuid,
				}).Info("Push price")
				if err != nil {
					log.WithFields(log.Fields{
						"error": err,
					}).Error("Error while push in redis stream")
				}
			}
		case <-quit:
			ticker.Stop()
			log.WithFields(log.Fields{
				"service": "random_price",
			}).Info("Stop service")
			return nil
		}
	}
}

func NewRandomPrice(priceRepository *repository.PriceRepository) *RandomPrice {
	return &RandomPrice{Repository: priceRepository}
}
