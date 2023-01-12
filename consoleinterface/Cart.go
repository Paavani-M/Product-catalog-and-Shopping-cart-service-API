package consoleinterface

import (
	"fmt"

	"task.com/handlers/query_helpers"
	"task.com/helpers"
	"task.com/typedefs"
)

func Cart() {
	fmt.Printf("1.AddItemtoCart\n2.AddItemstoCart\n3.Get\n4.Delete\n")
	fmt.Println("Enter your choice")

	var choice int
	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		AddtoCartCI()
	} else if choice == 2 {
		AddItemstoCartCI()
	} else if choice == 3 {
		GetCartCI()
	} else if choice == 4 {
		DeleteCartCI()
	}
}

func AddtoCartCI() {
	fmt.Println("Enter the reference id")
	var reference_id string
	var product_id int

	_, err := fmt.Scan(&reference_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the product id")
	_, err = fmt.Scan(&product_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	var quantity int
	fmt.Println("Enter the quantity")
	_, err = fmt.Scan(&quantity)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	res := query_helpers.AddItemtoCart(reference_id, product_id, quantity)

	fmt.Println(res)

	ContinueCa()
}

//----------------------------------------------------------------------------------------

func AddItemstoCartCI() {
	fmt.Println("Enter the reference id")
	var reference_id string

	_, err := fmt.Scan(&reference_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the number of products you want to add")
	var no_of_pro int

	_, err = fmt.Scan(&no_of_pro)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	var product_id int
	var quantity int
	for i := 0; i < int(no_of_pro); i++ {
		fmt.Println("Enter the product id")
		_, err := fmt.Scan(&product_id)
		if err != nil {
			helpers.CheckErr(err)
			helpers.LogError(err)
			return
		}

		fmt.Println("Enter the quantity")
		_, err = fmt.Scan(&quantity)
		if err != nil {
			helpers.CheckErr(err)
			helpers.LogError(err)
			return
		}

		res := query_helpers.AddItemtoCart(reference_id, product_id, quantity)

		fmt.Println(res)

	}

	ContinueCa()

}

//---------------------------------------------------------------------------------------------

func GetCartCI() {
	fmt.Println("Enter the reference id")
	var reference_id string

	_, err := fmt.Scan(&reference_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	var total float32
	db := helpers.SetupDB()

	rows, err := db.Query(helpers.Get_Cart, reference_id)

	if err != nil {
		fmt.Println("Query error")
		helpers.LogError(err)
		return
	}

	list_of_cart := []typedefs.Get_Cart{}

	for rows.Next() {
		new_cart := typedefs.Get_Cart{}

		err = rows.Scan(&new_cart.Reference_id, &new_cart.Product_Id, &new_cart.Product_Name, &new_cart.Price, &new_cart.Quantity)

		if err != nil {
			helpers.CheckErr(err)
			helpers.LogError(err)
		}

		total += (new_cart.Price * float32(new_cart.Quantity))

		list_of_cart = append(list_of_cart, new_cart)
	}

	if len(list_of_cart) == 0 {
		fmt.Println("Reference Id doesn't exists")
		return
	}

	var divider = "-----------------"
	for m1, m := range list_of_cart {
		//for k, v := range m {
		fmt.Println(m1, m)
		//}
		fmt.Println(divider)
	}
	fmt.Println("Total Cart Value:", total)

	ContinueCa()

}

//--------------------------------------------------------------------------------

func DeleteCartCI() {
	fmt.Println("Enter the reference id")
	var reference_id string

	_, err := fmt.Scan(&reference_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the product id")
	var product_id int

	_, err = fmt.Scan(&product_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	db := helpers.SetupDB()

	result, err := db.Exec(helpers.DeleteCart, reference_id, product_id)

	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	if rows != 1 {
		fmt.Println("Id doesn't exists")
	} else {
		fmt.Println("Deleted successfully")
	}

	ContinueCa()

}

func ContinueCa() {
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
