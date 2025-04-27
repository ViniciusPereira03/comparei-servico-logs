package levels_interface

type LevelsRepository interface {
	GetLevelByScore(score float32) (int, error)
}
