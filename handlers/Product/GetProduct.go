package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"task.com/helpers"
	"task.com/typedefs"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	Id := params["id"]

	a, _ := strconv.Atoi(Id)

	if a <= 0 {
		helpers.SendErrResponse(helpers.Error, helpers.ValidInput, w)
		return
	}

	db := helpers.SetupDB()

	rows, err := db.Query(helpers.GetProduct, Id)

	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
		helpers.LogError(err)
		return
	}

	defer rows.Close()

	if rows.Next() == false {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(helpers.Idnotexits)
	} else {
		var products []typedefs.Product_Master_Category_Helper
		product := typedefs.Product_Master_Category_Helper{}
		var spec []byte

		err = rows.Scan(&product.Product_id, &product.Name, &spec, &product.Sku, &product.Category_name, &product.Price)
		json.Unmarshal(spec, &product.Specification)

		if err != nil {
			helpers.CheckErr(err)
			helpers.LogError(err)
		}

		products = append(products, product)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}

}
