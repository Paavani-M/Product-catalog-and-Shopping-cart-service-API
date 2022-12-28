package typedefs

type Product_Master struct {
	Product_id    int               `json:"product_id"`
	Name          string            `json:"name"`
	Specification map[string]string `json:"specification"`
	Sku           string            `json:"sku"`
	Category_id   int               `json:"category_id"`
	Price         float32           `json:"price"`
	//Category_id   int               `json:"category_id"`
}
