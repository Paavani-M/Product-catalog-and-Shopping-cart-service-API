package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"task.com/helpers"
	"task.com/typedefs"
)

func Getcart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Ref_Id := params["ref_id"]

	var total float32
	db := helpers.SetupDB()

	rows, err := db.Query(helpers.Get_Cart, Ref_Id)

	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
		helpers.LogError(err)
		return
	}

	list_of_cart := []typedefs.Get_Cart{}

	for rows.Next() {
		new_cart := typedefs.Get_Cart{}

		err := rows.Scan(&new_cart.Reference_id, &new_cart.Product_Id, &new_cart.Product_Name, &new_cart.Price, &new_cart.Quantity)

		if err != nil {
			helpers.CheckErr(err)
			helpers.LogError(err)
		}

		total += (new_cart.Price * float32(new_cart.Quantity))

		list_of_cart = append(list_of_cart, new_cart)
	}

	if len(list_of_cart) == 0 {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(helpers.Refiddx)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	Cart := map[string]interface{}{
		"Cart Items":  list_of_cart,
		"Total Price": total,
	}

	json.NewEncoder(w).Encode(Cart)

}
