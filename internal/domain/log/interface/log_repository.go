package log_interface

import "comparei-servico-logs/internal/domain/log"

type LogRepository interface {
	CreateLog(log *log.Log) error
}
