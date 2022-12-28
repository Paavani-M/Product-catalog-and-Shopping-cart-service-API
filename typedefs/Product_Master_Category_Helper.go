package typedefs

type Product_Master_Category_Helper struct {
	Product_id    int               `json:"product_id"`
	Name          string            `json:"name"`
	Specification map[string]string `json:"specification"`
	Sku           string            `json:"sku"`
	Category_name string            `json:"category_name"`
	Price         float32           `json:"price"`
	//Category_id   int               `json:"category_id"`
}
