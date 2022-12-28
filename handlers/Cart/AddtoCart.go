package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func AddtoCart(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	cart := typedefs.Cart_Items{}

	err := json.Unmarshal(reqBody, &cart)
	if err != nil {
		fmt.Println(err)
	}
	var response = typedefs.Json_Response{}

	if cart.Reference_Id == "" {
		response = typedefs.Json_Response{Type: "missing", Message: "reference id has not passed"}
		json.NewEncoder(w).Encode(response)
	} else {

		db := helpers.SetupDB()

		fmt.Println("Inserting details into DB")

		result1, _ := alreadyExists(cart.Product_Id, cart.Quantity, cart.Reference_Id)
		//fmt.Println("alreadyexits:", err)

		user_existing_quantity, err := tq(cart.Reference_Id, cart.Product_Id)
		helpers.CheckErr(err)

		inventory_existing_quantity := inventory_quantity(cart.Product_Id)

		if result1 == true { //product already exists in the db just want to increase the quantity
			result, _ := canPurchase(cart.Product_Id, cart.Quantity) // checking whether stock is there
			if result == true {                                      //yes stock is there
				total_quantity := user_existing_quantity + cart.Quantity

				_, err = db.Exec("UPDATE cart_items SET quantity=$1 WHERE reference_id=$2 AND product_id=$3;", total_quantity, cart.Reference_Id, cart.Product_Id)

				reduced_quantity := inventory_existing_quantity - cart.Quantity

				// fmt.Println("inventory existing quantity", inventory_existing_quantity)
				// fmt.Println("cart quantity", cart.Quantity)
				// fmt.Println("reduced quantity", reduced_quantity)

				if reduced_quantity <= 0 {
					_, err = db.Exec("DELETE FROM inventory WHERE product_id=$1;", cart.Product_Id)
				} else {
					_, err = db.Exec("UPDATE inventory SET quantity=$1 WHERE product_id=$2;", reduced_quantity, cart.Product_Id)
				}

				response = typedefs.Json_Response{Type: "success", Message: "Added to cart!"}
				json.NewEncoder(w).Encode(response)

			} else {
				row := db.QueryRow("SELECT quantity FROM inventory WHERE product_id=$1", cart.Product_Id)

				var existing_quantity int

				row.Scan(&existing_quantity)

				response = typedefs.Json_Response{Type: "Insuifficient", Message: "Available quantity:" + fmt.Sprint(existing_quantity) + ", Enough Stock doesn't exists"}
				json.NewEncoder(w).Encode(response)
			}
		} else { // if the selected product doesnt already exists in the cart
			result, err := canPurchase(cart.Product_Id, cart.Quantity)

			helpers.CheckErr(err)
			if result != false {
				_, err := db.Exec("INSERT INTO cart_items(reference_id,product_id,quantity) VALUES($1,$2,$3);", cart.Reference_Id, cart.Product_Id, cart.Quantity)
				helpers.CheckErr(err)
				reduced_quantity := inventory_existing_quantity - cart.Quantity

				if reduced_quantity <= 0 {
					_, err = db.Exec("DELETE FROM inventory WHERE product_id=$1;", cart.Product_Id)
				} else {
					_, err = db.Exec("UPDATE inventory SET quantity=$1 WHERE product_id=$2;", reduced_quantity, cart.Product_Id)
				}

				response = typedefs.Json_Response{Type: "success", Message: "Added to cart!"}
				json.NewEncoder(w).Encode(response)

			} else {
				row := db.QueryRow("SELECT quantity FROM inventory WHERE product_id=$1", cart.Product_Id)

				var existing_quantity int

				row.Scan(&existing_quantity)

				response = typedefs.Json_Response{Type: "Insufficient", Message: "Available quantity:" + fmt.Sprint(existing_quantity) + ", Enough Stock doesn't exists"}
				json.NewEncoder(w).Encode(response)
			}
		}
	}
}

func inventory_quantity(pid int) int {
	db := helpers.SetupDB()

	row := db.QueryRow("SELECT quantity FROM inventory WHERE product_id=$1", pid)
	var quan int

	row.Scan(&quan)
	return quan
}

// to retrieve the current quantity of the user
func tq(rid string, pid int) (int, error) {
	db := helpers.SetupDB()

	row := db.QueryRow("SELECT quantity FROM cart_items WHERE reference_id = $1 AND product_id=$2", rid, pid)
	var quan int

	row.Scan(&quan)
	return quan, nil
}

func alreadyExists(pid int, quantity int, rid string) (bool, error) {
	var enough bool

	db := helpers.SetupDB()
	if err := db.QueryRow("SELECT quantity FROM cart_items WHERE product_id = $1 AND reference_id=$2", pid, rid).Scan(&enough); err != nil {
		//fmt.Println(err)
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("product_id %d doesn't exists for reference_id %s", pid, rid)
		}
	}
	return enough, nil
}

func canPurchase(id int, quantity int) (bool, error) {
	var enough bool

	db := helpers.SetupDB()

	if err := db.QueryRow("SELECT (quantity >= $1) FROM inventory WHERE product_id = $2", quantity, id).Scan(&enough); err != nil {
		//	fmt.Println(err)
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("canPurchase %d: unknown album", id)
		}
	}
	return enough, nil
}
