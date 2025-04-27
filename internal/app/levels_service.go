package app

import (
	levels_interface "comparei-servico-logs/internal/domain/levels/interface"
)

type LevelsService struct {
	mysqlRepo levels_interface.LevelsRepository
}

func NewLevelsService(mysqlRepo levels_interface.LevelsRepository) *LevelsService {
	return &LevelsService{mysqlRepo: mysqlRepo}
}

func (s *LevelsService) GetLevelByScore(score float32) (int, error) {
	return s.mysqlRepo.GetLevelByScore(score)
}
