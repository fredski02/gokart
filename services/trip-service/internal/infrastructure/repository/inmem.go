package repository

import (
	"context"

	"gokart/services/trip-service/internal/domain"
)

type inMemRepository struct {
	trips     map[string]*domain.TripModel
	rideFares map[string]*domain.RideFareModel
}

func NewInMemRepository() *inMemRepository {
	return &inMemRepository{
		trips:     make(map[string]*domain.TripModel),
		rideFares: make(map[string]*domain.RideFareModel),
	}
}

func (memStore *inMemRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	memStore.trips[trip.ID.Hex()] = trip
	return trip, nil
}
