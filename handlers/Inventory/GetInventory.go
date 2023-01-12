package handlers

import (
	"encoding/json"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	db := helpers.SetupDB()

	rows, err := db.Query(helpers.GetInventory)
	defer rows.Close()

	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
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

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory_result)
}
