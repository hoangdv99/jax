package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"jax.hoangdv99/internal/constant"
)

type Store struct {
	Id        int64     `json:"id"`
	Url       string    `json:"url"`
	Platform  string    `json:"platform"`
	IsActive  bool      `json:"isActive"`
	Tags      []Tag     `json:"tags"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type StoreModel struct {
	DB *sql.DB
}

var (
	ErrNotFoundPlatform = errors.New("not found platform")
	ErrDuplicatedStore  = errors.New("duplicated store")
)

func (m StoreModel) GetStorePlatform(storeUrl string) (string, error) {
	for _, platform := range constant.LIST_PLATFORM {
		url := fmt.Sprintf("%s/%s?%s=1&%s=1", storeUrl, platform.ProductUrl, platform.LimitKey, platform.PageKey)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			return platform.Type, nil
		}
	}

	return "", nil
}

func (m StoreModel) Insert(user *User, store *Store, tagIds []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	query := `INSERT INTO stores (url, platform, is_active) VALUES (?, ?, ?);`
	args := []any{store.Url, store.Platform, store.IsActive}
	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `INSERT INTO user_stores (user_id, store_id) VALUES (?, ?);`
	storeId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return ErrDuplicatedStore
		}
		return err
	}
	args = []any{user.Id, storeId}
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(tagIds) > 0 {
		query = `INSERT INTO user_tags (store_id, tag_id) VALUES `
		args = []any{}
		values := []string{}
		for _, tagId := range tagIds {
			values = append(values, "(?, ?)")
			args = append(args, storeId, tagId)
		}
		query += strings.Join(values, ", ")
		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (m StoreModel) GetByUrl(url string) (*Store, error) {
	query := `SELECT id, url, platform, is_active FROM stores WHERE url = ?;`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var store Store
	err := m.DB.QueryRowContext(ctx, query, url).Scan(&store.Id, &store.Url, &store.Platform, &store.IsActive)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &store, nil
}

func (m StoreModel) GetById(id int64) (*Store, error) {
	query := `SELECT id, url, platform, is_active FROM stores WHERE id = ?;`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var store Store
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&store.Id, &store.Url, &store.Platform, &store.IsActive)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &store, nil
}

func (m StoreModel) UpdateUserStore(user *User, store *Store, tagIds []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO user_stores (user_id, store_id) VALUE (?, ?);`
	args := []any{user.Id, store.Id}
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return ErrDuplicatedStore
		}
		return err
	}

	if len(tagIds) > 0 {
		query = `INSERT INTO user_tags (store_id, tag_id) VALUES `
		args = []any{}
		for _, tagId := range tagIds {
			query += "(?, ?) "
			args = append(args, store.Id, tagId)
		}
		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (m StoreModel) List(userId int64) ([]Store, error) {
	query := `
		SELECT
			s.id,
			s.url,
			s.platform,
			s.is_active,
			t.id,
			t.name
		FROM stores AS s
		INNER JOIN user_stores AS us ON us.store_id = s.id
		LEFT JOIN user_tags AS ut ON ut.store_id = s.id
		INNER JOIN tags AS t ON ut.tag_id = t.id
		WHERE us.user_id = ?;
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	storeMap := make(map[int64]*Store)
	for rows.Next() {
		var storeId int64
		var url, platform, tagName string
		var isActive bool
		var tagId int64
		err := rows.Scan(&storeId, &url, &platform, &isActive, &tagId, &tagName)
		if err != nil {
			return nil, err
		}

		if _, exists := storeMap[storeId]; !exists {
			storeMap[storeId] = &Store{
				Id:       storeId,
				Url:      url,
				Platform: platform,
				IsActive: isActive,
				Tags:     []Tag{},
			}
		}

		storeMap[storeId].Tags = append(storeMap[storeId].Tags, Tag{
			Id:   tagId,
			Name: tagName,
		})
	}

	stores := make([]Store, 0, len(storeMap))
	for _, store := range storeMap {
		stores = append(stores, *store)
	}

	return stores, nil
}

func (m StoreModel) UpdateStoreTags(storeId int64, tagIds []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	// Delete existing tags
	query := `DELETE FROM user_tags WHERE store_id = ?;`
	_, err = tx.ExecContext(ctx, query, storeId)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert new tags
	if len(tagIds) > 0 {
		query = `INSERT INTO user_tags (store_id, tag_id) VALUES `
		args := []any{}
		values := []string{}
		for _, tagId := range tagIds {
			values = append(values, "(?, ?)")
			args = append(args, storeId, tagId)
		}
		query += strings.Join(values, ", ")
		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
