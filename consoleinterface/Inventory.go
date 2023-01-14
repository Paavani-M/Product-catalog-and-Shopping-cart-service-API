package consoleinterface

import (
	"fmt"

	"task.com/helpers"
	"task.com/typedefs"
)

func Inventory() {
	fmt.Printf("1.Add\n2.Get\n3.Update\n4.Delete\n")
	fmt.Println("Enter your choice")

	var choice int
	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		InsertInventoryCI()
	} else if choice == 2 {
		GetInventoryCI()
	} else if choice == 3 {
		UpdateInventoryCI()
	} else if choice == 4 {
		DeleteInventoryCI()
	}

}

func InsertInventoryCI() {
	fmt.Println("Enter product id")
	var product_id int
	var quantity int

	_, err := fmt.Scan(&product_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter quantity")
	_, err = fmt.Scan(&quantity)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	if product_id <= 0 || quantity <= 0 {
		fmt.Println("Enter positive value")
		return
	}

	db := helpers.SetupDB()

	result, err := db.Exec(helpers.InsertUpdateInventory, product_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	if rows == 1 {
		fmt.Println("Product Id already exists, please update it")
		return
	}

	_, err = db.Exec(helpers.InsertInventory, product_id, quantity)
	if err != nil {
		fmt.Println("Query error")
		helpers.LogError(err)
		return
	}

	fmt.Println("Your data has been inserted successfuly")

	ContinueI()
}

//------------------------------------------------------------------------------

func GetInventoryCI() {
	db := helpers.SetupDB()

	rows, err := db.Query(helpers.GetInventory)
	defer rows.Close()

	if err != nil {
		fmt.Println("Query Error")
		helpers.LogError(err)
		return
	}

	inventory_result := []typedefs.Inventory{}
	for rows.Next() {
		newInventory := typedefs.Inventory{}

		err = rows.Scan(&newInventory.Product_id, &newInventory.Quantity)

		if err != nil {
			helpers.CheckErr(err)
			helpers.LogError(err)
		}

		inventory_result = append(inventory_result, newInventory)
	}

	response := []map[string]any{}

	for _, v := range inventory_result {
		newInventory := map[string]any{
			"Product Id": v.Product_id,
			"Quantity":   v.Quantity,
		}
		response = append(response, newInventory)
	}

	var divider = "-----------------"
	for _, m := range response {
		for k, v := range m {
			fmt.Println(k, ":", v)
		}
		fmt.Println(divider)
	}

	ContinueI()

}

//---------------------------------------------------------------------

func UpdateInventoryCI() {
	fmt.Println("Enter the product id")
	var product_id int
	var Quantity int

	_, err := fmt.Scan(&product_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the quantity")
	_, err = fmt.Scan(&Quantity)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	if product_id <= 0 || Quantity <= 0 {
		fmt.Println("Enter a positive input")
		return
	}

	db := helpers.SetupDB()

	result, err := db.Exec(helpers.UpdateInventory, Quantity, product_id)

	if err != nil {
		fmt.Println("Query Error")
		helpers.LogError(err)
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	if rows != 1 {
		fmt.Println("Product ID doesn't exists")
	} else {
		fmt.Println("Updated successfully")
	}

	ContinueI()

}

//-------------------------------------------------------------------------------------------------------

func DeleteInventoryCI() {
	fmt.Println("Enter the product id")
	var product_id int

	_, err := fmt.Scan(&product_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	if product_id <= 0 {
		fmt.Println("Enter a positive input")
		return
	}

	db := helpers.SetupDB()

	result, err := db.Exec(helpers.DeleteInventory, product_id)
	if err != nil {
		fmt.Println("Query Error")
		helpers.LogError(err)
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	if rows == 0 {
		fmt.Println("Product id doesn't exists")
	} else {
		fmt.Println("Deleted successfully")
	}

	ContinueI()

}

func ContinueI() {
	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err := fmt.Scan(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		ConsoleMain()
	} else {
		return
	}
}
