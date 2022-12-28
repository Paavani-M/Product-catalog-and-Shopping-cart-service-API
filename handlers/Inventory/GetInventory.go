package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	db := helpers.SetupDB()

	fmt.Println("Getting inventory db")
	rows, err := db.Query("SELECT * FROM inventory")
	defer rows.Close()
	// check errors
	helpers.CheckErr(err)

	if err != nil {
		log.Fatal(err)
	}

	inventory_result := []typedefs.Inventory{}
	for rows.Next() {
		newInventory := typedefs.Inventory{}

		err = rows.Scan(&newInventory.Product_id, &newInventory.Quantity)

		if err != nil {
			fmt.Println(err)
		}

		inventory_result = append(inventory_result, newInventory)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory_result)
}
