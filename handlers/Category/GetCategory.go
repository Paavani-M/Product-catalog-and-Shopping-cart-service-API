package handlers

import (
	"encoding/json"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func GetCatgory(w http.ResponseWriter, r *http.Request) {
	db := helpers.SetupDB()

	rows, err := db.Query(helpers.GetCatgory)

	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
		helpers.LogError(err)
		return
	}

	defer rows.Close()

	categories := []typedefs.Category_Master{}
	for rows.Next() {
		newCategory := typedefs.Category_Master{}

		err = rows.Scan(&newCategory.Category_id, &newCategory.Name)

		if err != nil {
			helpers.CheckErr(err)
			helpers.LogError(err)
		}

		categories = append(categories, newCategory)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
