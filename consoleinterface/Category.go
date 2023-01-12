package consoleinterface

import (
	"fmt"

	"task.com/helpers"
	"task.com/typedefs"
)

func CategoryMaster() {
	fmt.Printf("1.Add\n2.Get\n3.Update\n4.Delete\n")
	fmt.Println("Enter your choice")

	var choice int
	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		InsertCategoryCI()
	} else if choice == 2 {
		GetCategoryCI()
	} else if choice == 3 {
		UpdateCategoryCI()
	} else if choice == 4 {
		DeleteCategoryCI()
	}
}

func InsertCategoryCI() {
	fmt.Println("Enter the category id")
	var category_id int
	var name string

	_, err := fmt.Scan(&category_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Please enter the product name")
	_, err = fmt.Scan(&name)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	if category_id <= 0 {
		fmt.Println("Enter a positive input")
		return
	}

	db := helpers.SetupDB()

	_, err = db.Exec(helpers.InsertCategory, category_id, name)
	if err != nil {
		fmt.Println("Query Error")
		helpers.LogError(err)
		return
	}

	fmt.Println("Your data has been inserted successfully")

	ContinueC()
}

//--------------------------------------------------------------------------------------------------

func GetCategoryCI() {
	db := helpers.SetupDB()

	rows, err := db.Query(helpers.GetCatgory)

	if err != nil {
		fmt.Println("Query Error")
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

	var divider = "-----------------"
	for m1, m := range categories {
		fmt.Println(m1, m)
		fmt.Println(divider)
	}

	ContinueC()
}

//--------------------------------------------------------------------------------------

func UpdateCategoryCI() {
	fmt.Println("Enter the category id")
	var category_id int
	var name string

	_, err := fmt.Scan(&category_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the category name")
	_, err = fmt.Scan(&name)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	if category_id <= 0 {
		fmt.Println("Enter positive input")
		return
	}

	db := helpers.SetupDB()

	result, err := db.Exec(helpers.UpdateCategory, name, category_id)

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
		fmt.Println("Category Id doesn't exists")
	} else {
		fmt.Println("Updated successfully")
	}

	ContinueC()

}

//---------------------------------------------------------------------------------

func DeleteCategoryCI() {
	fmt.Println("Enter the category id")
	var category_id int

	_, err := fmt.Scan(&category_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	if category_id <= 0 {
		fmt.Println("Enter positive input")
		return
	}

	db := helpers.SetupDB()

	result, err := db.Exec(helpers.DeleteCategory, category_id)
	if err != nil {
		fmt.Println("Query error")
		helpers.LogError(err)
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	if rows != 1 {
		fmt.Println("Category Id doesn't exists")
	} else {
		fmt.Println("Deleted successfully")
	}

	ContinueC()

}

func ContinueC() {
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
