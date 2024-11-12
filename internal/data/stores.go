package data

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"jax.hoangdv99/internal/constant"
)

type Store struct {
	Id        int64     `json:"id"`
	Url       string    `json:"url"`
	Platform  string    `json:"platform"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type StoreModel struct {
	DB *sql.DB
}

var (
	ErrNotFoundPlatform = errors.New("not found platform")
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
