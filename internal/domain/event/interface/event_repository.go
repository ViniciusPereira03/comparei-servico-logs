package event_interface

import "comparei-servico-logs/internal/domain/event"

type EventRepository interface {
	GetEventById(id int) (*event.Event, error)
}
