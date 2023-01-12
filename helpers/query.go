package helpers

//product
var InsertProduct string = "INSERT INTO product_master(product_id,name,specification,sku,category_id,price) VALUES($1,$2,$3,$4,$5,$6);"
var GetProducts string = "SELECT product_id,name,specification FROM product_master ORDER BY product_id"
var GetProduct string = "SELECT product_master.product_id,product_master.name,product_master.specification,product_master.sku,category_master.name,product_master.price FROM product_master JOIN category_master ON product_master.category_id=category_master.category_id WHERE product_id=$1"
var UpdateProduct string = "UPDATE product_master SET Name=$1, Specification=$3, sku=$4, category_id=$5, price=$6 WHERE Product_id=$2;"
var DeleteProduct string = "DELETE FROM product_master where product_id = $1"
var GetProductAll string = "SELECT * FROM product_master where product_id=$1"

//inventory
var InsertInventory string = "INSERT INTO inventory(product_id,quantity) VALUES($1,$2);"
var GetInventory string = "SELECT * FROM inventory ORDER BY product_id"
var UpdateInventory string = "UPDATE inventory SET quantity=$1 WHERE product_Id=$2;"
var DeleteInventory string = "DELETE FROM inventory where product_id = $1"
var InsertUpdateInventory string = "SELECT * FROM inventory WHERE product_id=$1;"

//Category
var InsertCategory string = "INSERT INTO category_master(category_id,name) VALUES($1,$2);"
var GetCatgory string = "SELECT * FROM category_master ORDER BY category_id"
var UpdateCategory string = "UPDATE category_master SET Name=$1 WHERE Category_Id=$2;"
var DeleteCategory string = "DELETE FROM category_master where category_id = $1"

//cart
var CreateRefId string = "INSERT INTO cart_reference(reference_id,created_at) VALUES($1,$2);"
var Get_Cart string = "SELECT cart_items.reference_id,cart_items.product_id,product_master.name,product_master.price,cart_items.quantity FROM (cart_items JOIN product_master ON cart_items.product_id = product_master.product_id) WHERE cart_items.reference_id=$1"
var DeleteCart string = "DELETE FROM cart_items where reference_id=$1 and product_id = $2"
var InsertCart string = "INSERT INTO cart_items(reference_id,product_id,quantity) VALUES($1,$2,$3);"

//product already exists in the db just want to increase the quantity and stock is there
var UpdateCartIQ string = "UPDATE cart_items SET quantity=$1 WHERE reference_id=$2 AND product_id=$3;"

//stock is not there
var UpdateCart string = "SELECT quantity FROM inventory WHERE product_id=$1"

// to retrieve the current quantity of the user
var CurrentQuantity string = "SELECT quantity FROM cart_items WHERE reference_id = $1 AND product_id=$2"

// alreadyexists
var AlreadyExists string = "SELECT quantity FROM cart_items WHERE product_id = $1 AND reference_id=$2"

//canPurchase
var CanPurchase string = "SELECT (quantity >= $1) FROM inventory WHERE product_id = $2"
