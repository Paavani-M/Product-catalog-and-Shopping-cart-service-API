package typedefs

type Cart_Items struct {
	Reference_Id string `json:"reference_id"`
	Product_Id   int    `json:"product_id"`
	Quantity     int    `json:"quantity"`
}
