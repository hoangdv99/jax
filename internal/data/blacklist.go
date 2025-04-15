package data

import "database/sql"

type BlacklistItem struct {
	Id       int64  `json:"id"`
	StoreUrl string `json:"storeUrl"`
	Scope    string `json:"scope"`
}

type BlacklistModel struct {
	DB *sql.DB
}

func (m BlacklistModel) GetBlacklist() ([]BlacklistItem, error) {
	rows, err := m.DB.Query("SELECT id, store_url, scope FROM blacklist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blacklist []BlacklistItem
	for rows.Next() {
		var item BlacklistItem
		if err := rows.Scan(&item.Id, &item.StoreUrl, &item.Scope); err != nil {
			return nil, err
		}
		blacklist = append(blacklist, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return blacklist, nil
}
