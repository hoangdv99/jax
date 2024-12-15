package data

import (
	"context"
	"database/sql"
	"time"
)

type Tag struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"userId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type TagModel struct {
	DB *sql.DB
}

func (m TagModel) Insert(tag *Tag) (int64, error) {
	query := `
		INSERT INTO tags(user_id, name) VALUE (?, ?);
	`
	args := []interface{}{tag.UserId, tag.Name}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}
