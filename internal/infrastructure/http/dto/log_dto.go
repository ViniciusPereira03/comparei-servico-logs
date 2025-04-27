package dto

import (
	"comparei-servico-logs/internal/domain/log"
)

type CreateLogDTO struct {
	EventID int `json:"event_id" validate:"required"`
}

func (dto *CreateLogDTO) ParseToLog(userId string) *log.Log {
	return &log.Log{
		UserID:  userId,
		EventID: dto.EventID,
	}
}
