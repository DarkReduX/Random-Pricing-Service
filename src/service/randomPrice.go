package service

//https://ably.com/blog/event-streaming-with-redis-and-golang
import (
	"context"
	"encoding/json"
	"main/src/internal/data"
	"main/src/internal/data/repository"
	"math"
	"math/rand"
	"os"
	"time"
)

type RandomPrice struct {
	Repository repository.PriceRepository
}

func (s *RandomPrice) UpdatePriceLoop(prices map[string]data.SymbolPrice) error {
	queueKey := os.Getenv("REDIS_QUEUE_KEY")
	ticker := time.NewTicker(time.Second * 10)

	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			for _, price := range prices {
				priceBytes, err := json.Marshal(price)
				ctx, _ := context.WithTimeout(context.Background(), time.Second*1)

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
