package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	cart "task.com/handlers/Cart"
	category "task.com/handlers/Category"
	inventory "task.com/handlers/Inventory"
	product "task.com/handlers/Product"
)

func Start() {

	router := mux.NewRouter()

	// Route handles & endpoints
	// Product
	router.HandleFunc("/insertproduct/", product.InsertProduct).Methods("POST")
	router.HandleFunc("/getproduct/{id}/", product.GetProduct).Methods("GET")
	router.HandleFunc("/getproducts/{pageno}/", product.GetProducts).Methods("GET")
	router.HandleFunc("/deleteproduct/{id}/", product.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/updateproduct/", product.UpdateProduct).Methods("PUT")

	//Cateory
	router.HandleFunc("/insertcategory/", category.InsertCategory).Methods("POST")
	router.HandleFunc("/getcategory/", category.GetCatgory).Methods("GET")
	router.HandleFunc("/deletecategory/{id}/", category.DeleteCategory).Methods("DELETE")
	router.HandleFunc("/updatecategory/", category.UpdateCategory).Methods("PUT")

	//Inventory
	router.HandleFunc("/insertinventory/", inventory.InsertInventory).Methods("POST")
	router.HandleFunc("/getinventory/", inventory.GetInventory).Methods("GET")
	router.HandleFunc("/deleteinventory/{id}/", inventory.DeleteInventory).Methods("DELETE")
	router.HandleFunc("/updateinventory/", inventory.UpdateInventory).Methods("PUT")

	//Cart
	router.HandleFunc("/cart/create_ref_id/", cart.CreateRefId).Methods("POST")
	router.HandleFunc("/cart/get_cart/{ref_id}/", cart.Getcart).Methods("GET")
	router.HandleFunc("/addtocart/", cart.AddtoCart).Methods("POST")
	router.HandleFunc("/additemstocart/", cart.AddItemstoCart).Methods("POST")
	router.HandleFunc("/deletecart/", cart.DeleteCart).Methods("DELETE")

	// serve the app
	fmt.Println("Listening in port 7171:")
	log.Fatal(http.ListenAndServe(":7171", router))
}
