package query_helpers

import (
	"database/sql"
	"fmt"

	"task.com/helpers"
	"task.com/typedefs"
)

func AddItemtoCart(Reference_Id string, Product_Id int, Quantity int) typedefs.Json_Response {
	if Reference_Id == "" {
		return typedefs.Json_Response{Type: helpers.Missing, Message: helpers.Refidnp}
	} else {

		db := helpers.SetupDB()

		result1, _ := alreadyExists(Product_Id, Quantity, Reference_Id)

		user_existing_quantity, err := tq(Reference_Id, Product_Id)
		if err != nil {
			helpers.CheckErr(err)
			helpers.LogError(err)
		}

		inventory_existing_quantity := inventory_quantity(Product_Id)

		if result1 == true { //product already exists in the db just want to increase the quantity
			result, _ := canPurchase(Product_Id, Quantity) // checking whether stock is there
			if result == true {                            //yes stock is there
				total_quantity := user_existing_quantity + Quantity

				_, err = db.Exec(helpers.UpdateCartIQ, total_quantity, Reference_Id, Product_Id)

				if err != nil {
					helpers.CheckErr(err)
					helpers.LogError(err)
				}

				reduced_quantity := inventory_existing_quantity - Quantity

				if reduced_quantity <= 0 {
					_, err = db.Exec(helpers.DeleteInventory, Product_Id)
					if err != nil {
						helpers.CheckErr(err)
						helpers.LogError(err)
					}

				} else {
					_, err = db.Exec(helpers.UpdateInventory, reduced_quantity, Product_Id)
					if err != nil {
						helpers.CheckErr(err)
						helpers.LogError(err)
					}

				}

				return typedefs.Json_Response{Type: helpers.Success, Message: helpers.CartSuccess}

			} else {
				row := db.QueryRow(helpers.UpdateCart, Product_Id)

				var existing_quantity int

				row.Scan(&existing_quantity)

				return typedefs.Json_Response{Type: helpers.Insufficient, Message: helpers.Aq + fmt.Sprint(existing_quantity) + helpers.Esde}

			}
		} else { // if the selected product doesnt already exists in the cart
			result, err := canPurchase(Product_Id, Quantity)

			if err != nil {
				helpers.CheckErr(err)
				helpers.LogError(err)
			}
			if result != false {
				_, err := db.Exec(helpers.InsertCart, Reference_Id, Product_Id, Quantity)
				if err != nil {
					helpers.CheckErr(err)
					helpers.LogError(err)
				}
				reduced_quantity := inventory_existing_quantity - Quantity

				if reduced_quantity <= 0 {
					_, err = db.Exec(helpers.DeleteInventory, Product_Id)
					if err != nil {
						helpers.CheckErr(err)
						helpers.LogError(err)
					}

				} else {
					_, err = db.Exec(helpers.UpdateInventory, reduced_quantity, Product_Id)
					if err != nil {
						helpers.CheckErr(err)
						helpers.LogError(err)
					}

				}

				return typedefs.Json_Response{Type: helpers.Success, Message: helpers.CartSuccess}

			} else {
				row := db.QueryRow(helpers.UpdateCart, Product_Id)

				var existing_quantity int

				row.Scan(&existing_quantity)

				return typedefs.Json_Response{Type: helpers.Insufficient, Message: helpers.Aq + fmt.Sprint(existing_quantity) + helpers.Esde}
			}
		}
	}
}

func inventory_quantity(pid int) int {
	db := helpers.SetupDB()

	row := db.QueryRow(helpers.UpdateCart, pid)
	var quan int

	row.Scan(&quan)
	return quan
}

// to retrieve the current quantity of the user
func tq(rid string, pid int) (int, error) {
	db := helpers.SetupDB()

	row := db.QueryRow(helpers.CurrentQuantity, rid, pid)
	var quan int

	row.Scan(&quan)
	return quan, nil
}

func alreadyExists(pid int, quantity int, rid string) (bool, error) {
	var quantity_ int

	db := helpers.SetupDB()
	if err := db.QueryRow(helpers.AlreadyExists, pid, rid).Scan(&quantity_); err != nil {

		if err == sql.ErrNoRows {
			return false, fmt.Errorf("product_id %d doesn't exists for reference_id %s", pid, rid)
		}
	}

	return true, nil
}

func canPurchase(id int, quantity int) (bool, error) {
	var enough bool

	db := helpers.SetupDB()

	if err := db.QueryRow(helpers.CanPurchase, quantity, id).Scan(&enough); err != nil {

		if err == sql.ErrNoRows {
			return false, fmt.Errorf("canPurchase %d: unknown album", id)
		}
	}
	return enough, nil
}
