package event

type Event struct {
	ID    int     `json:"id"`
	Event string  `json:"event"`
	Score float32 `json:"score"`
}
