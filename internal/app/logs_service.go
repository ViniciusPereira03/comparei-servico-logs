package app

import (
	event_interface "comparei-servico-logs/internal/domain/event/interface"
	levels_interface "comparei-servico-logs/internal/domain/levels/interface"
	"comparei-servico-logs/internal/domain/log"
	log_interface "comparei-servico-logs/internal/domain/log/interface"
	user_interface "comparei-servico-logs/internal/domain/user/interface"
	"comparei-servico-logs/internal/infrastructure/messaging/publisher"
)

// LogService agora recebe explicitamente todas as dependências
type LogService struct {
	mysqlRepo     log_interface.LogRepository
	eventService  event_interface.EventRepository
	userService   user_interface.UserRepository
	levelsService levels_interface.LevelsRepository
}

func NewLogService(
	mysqlRepo log_interface.LogRepository,
	es event_interface.EventRepository,
	us user_interface.UserRepository,
	ls levels_interface.LevelsRepository,
) *LogService {
	return &LogService{
		mysqlRepo:     mysqlRepo,
		eventService:  es,
		userService:   us,
		levelsService: ls,
	}
}

func (s *LogService) CreateLog(log *log.Log) error {
	event, err := s.eventService.GetEventById(log.EventID)
	if err != nil {
		return err
	}
	log.Score = event.Score

	err = s.mysqlRepo.CreateLog(log)
	if err != nil {
		return err
	}

	user, err := s.userService.GetUserById(log.UserID)
	if err != nil {
		return err
	}

	newUserScore := user.Score + log.Score

	user, err = s.userService.UpdateUserScore(user.ID, newUserScore)
	if err != nil {
		return err
	}

	newLevel, err := s.levelsService.GetLevelByScore(user.Score)
	if err != nil {
		return err
	}
	if newLevel != user.Level {
		user, err = s.userService.UpdateUserLevel(user.ID, newLevel)
		if err == nil {
			err = publisher.PubUpdateLevelUser(user)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
