package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"task.com/handlers/query_helpers"
	"task.com/helpers"
)

func AddtoCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	Reference_Id := urlQuery.Get("ref")
	Quantity := urlQuery.Get("quantity")
	Product_Id := urlQuery.Get("product_id")

	quan, err := strconv.Atoi(Quantity)
	helpers.CheckErr(err)

	p_id, err := strconv.Atoi(Product_Id)
	helpers.CheckErr(err)

	res := query_helpers.AddItemtoCart(Reference_Id, p_id, quan)

	json.NewEncoder(w).Encode(res)

}
