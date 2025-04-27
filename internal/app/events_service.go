package app

import (
	"comparei-servico-logs/internal/domain/event"
	event_interface "comparei-servico-logs/internal/domain/event/interface"
)

type EventService struct {
	mysqlRepo event_interface.EventRepository
}

func NewEventService(mysqlRepo event_interface.EventRepository) *EventService {
	return &EventService{mysqlRepo: mysqlRepo}
}

func (s *EventService) GetEventById(id int) (*event.Event, error) {
	return s.mysqlRepo.GetEventById(id)
}
