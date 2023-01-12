package consoleinterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"

	"task.com/helpers"
	"task.com/typedefs"
)

func ProductMaster() {
	fmt.Printf("1.Add\n2.GetProduct\n3.GetProducts\n4.Update\n5.Delete\n")
	fmt.Println("Enter your choice")

	var choice int
	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println(err)
	}

	if choice == 1 {
		InsertProductCI()
	} else if choice == 2 {
		GetProductCI()
	} else if choice == 3 {
		GetProductsCI()
	} else if choice == 4 {
		UpdateProductCI()
	} else if choice == 5 {
		DeleteProductCI()
	}

}

func InsertProductCI() {
	fmt.Println("Enter the product id")
	var product_id int
	var name string

	_, err := fmt.Scan(&product_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the product name")
	_, err = fmt.Scan(&name)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the specification")
	var specification string
	_, err = fmt.Scan(&specification)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the SKU")
	var sku string
	_, err = fmt.Scan(&sku)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the Category id")
	var category_id int
	_, err = fmt.Scan(&category_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the Price")
	var price float32
	_, err = fmt.Scan(&price)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	InsertProductCIF(product_id, name, specification, sku, category_id, price)
}

func InsertProductCIF(pid int, name string, spec string, sku string, cid int, price float32) {

	if pid <= 0 || cid <= 0 || price <= 0 {
		fmt.Println("Enter a positive input")
		return
	}

	db := helpers.SetupDB()

	var price_round_off = math.Floor(float64(price)*100) / 100

	_, err := db.Exec(helpers.InsertProduct, pid, name, string(spec), sku, cid, price_round_off)

	if err != nil {
		fmt.Println("Query error")
		helpers.LogError(err)
		return
	}

	fmt.Println("Data has been inserted successfully")

	Continue()

}

//------------------------------------------------------------------------------

func GetProductCI() {
	fmt.Println("Enter the valid product id")
	var product_id int

	_, err := fmt.Scan(&product_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	GetProductCIF(product_id)
}

func GetProductCIF(pid int) {
	if pid <= 0 {
		fmt.Println("Enter a positive value")
		return
	}

	db := helpers.SetupDB()

	rows, err := db.Query(helpers.GetProduct, pid)

	if err != nil {
		fmt.Println("query error")
		helpers.LogError(err)
		return
	}

	defer rows.Close()

	if rows.Next() == false {
		fmt.Println("Product_id not found")
		return
	} else {

		product := typedefs.Product_Master_Category_Helper{}
		var spec []byte

		err = rows.Scan(&product.Product_id, &product.Name, &spec, &product.Sku, &product.Category_name, &product.Price)
		json.Unmarshal(spec, &product.Specification)

		if err != nil {
			helpers.CheckErr(err)
		}

		fmt.Println("Product ID:", product.Product_id)
		fmt.Println("Name:", product.Name)
		fmt.Println("Specification:", createKeyValuePairs(product.Specification))
		fmt.Println("Sku:", product.Sku)
		fmt.Println("Category Name:", product.Category_name)
		fmt.Println("Price:", product.Price)
	}

	Continue()
}

func createKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"", key, value)
	}
	return b.String()
}

//----------------------------------------------------------------------------------------------

