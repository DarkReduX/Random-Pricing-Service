package main

import (
	log "github.com/sirupsen/logrus"
	"main/src/internal/data"
	"main/src/internal/repository"
	"main/src/service"
	"os"
	"time"
)

const (
	envRedisURI      = "REDIS_URI"
	envRedisUserName = "REDIS_USERNAME"
	envRedisPassword = "REDIS_PASS"
)

func main() {
	log.Infof("Start application")
	pr := map[string]data.SymbolPrice{
		"1": {
			Uuid:   time.Now().Unix(),
			Symbol: "apple",
		},
		"2": {
			Uuid:   time.Now().Unix(),
			Symbol: "lemon",
		},
		"3": {
			Uuid:   time.Now().Unix(),
			Symbol: "samsung",
		},
		"4": {
			Uuid:   time.Now().Unix(),
			Symbol: "ban",
		},
	}

	repository := repository.NewPriceRepository(os.Getenv(envRedisURI), os.Getenv(envRedisUserName), os.Getenv(envRedisPassword), 0)
	if repository == nil {
		log.WithFields(log.Fields{
			"repository": "price repository",
		}).Fatal("Couldn't create repository")
	}
	log.WithFields(log.Fields{
		"repository":    "price_repository",
		"database type": "redis",
	}).Info("Created repository ")

	randomPrice := service.NewRandomPrice(repository)

	if randomPrice == nil {
		log.WithFields(log.Fields{
			"service": "random_price",
		}).Fatal("Couldn't create service")
	}
	log.WithFields(log.Fields{
		"service": "random_price",
	}).Info("Created service")
	if err := randomPrice.UpdatePriceLoop(pr); err != nil {
		log.Fatal(err)
	}
}
