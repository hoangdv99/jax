package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"jax.hoangdv99/internal/constant"
)

type Product struct {
	Id            int64     `json:"id"`
	StoreUrl      string    `json:"storeUrl"`
	Price         string    `json:"price"`
	Images        []string  `json:"images"`
	CreatedDate   time.Time `json:"createdDate"`
	SourcePostUrl string    `json:"sourcePostUrl"`
	FeatureMedia  int64     `json:"featureMedia"`
}

type ProductModel struct {
	DB *sql.DB
}

func (m ProductModel) GetProducts(userId int64, storeIds []int64, page, limit int) ([]Product, error) {
	stores, err := m.getStoresByIds(userId, storeIds)
	if err != nil {
		return nil, err
	}

	products := []Product{}

	for _, store := range stores {
		var platform constant.Platform
		for _, p := range constant.LIST_PLATFORM {
			if p.Type == store.Platform {
				platform = p
				break
			}
		}
		url := fmt.Sprintf("%s/%s?%s=%d&%s=%d", store.Url, platform.ProductUrl, platform.LimitKey, limit, platform.PageKey, page)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", constant.USER_AGENT)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			continue
		}
		defer res.Body.Close()

		var fetchedProducts []Product
		// Assuming you have a function to parse the response body into products
		err = parseResponseBody(store.Url, res.Body, store.Platform, &fetchedProducts)
		if err != nil {
			continue
		}
		products = append(products, fetchedProducts...)
	}
	return products, nil
}

func (m ProductModel) getStoresByIds(userId int64, storeIds []int64) ([]Store, error) {
	var query string
	var args []any

	if len(storeIds) == 0 {
		query = `
            SELECT id, url, platform, is_active
            FROM stores AS s
			JOIN user_stores AS us ON s.id = us.store_id
            WHERE us.user_id = ? AND s.is_active = true
        `
		args = append(args, userId)
	} else {
		query = `
            SELECT id, url, platform, is_active
            FROM stores
            WHERE id IN (?) AND is_active = true
        `
		args = make([]any, len(storeIds))
		for i, id := range storeIds {
			args[i] = id
		}
		query = strings.Replace(query, "?", strings.Join(strings.Split(strings.Repeat("?", len(storeIds)), ""), ","), 1)
	}

	rows, err := m.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stores := []Store{}
	for rows.Next() {
		var store Store
		err := rows.Scan(&store.Id, &store.Url, &store.Platform, &store.IsActive)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stores, nil
}

func (m ProductModel) GetCollectionProducts(store Store, collectionId int64, handle string, page, limit int) ([]Product, error) {
	var platform constant.Platform
	for _, p := range constant.LIST_PLATFORM {
		if p.Type == store.Platform {
			platform = p
			break
		}
	}

	urlFormat := store.Url + platform.CollectionProductUrl
	var url string
	if platform.Type == "shopify" {
		url = fmt.Sprintf(urlFormat, handle, page, limit)
	} else {
		url = fmt.Sprintf(urlFormat, collectionId, page, limit)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", constant.USER_AGENT)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var products []Product
	err = parseResponseBody(store.Url, res.Body, store.Platform, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func parseResponseBody(storeUrl string, body io.Reader, platform string, products *[]Product) error {
	if platform == "shopify" {
		return parseProductsShopify(storeUrl, body, products)
	} else if platform == "shopbase" {
		return parseProductsShopbase(storeUrl, body, products)
	} else if platform == "woocommerce" {
		return parseProductsWooCommerce(storeUrl, body, products)
	} else {
		return fmt.Errorf("unsupported platform: %s", platform)
	}
}

func parseProductsShopify(storeUrl string, body io.Reader, products *[]Product) error {
	var response struct {
		Products []struct {
			Id       int64 `json:"id"`
			Variants []struct {
				Price string `json:"price"`
			} `json:"variants"`
			Images []struct {
				Src string `json:"src"`
			} `json:"images"`
			Handle      string `json:"handle"`
			CreatedDate string `json:"created_at"`
		} `json:"products"`
	}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&response)
	if err != nil {
		return err
	}

	for _, p := range response.Products {
		createdDate, err := time.Parse(time.RFC3339, p.CreatedDate)
		if err != nil {
			return err
		}

		imageUrls := []string{}
		for _, image := range p.Images {
			imageUrls = append(imageUrls, image.Src)
		}

		product := Product{
			Id:            p.Id,
			StoreUrl:      storeUrl,
			Price:         p.Variants[0].Price,
			Images:        imageUrls,
			CreatedDate:   createdDate,
			SourcePostUrl: storeUrl + "/products" + p.Handle,
		}
		*products = append(*products, product)
	}

	return nil
}

func parseProductsShopbase(storeUrl string, body io.Reader, products *[]Product) error {
	var response struct {
		Result struct {
			Items []struct {
				Id       int64  `json:"id"`
				Handle   string `json:"handle"`
				Variants []struct {
					Price float64 `json:"price"`
				} `json:"variants"`
				Images []struct {
					Src string `json:"src"`
				} `json:"images"`
				CreatedDate int64 `json:"created_at"`
			} `json:"items"`
		} `json:"result"`
	}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&response)
	if err != nil {
		return err
	}

	for _, p := range response.Result.Items {
		createdDate := time.Unix(p.CreatedDate, 0)

		imageUrls := []string{}
		for _, image := range p.Images {
			imageUrls = append(imageUrls, image.Src)
		}

		product := Product{
			Id:            p.Id,
			StoreUrl:      storeUrl,
			Price:         strconv.FormatFloat(p.Variants[0].Price, 'f', 2, 64),
			Images:        imageUrls,
			CreatedDate:   createdDate,
			SourcePostUrl: storeUrl + "/products" + p.Handle,
		}
		*products = append(*products, product)
	}

	return nil
}

func parseProductsWooCommerce(storeUrl string, body io.Reader, products *[]Product) error {
	var response []struct {
		Id           int64  `json:"id"`
		Link         string `json:"link"`
		FeatureMedia int64  `json:"featured_media"`
		CreatedDate  string `json:"date"`
	}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&response)
	if err != nil {
		return err
	}

	for _, p := range response {
		createdDate, err := time.Parse(time.RFC3339, p.CreatedDate+"Z")
		if err != nil {
			return err
		}

		product := Product{
			Id:            p.Id,
			StoreUrl:      storeUrl,
			Price:         "",
			Images:        []string{},
			CreatedDate:   createdDate,
			FeatureMedia:  p.FeatureMedia,
			SourcePostUrl: p.Link,
		}
		*products = append(*products, product)
	}

	return nil
}