func GetProductsCI() {
	fmt.Println("Enter the page number")
	var pageno int
	var noofitems int

	_, err := fmt.Scan(&pageno)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	fmt.Println("Enter the no. of items to be displayed per page")
	_, err = fmt.Scan(&noofitems)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	a := pageno
	b := noofitems

	var ls int
	if a <= 0 && b <= 0 {
		ls = GetProductsNoo(1, 20)
		b = 20
	} else if a >= 0 && b <= 0 {
		ls = GetProductsNoo(a, 20)
		b = 20
	} else if a <= 0 && b >= 0 {
		ls = GetProductsNoo(1, b)
	} else {
		ls = GetProductsNoo(a, b)
	}

	limit_start := ls

	db := helpers.SetupDB()

	fmt.Println("Getting products with minimum details...")
	rows, err := db.Query(helpers.GetProducts)
	defer rows.Close()

	if err != nil {
		fmt.Println("Query error")
		helpers.LogError(err)
		return
	}

	products := []typedefs.Product_Master{}
	for rows.Next() {
		newProduct := typedefs.Product_Master{}
		var spec []byte

		err = rows.Scan(&newProduct.Product_id, &newProduct.Name, &spec)
		json.Unmarshal(spec, &newProduct.Specification)

		if err != nil {
			helpers.CheckErr(err)
			helpers.LogError(err)
		}

		products = append(products, newProduct)
	}

	response := []map[string]any{}

	for _, v := range products {
		newProduct := map[string]any{
			"Specification": v.Specification,
			"name":          v.Name,
			"product_id":    v.Product_id,
		}
		response = append(response, newProduct)
	}

	limit_end := int(math.Min(float64(limit_start+b), float64(len(response))))
	var divider = "-----------------"
	if limit_start <= len(response)-1 {
		for _, m := range response[limit_start:limit_end] {
			for k, v := range m {
				fmt.Println(k, ":", v)
			}
			fmt.Println(divider)
		}
	} else {
		fmt.Println("Doesn't have enough products to display")

	}

	Continue()
}

func GetProductsNoo(pn int, lpp int) int {
	limit_start := (pn - 1) * lpp

	return limit_start
}

//-------------------------------------------------------------------

func UpdateProductCI() {
	fmt.Println("Enter product id that has you wish to update")
	var product_id int

	_, err := fmt.Scan(&product_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	var name string
	fmt.Println("Enter the product name")
	_, err = fmt.Scan(&name)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the specification")
	var specification string
	_, err = fmt.Scan(&specification)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the SKU")
	var sku string
	_, err = fmt.Scan(&sku)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the Category id")
	var category_id int
	_, err = fmt.Scan(&category_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	fmt.Println("Enter the Price")
	var price float32
	_, err = fmt.Scan(&price)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	UpdateProductCIF(product_id, name, specification, sku, category_id, price)
}

func UpdateProductCIF(pid int, name string, spec string, sku string, cid int, price float32) {

	db := helpers.SetupDB()

	rows, err := db.Query(helpers.GetProductAll, pid)
	defer rows.Close()

	if err != nil {
		fmt.Println("Query error")
		helpers.LogError(err)
		return
	}

	result := typedefs.Product_Master{}

	var speci []byte

	for rows.Next() {
		rows.Scan(&result.Product_id, &result.Name, &speci, &result.Sku, &result.Category_id, &result.Price)
	}

	var jsonStr []byte

	if name == "nil" {
		name = result.Name
	}

	if spec != "nil" {
		json.Unmarshal([]byte(spec), &result.Specification)
		jsonStr, _ = json.Marshal(result.Specification)
	} else {
		json.Unmarshal(speci, &result.Specification)
		jsonStr, _ = json.Marshal(result.Specification)
	}

	if sku == "nil" {
		sku = result.Sku
	}

	if cid == 0 {
		cid = result.Category_id
	}

	if price == 0 {
		price = result.Price
	}

	a, err := db.Exec(helpers.UpdateProduct, name, pid, string(jsonStr), sku, cid, price)

	if err != nil {
		fmt.Println("Query error")
		helpers.LogError(err)
		return
	}

	b, err := a.RowsAffected()
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	if b != 1 {
		fmt.Println("Product ID doesn't exists")
	} else {
		fmt.Println("Updated successfully")
	}

	Continue()

}

//------------------------------------------------------------------------------

func DeleteProductCI() {
	fmt.Println("Enter the valid product id")
	var product_id int

	_, err := fmt.Scan(&product_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
		return
	}

	if product_id <= 0 {
		fmt.Println("Enter a valid positive product id")
		return
	}

	db := helpers.SetupDB()
	result, err := db.Exec(helpers.DeleteProduct, product_id)

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

	if rows != 1 {
		fmt.Println("Product Id doesn't exists")
	} else {
		fmt.Println("Deleted a product from DB")
	}

	Continue()

}

func Continue() {
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
