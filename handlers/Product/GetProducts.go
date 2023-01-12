package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"task.com/helpers"
	"task.com/typedefs"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	pageno := r.URL.Query().Get("page")
	noofitems := r.URL.Query().Get("items_per_page")

	a, _ := strconv.Atoi(pageno)
	b, _ := strconv.Atoi(noofitems)

	var ls int
	if a <= 0 && b <= 0 {
		ls = GetProductsNo(1, 20)
		b = 20
	} else if a >= 0 && b <= 0 {
		ls = GetProductsNo(a, 20)
		b = 20
	} else if a <= 0 && b >= 0 {
		ls = GetProductsNo(1, b)
	} else {
		ls = GetProductsNo(a, b)
	}

	limit_start := ls

	db := helpers.SetupDB()

	rows, err := db.Query(helpers.GetProducts)
	defer rows.Close()

	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
		helpers.LogError(err)
		return
	}

	products := []typedefs.Product_Master{}
	for rows.Next() {
		newProduct := typedefs.Product_Master{}
		var spec []byte

		err = rows.Scan(&newProduct.Product_id, &newProduct.Name, &spec)
		json.Unmarshal(spec, &newProduct.Specification)

		if err != nil {
			helpers.CheckErr(err)
			helpers.LogError(err)
		}

		products = append(products, newProduct)
	}

	response := []map[string]any{}

	for _, v := range products {
		newProduct := map[string]any{
			"Specification": v.Specification,
			"name":          v.Name,
			"product_id":    v.Product_id,
		}
		response = append(response, newProduct)
	}
	limit_end := int(math.Min(float64(limit_start+b), float64(len(response))))

	if limit_start <= len(response)-1 {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response[limit_start:limit_end])
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(helpers.A)
	}

}

func GetProductsNo(pn int, lpp int) int {
	limit_start := (pn - 1) * lpp

	return limit_start
}
