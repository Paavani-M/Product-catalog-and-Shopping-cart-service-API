package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"task.com/helpers"
	"task.com/typedefs"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("inside get product function")
	Id := params["id"]

	db := helpers.SetupDB()

	fmt.Println("Getting details...")

	rows, err := db.Query("SELECT product_master.product_id,product_master.name,product_master.specification,product_master.sku,category_master.name,product_master.price FROM product_master JOIN category_master ON product_master.category_id=category_master.category_id WHERE product_id=$1", Id)

	defer rows.Close()

	if rows.Next() == false {
		fmt.Println("Product_id not found")
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Product-id doesn't exists")
	} else {
		// check errors
		helpers.CheckErr(err)

		var products []typedefs.Product_Master_Category_Helper
		product := typedefs.Product_Master_Category_Helper{}
		var spec []byte

		err = rows.Scan(&product.Product_id, &product.Name, &spec, &product.Sku, &product.Category_name, &product.Price)
		json.Unmarshal(spec, &product.Specification)

		// check errors
		helpers.CheckErr(err)

		products = append(products, product)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}

}
