CREATE DATABASE pro_cat_shop_cart_service;

CREATE TABLE Product_master (
Product_Id INT PRIMARY KEY,
Name VARCHAR,
Specification JSON,
SKU VARCHAR,
Category_Id INT,
Price FLOAT,
CONSTRAINT fk_Product_master_Category_master FOREIGN KEY(Category_Id) REFERENCES Category_master(Category_Id)
);

CREATE TABLE Category_master (
Category_Id INT PRIMARY KEY,
Name VARCHAR );

CREATE TABLE Inventory (
Product_Id INT,
Quantity INT,
CONSTRAINT fk_Inventory_Product_master FOREIGN KEY(Product_Id) REFERENCES Product_master(Product_Id) );

CREATE TABLE Cart_reference (
Reference_Id VARCHAR PRIMARY KEY,
Created_At timestamp without time zone
);

CREATE TABLE Cart_items (
Reference_Id VARCHAR,
Product_Id INT,
Quantity INT,
PRIMARY KEY(Reference_Id,Product_Id),
CONSTRAINT fk_Cart_items_Product_master FOREIGN KEY(Product_Id) REFERENCES Product_master(Product_Id) 
);