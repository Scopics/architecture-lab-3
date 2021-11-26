-- Database: restaurant

-- DROP DATABASE "restaurant";
-- CREATE DATABASE "restaurant";

-- Create tables

CREATE TABLE IF NOT EXISTS "menu" (
  "id" serial NOT NULL,
  "name" varchar(50) NOT NULL UNIQUE,
  "price" integer NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "orders" (
  "id" serial NOT NULL,
  "table" integer NOT NULL,
  "date" timestamp DEFAULT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "order_details" (
  "order_id" integer NOT NULL,
  "meal_id" integer NOT NULL,
  "quantity" integer DEFAULT 1 NOT NULL,
  PRIMARY KEY ("order_id", "meal_id"),
  CONSTRAINT order_details_meal_id_menu_id_foreign FOREIGN KEY ("meal_id") REFERENCES menu ("id"),
  CONSTRAINT order_details_order_id_orders_id_foreign FOREIGN KEY ("order_id") REFERENCES orders ("id")
);

CREATE TABLE IF NOT EXISTS "constants" (
  "id" serial NOT NULL,
  "name" varchar(50) NOT NULL,
  "value" decimal NOT NULL,
  PRIMARY KEY ("id")
);


-- Insert data

INSERT INTO "constants" ("name", "value") VALUES ('tax', 0.08);
INSERT INTO "constants" ("name", "value") VALUES ('tip', 0.15);

INSERT INTO "menu" ("name", "price") VALUES ('Fried rice with vegetables', 12);
INSERT INTO "menu" ("name", "price") VALUES ('Spicy pasta', 9);
INSERT INTO "menu" ("name", "price") VALUES ('Chicken salad', 13);
INSERT INTO "menu" ("name", "price") VALUES ('Beefsteak', 15);
INSERT INTO "menu" ("name", "price") VALUES ('Chocolate cake', 10);
INSERT INTO "menu" ("name", "price") VALUES ('Juice', 6);
INSERT INTO "menu" ("name", "price") VALUES ('Coffee', 4);

INSERT INTO "orders" ("table", "date") VALUES (1, '2021-11-21 17:04');
INSERT INTO "orders" ("table", "date") VALUES (2, '2021-11-21 16:44');
INSERT INTO "orders" ("table", "date") VALUES (2, '2021-11-21 18:14');
INSERT INTO "orders" ("table", "date") VALUES (4, '2021-11-21 17:34');

INSERT INTO "order_details" ("order_id", "meal_id", "quantity") VALUES (1, 1, 1);
INSERT INTO "order_details" ("order_id", "meal_id", "quantity") VALUES (1, 3, 1);

INSERT INTO "order_details" ("order_id", "meal_id", "quantity") VALUES (2, 7, 2);
INSERT INTO "order_details" ("order_id", "meal_id", "quantity") VALUES (2, 5, 2);

INSERT INTO "order_details" ("order_id", "meal_id") VALUES (3, 2);

INSERT INTO "order_details" ("order_id", "meal_id", "quantity") VALUES (4, 2, 1);
INSERT INTO "order_details" ("order_id", "meal_id", "quantity") VALUES (4, 4, 1);
INSERT INTO "order_details" ("order_id", "meal_id", "quantity") VALUES (4, 6, 2);


-- Create functions

CREATE OR REPLACE FUNCTION get_constant(const_name TEXT) RETURNS DECIMAL
LANGUAGE SQL   
AS   
$$  
    SELECT value FROM constants
	WHERE name = const_name;
$$; 

CREATE OR REPLACE FUNCTION get_total_price(ord_id INT) 
RETURNS TABLE (order_id INT, total_price BIGINT, total_price_without_tax NUMERIC, recommended_tips NUMERIC)
LANGUAGE SQL   
AS   
$$  
    SELECT od.order_id, 
	SUM(m.price * od.quantity) AS total_price, 
	SUM(m.price * od.quantity) * (1 - get_constant('tax')) AS total_price_without_tax,
	SUM(m.price * od.quantity) * get_constant('tip') AS recommended_tips
FROM order_details od
JOIN menu m ON m.id = od.meal_id
JOIN orders o ON o.id = od.order_id
GROUP BY od.order_id
HAVING od.order_id = ord_id;  
$$; 

