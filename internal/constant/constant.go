package constant

const (
	USER_TYPE_SUPERVISOR string = "supervisor"
	USER_TYPE_USER       string = "user"
	USER_TYPE_ADMIN      string = "admin"
)

const (
	USER_STATUS_FIRESTORE_PENDING  int = -10 // migrated user needs changing password
	USER_STATUS_DEACTIVE           int = -1
	USER_STATUS_WAITING_ACTIVATION int = 0  // waiting activation by confirming email
	USER_STATUS_WAITING_APPROVAL   int = 10 // waiting admin's approval
	USER_STATUS_ACTIVE             int = 999
)

type Platform struct {
	Type          string
	ProductUrl    string
	LimitKey      string
	PageKey       string
	CollectionUrl string
}

var WOOCOMMERCE = Platform{
	Type:          "woocommerce",
	ProductUrl:    "/wp-json/wp/v2/product",
	LimitKey:      "per_page",
	PageKey:       "page",
	CollectionUrl: "/wp-json/wp/v2/categories",
}

var SHOPIFY = Platform{
	Type:          "shopify",
	ProductUrl:    "/products.json",
	LimitKey:      "limit",
	PageKey:       "page",
	CollectionUrl: "/collections.json",
}

var SHOPBASE = Platform{
	Type:          "shopbase",
	ProductUrl:    "/api/catalog/next/products.json",
	LimitKey:      "limit",
	PageKey:       "page",
	CollectionUrl: "/api/catelog/next/collections.json",
}

var LIST_PLATFORM = []Platform{WOOCOMMERCE, SHOPIFY, SHOPBASE}

var USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36"
