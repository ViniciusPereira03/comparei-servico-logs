package app

import (
	"comparei-servico-logs/internal/domain/log"
	log_interface "comparei-servico-logs/internal/domain/log/interface"
)

type LogService struct {
	mysqlRepo log_interface.LogRepository
}

func NewLogService(mysqlRepo log_interface.LogRepository) *LogService {
	return &LogService{mysqlRepo: mysqlRepo}
}

func (s *LogService) CreateLog(log *log.Log) error {

	/**
	*	1. Capturar o score do evento informado
	*	2. Adicionar score ao payload de cadastro (insert)
	*	3. Calcular o nível do usuário após cadastro de log
	 */

	return nil
}
