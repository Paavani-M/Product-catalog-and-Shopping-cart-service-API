package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	delete := typedefs.Delete_Cart{}

	err := json.Unmarshal(reqBody, &delete)
	if err != nil {
		fmt.Println(err)
	}

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	fmt.Println("Deleting from cart....")

	result, err := db.Exec("DELETE FROM cart_items where reference_id=$1 and product_id = $2", delete.Reference_Id, delete.Product_Id)

	helpers.CheckErr(err)

	rows, err := result.RowsAffected()

	helpers.CheckErr(err)

	if rows != 1 {
		response = typedefs.Json_Response{Type: "missing", Message: "product id or reference_id doesn't exists"}
	} else {
		response = typedefs.Json_Response{Type: "success", Message: "Deleted successfully!"}
	}
	json.NewEncoder(w).Encode(response)

}
