package repository

import (
	"comparei-servico-logs/internal/domain/event"
	"comparei-servico-logs/internal/domain/log"
	"comparei-servico-logs/internal/domain/user"
	"database/sql"
	"fmt"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

// Logs
func (r *MySQLRepository) CreateLog(log *log.Log) error {
	_, err := r.db.Exec("INSERT INTO logs (user_id, event_id, score) VALUES (?, ?, ?)", log.UserID, log.EventID, log.Score)
	return err
}

// Users
func (r *MySQLRepository) CreateUser(user *user.User) error {
	user.Score = 0
	_, err := r.db.Exec("INSERT INTO users (id, level, score, status) VALUES (?, ?, ?, ?)", user.ID, user.Level, user.Score, user.Status)
	return err
}

func (r *MySQLRepository) GetUserById(id string) (*user.User, error) {
	var u user.User
	// Liste só as colunas que existem em user.User
	row := r.db.QueryRow(`
        SELECT id, status, score, level 
          FROM users 
         WHERE id = ?`, id,
	)

	// A ordem aqui TEM de bater com o SELECT acima
	err := row.Scan(
		&u.ID,
		&u.Status,
		&u.Score,
		&u.Level,
	)
	if err != nil {
		return nil, fmt.Errorf("GetUserById.Scan: %w", err)
	}
	return &u, nil
}

func (r *MySQLRepository) UpdateUserScore(user_id string, new_score float32) (*user.User, error) {
	_, err := r.db.Exec("UPDATE users SET score = ? WHERE id = ?", new_score, user_id)
	if err != nil {
		return nil, err
	}
	return r.GetUserById(user_id)
}

func (r *MySQLRepository) UpdateUserLevel(user_id string, new_level int) (*user.User, error) {
	_, err := r.db.Exec("UPDATE users SET level = ? WHERE id = ?", new_level, user_id)
	if err != nil {
		return nil, err
	}
	return r.GetUserById(user_id)
}

// Events
func (r *MySQLRepository) GetEventById(id int) (*event.Event, error) {
	var event event.Event
	err := r.db.QueryRow("SELECT * FROM events WHERE id = ?", id).Scan(
		&event.ID,
		&event.Event,
		&event.Score,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// Levels
func (r *MySQLRepository) GetLevelByScore(score float32) (int, error) {
	var level int
	err := r.db.
		QueryRow(
			`SELECT level 
             FROM levels 
             WHERE initial_score <= ? 
             ORDER BY initial_score DESC 
             LIMIT 1`, score,
		).
		Scan(&level)
	if err != nil {
		return 0, err
	}
	return level, nil
}
