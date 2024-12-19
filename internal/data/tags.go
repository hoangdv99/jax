package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Tag struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"-"`
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

func (m TagModel) ListByUserId(userId int64) ([]Tag, error) {
	query := `SELECT id, name FROM tags where user_id = ?;`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (m TagModel) GetByName(name string) (*Tag, error) {
	query := "SELECT id, name FROM tags WHERE name = ? LIMIT 1;"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tag Tag
	err := m.DB.QueryRowContext(ctx, query, name).Scan(&tag.Id, &tag.Name)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (m TagModel) GetById(id int64) (*Tag, error) {
	query := "SELECT id, name FROM tags WHERE id = ? LIMIT 1;"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tag Tag
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&tag.Id, &tag.Name)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &tag, nil
}

func (m TagModel) Update(tag *Tag) error {
	query := `UPDATE tags SET name = ? WHERE id = ?`
	args := []interface{}{tag.Name, tag.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
