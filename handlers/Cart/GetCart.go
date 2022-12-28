package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"task.com/helpers"
	"task.com/typedefs"
)

func Getcart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("inside get cart function")
	Ref_Id := params["ref_id"]

	var total float32
	db := helpers.SetupDB()

	rows, err := db.Query("SELECT cart_items.reference_id,cart_items.product_id,product_master.name,product_master.price,cart_items.quantity FROM (cart_items JOIN product_master ON cart_items.product_id = product_master.product_id) WHERE cart_items.reference_id=$1", Ref_Id)

	helpers.CheckErr(err)

	if rows.Next() == false {
		fmt.Println("Reference_id not found")
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Reference_id doesn't exists")
	} else {

		list_of_cart := []typedefs.Get_Cart{}

		new_cart := typedefs.Get_Cart{}

		err := rows.Scan(&new_cart.Reference_id, &new_cart.Product_Id, &new_cart.Product_Name, &new_cart.Price, &new_cart.Quantity)

		helpers.CheckErr(err)

		total += (new_cart.Price * float32(new_cart.Quantity))

		list_of_cart = append(list_of_cart, new_cart)

		if len(list_of_cart) == 0 {
			json.NewEncoder(w).Encode("NO DATA FOUND!")
		}

		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(list_of_cart)
		res := fmt.Sprintln("Total cart value", total)
		err = json.NewEncoder(w).Encode(res)

	}

}
