package user

type User struct {
	ID     string  `json:"id"`
	Status int     `json:"status"`
	Score  float32 `json:"score"`
	Level  int     `json:"level"`
}
