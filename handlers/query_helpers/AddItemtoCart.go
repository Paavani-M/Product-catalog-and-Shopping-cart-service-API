package query_helpers

import (
	"database/sql"
	"fmt"

	"task.com/helpers"
	"task.com/typedefs"
)

func AddItemtoCart(Reference_Id string, Product_Id int, Quantity int) typedefs.Json_Response {
	if Reference_Id == "" {
		return typedefs.Json_Response{Type: "missing", Message: "reference id has not passed"}
		// json.NewEncoder(w).Encode(response)
	} else {

		db := helpers.SetupDB()

		fmt.Println("Inserting details into DB")

		result1, _ := alreadyExists(Product_Id, Quantity, Reference_Id)
		//fmt.Println("alreadyexits:", err)
		// fmt.Println(Product_Id, Reference_Id, Quantity)
		user_existing_quantity, err := tq(Reference_Id, Product_Id)
		helpers.CheckErr(err)

		inventory_existing_quantity := inventory_quantity(Product_Id)

		if result1 == true { //product already exists in the db just want to increase the quantity
			result, _ := canPurchase(Product_Id, Quantity) // checking whether stock is there
			if result == true {                            //yes stock is there
				total_quantity := user_existing_quantity + Quantity

				_, err = db.Exec("UPDATE cart_items SET quantity=$1 WHERE reference_id=$2 AND product_id=$3;", total_quantity, Reference_Id, Product_Id)

				reduced_quantity := inventory_existing_quantity - Quantity

				// fmt.Println("inventory existing quantity", inventory_existing_quantity)
				// fmt.Println("cart quantity", cart.Quantity)
				// fmt.Println("reduced quantity", reduced_quantity)

				if reduced_quantity <= 0 {
					_, err = db.Exec("DELETE FROM inventory WHERE product_id=$1;", Product_Id)
				} else {
					_, err = db.Exec("UPDATE inventory SET quantity=$1 WHERE product_id=$2;", reduced_quantity, Product_Id)
				}

				return typedefs.Json_Response{Type: "success", Message: "Added to cart!"}
				// json.NewEncoder(w).Encode(response)

			} else {
				row := db.QueryRow("SELECT quantity FROM inventory WHERE product_id=$1", Product_Id)

				var existing_quantity int

				row.Scan(&existing_quantity)

				return typedefs.Json_Response{Type: "Insuifficient", Message: "Available quantity:" + fmt.Sprint(existing_quantity) + ", Enough Stock doesn't exists"}
				// json.NewEncoder(w).Encode(response)
			}
		} else { // if the selected product doesnt already exists in the cart
			result, err := canPurchase(Product_Id, Quantity)

			// fmt.Println(result)
			// fmt.Println("inside else")
			helpers.CheckErr(err)
			if result != false {
				_, err := db.Exec("INSERT INTO cart_items(reference_id,product_id,quantity) VALUES($1,$2,$3);", Reference_Id, Product_Id, Quantity)
				helpers.CheckErr(err)
				reduced_quantity := inventory_existing_quantity - Quantity

				if reduced_quantity <= 0 {
					_, err = db.Exec("DELETE FROM inventory WHERE product_id=$1;", Product_Id)
				} else {
					_, err = db.Exec("UPDATE inventory SET quantity=$1 WHERE product_id=$2;", reduced_quantity, Product_Id)
				}

				return typedefs.Json_Response{Type: "success", Message: "Added to cart!"}
				// json.NewEncoder(w).Encode(response)

			} else {
				row := db.QueryRow("SELECT quantity FROM inventory WHERE product_id=$1", Product_Id)

				var existing_quantity int

				row.Scan(&existing_quantity)

				return typedefs.Json_Response{Type: "Insufficient", Message: "Available quantity:" + fmt.Sprint(existing_quantity) + ", Enough Stock doesn't exists"}
				// json.NewEncoder(w).Encode(response)
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
	var quantity_ int

	db := helpers.SetupDB()
	if err := db.QueryRow("SELECT quantity FROM cart_items WHERE product_id = $1 AND reference_id=$2", pid, rid).Scan(&quantity_); err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("product_id %d doesn't exists for reference_id %s", pid, rid)
		}
	}
	//fmt.Println("enough", enough)
	return true, nil
}

func canPurchase(id int, quantity int) (bool, error) {
	var enough bool

	db := helpers.SetupDB()

	// fmt.Println(id)
	// fmt.Println(quantity)

	if err := db.QueryRow("SELECT (quantity >= $1) FROM inventory WHERE product_id = $2", quantity, id).Scan(&enough); err != nil {

		if err == sql.ErrNoRows {
			return false, fmt.Errorf("canPurchase %d: unknown album", id)
		}
	}
	return enough, nil
}
