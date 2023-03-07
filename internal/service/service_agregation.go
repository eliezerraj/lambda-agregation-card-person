package service

import (
	"github.com/lambda-agregation-card-person/internal/core/domain"

)

func (s *AgregationService) AddAgregation(agregation domain.AgregationCardPerson) (*domain.AgregationCardPerson, error){
	childLogger.Debug().Msg("AddAgregation")

	// Add new card
	c, err := s.agregationRepository.AddAgregation(agregation)
	if err != nil {
		return nil, err
	}

	// Stream new card
	/*eventType := "add-new-card"
	err = s.cardNotification.PutEvent(*c, eventType)
	if err != nil {
		return nil, err
	}*/

	return c, nil
}

func (s *AgregationService) GetAgregation(agregation domain.AgregationCardPerson) (*domain.AgregationCardPerson, error){
	childLogger.Debug().Msg("GetAgregation")

	c, err := s.agregationRepository.GetAgregation(agregation)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *AgregationService) SetAgregationStatus(agregation domain.AgregationCardPerson) (*domain.AgregationCardPerson, error){
	childLogger.Debug().Msg("SetAgregationStatus")

	// Change status card, as the DB is a Dynamo de AddCard is a Upsert
	c, err := s.agregationRepository.AddAgregation(agregation)
	if err != nil {
		return nil, err
	}

	// Stream new card
	/*eventType := "change-status-card"
	err = s.cardNotification.PutEvent(*c, eventType)
	if err != nil {
		return nil, err
	}*/

	return c, nil
}