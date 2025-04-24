package dto

import "comparei-servico-logs/internal/domain/log"

type CreateLogDTO struct {
	UserID  string `json:"user_id" validate:"required"`
	EventID int    `json:"event_id" validate:"required"`
}

func (dto *CreateLogDTO) ParseToLog() *log.Log {
	return &log.Log{
		UserID:  dto.UserID,
		EventID: dto.EventID,
	}
}
