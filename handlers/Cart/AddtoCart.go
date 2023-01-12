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

	a, _ := strconv.Atoi(Quantity)
	b, _ := strconv.Atoi(Product_Id)

	if a <= 0 || b <= 0 {
		helpers.SendErrResponse(helpers.Error, helpers.ValidInput, w)
		return
	}

	quan, err := strconv.Atoi(Quantity)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	p_id, err := strconv.Atoi(Product_Id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	res := query_helpers.AddItemtoCart(Reference_Id, p_id, quan)

	json.NewEncoder(w).Encode(res)

}
