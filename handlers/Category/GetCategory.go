package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func GetCatgory(w http.ResponseWriter, r *http.Request) {
	db := helpers.SetupDB()

	fmt.Println("inside getcategory function")
	rows, err := db.Query("SELECT * FROM category_master")

	defer rows.Close()

	// check errors
	helpers.CheckErr(err)

	if err != nil {
		log.Fatal(err)
	}

	categories := []typedefs.Category_Master{}
	for rows.Next() {
		newCategory := typedefs.Category_Master{}

		err = rows.Scan(&newCategory.Category_id, &newCategory.Name)

		if err != nil {
			fmt.Println(err)
		}

		categories = append(categories, newCategory)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
