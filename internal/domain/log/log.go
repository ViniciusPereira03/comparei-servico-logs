package log

type Log struct {
	ID      int     `json:"id"`
	UserID  string  `json:"user_id"`
	EventID int     `json:"event_id"`
	Score   float32 `json:"score"`
}
