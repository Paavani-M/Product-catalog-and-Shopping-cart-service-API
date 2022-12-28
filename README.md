
# Product catalog and Shopping cart Service API 

This is a Product catalog and Shopping cart service API that has products list in which customers can add or update or view or delete products in the cart.

# Prerequisites

Visual Studio Code\
Postman\
Go 1.19.3\
Postgresql database and pgAdmin 4

# Dependenices

go get github.com/gorilla/mux\
go get github.com/lib/pq\
go get github.com/google/uuid

# API Reference

## Product_master

### Request Body

JSON:

Product_id    (int)               
Name          (string)            
Specification (map[string]string)                                                                                                               
Sku           (string)            
Category_id   (int)               
Price         (float32)           

### InsertProduct

Inserts product details to the database : POST\
/insertproduct/

### GetProduct

Displays a product from the database using product id : GET\
/getproduct/{id}/

### GetProducts

Display all products from the database with minimum details and with page limit of 20 : GET\
/getproducts/{pageno}/

### UpdateProduct

Updates the product details by using product id : PUT\
/updateproduct/

### DeleteProduct

Deletes a product from the database using product id : DELETE\
/deleteproduct/{id}/

## Category_master

JSON:

Category_id (int)    
Name        (string) 

### InsertCategory

Inserts category details into the database : POST\
/insertcategory/

### GetCategory

Displays all available categories : GET\
/getcategory/

### UpdateCategory

Updates the category name by using category id : PUT\
/updatecategory/

### DeleteCategory

Deletes a category from database using category id: DELETE\
/deletecategory/{id}/

## Inventory

JSON:

Product_id (int)\
Quantity   (int)

### InsertInventory

Inserts iventory details into the database : POST\
/insertinventory/

### GetInventory

Displays all available stocks : GET\
/getinventory/

### UpdateInventory

Updates the quantity by using product id : PUT\
/updateinventory/

### DeleteInventory

Deletes a stock from database using product id : DELETE\
/deleteinventory/{id}/

## Cart

JSON:

Reference_Id (string)\
Product_Id   (int)    
Quantity     (int)    

### Generate cart reference id

Creating a unique user reference id using UUID and inserting it to the cart reference table with timestamp : POST\
/cart/create_ref_id/

### AddItemtoCart

Add new item to the cart or increases the existing quantity : POST\
/addtocart/

### AddItemstoCart

Add multiple items to the cart or increases the existing quantity : POST\
/additemstocart/

### GetCart

Displays the cart from the database using the reference id : GET\
/cart/get_cart/{ref_id}/

### DeleteCart

Deletes the product id from the cart using reference id : DELETE\
/deletecart/






