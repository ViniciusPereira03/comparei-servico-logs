package promer

import "comparei-servico-logs/internal/domain/log"

type Promer struct {
	ID     int    `json:"id"`
	UserID string `json:"user_id"`
}

func (dto *Promer) ParseToCreateLog() *log.Log {
	return &log.Log{
		UserID:           dto.UserID,
		EventID:          1,
		MercadoProdutoID: dto.ID,
	}
}

func (dto *Promer) ParseToUpdateLog() *log.Log {
	return &log.Log{
		UserID:           dto.UserID,
		EventID:          2,
		MercadoProdutoID: dto.ID,
	}
}
