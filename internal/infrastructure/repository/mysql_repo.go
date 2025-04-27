package repository

import (
	"comparei-servico-logs/internal/domain/log"
	"comparei-servico-logs/internal/domain/user"
	"database/sql"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) CreateLog(log *log.Log) error {
	_, err := r.db.Exec("INSERT INTO logs (user_id, event_id, score) VALUES (?, ?, ?)", log.UserID, log.EventID, log.Score)
	return err
}

func (r *MySQLRepository) CreateUser(user *user.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, level, score, status) VALUES (?, ?, ?)", user.ID, user.Level, 0, user.Status)
	return err
}
