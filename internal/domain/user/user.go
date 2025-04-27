package user

type User struct {
	ID     int     `json:"id"`
	Status int     `json:"status"`
	Score  float32 `json:"score"`
	Level  int     `json:"level"`
}
