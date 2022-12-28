package typedefs

type Get_Cart struct {
	Reference_id string  `json:"reference_id"`
	Product_Id   int     `json:"product_id"`
	Product_Name string  `json:"product_name"`
	Price        float32 `json:"price"`
	Quantity     int     `json:"quantity"`
}
