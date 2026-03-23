package service

import (
	"context"

	"gokart/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tripService struct {
	repo domain.TripRepository
}

func NewTripService(repository domain.TripRepository) *tripService {
	return &tripService{
		repo: repository,
	}
}

func (service *tripService) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	trip := &domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserID:   fare.UserID,
		Status:   "pending",
		RideFare: fare,
	}
	return service.repo.CreateTrip(ctx, trip)
}
