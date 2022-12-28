package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"task.com/helpers"
	"task.com/typedefs"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("inside getproducts function")
	pageno := params["pageno"]

	a, _ := strconv.Atoi(pageno)
	limit_start := (a - 1) * 20

	db := helpers.SetupDB()

	fmt.Println("Getting products with minimum details...")
	rows, err := db.Query("SELECT product_id,name,specification FROM product_master")
	defer rows.Close()
	// check errors
	helpers.CheckErr(err)

	if err != nil {
		log.Fatal(err)
	}

	products := []typedefs.Product_Master{}
	for rows.Next() {
		newProduct := typedefs.Product_Master{}
		var spec []byte

		err = rows.Scan(&newProduct.Product_id, &newProduct.Name, &spec)
		json.Unmarshal(spec, &newProduct.Specification)

		if err != nil {
			fmt.Println(err)
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
	limit_end := int(math.Min(float64(limit_start+20), float64(len(response))))

	if limit_start <= len(response)-1 {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response[limit_start:limit_end])
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Doesn't have products to display")
	}

}
