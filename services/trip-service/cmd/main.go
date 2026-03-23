package main

import (
	"context"
	"log"
	"time"

	"gokart/services/trip-service/internal/domain"
	"gokart/services/trip-service/internal/infrastructure/repository"
	"gokart/services/trip-service/internal/service"
)

func main() {
	ctx := context.Background()
	inMemRepo := repository.NewInMemRepository()
	srv := service.NewTripService(inMemRepo)
	fare := &domain.RideFareModel{
		UserID: "42",
	}
	t, err := srv.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	// temporary.
	for {
		time.Sleep(time.Second)
	}
}
