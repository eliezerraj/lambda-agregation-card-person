package service

import (
	"github.com/rs/zerolog/log"
//	"github.com/lambda-card/internal/adapter/notification"

	"github.com/lambda-agregation-card-person/internal/repository"

)

var childLogger = log.With().Str("service", "AgregationService").Logger()

type AgregationService struct {
	agregationRepository repository.AgregationRepository
//	cardNotification notification.CardNotification
}

func NewAgregationService(agregationRepository repository.AgregationRepository) *AgregationService{
	childLogger.Debug().Msg("NewAgregationService")

	return &AgregationService{
		agregationRepository: agregationRepository,
	//	cardNotification: cardNotification,
	}
}