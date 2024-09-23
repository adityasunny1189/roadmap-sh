SELECT * FROM products p 
JOIN 
inventory i ON p.id = i.product_id 
WHERE 
i.quantity > 0 AND p.price > 10000;

-- Out of stock product 
SELECT * FROM products p 
JOIN 
inventory i ON p.id = i.product_id 
WHERE 
i.quantity = 0;

-- Create a product view of product in stock
CREATE VIEW product_in_stock_view AS
SELECT * FROM products p
JOIN
inventory i ON p.id = i.product_id
WHERE
i.quantity > 0;

-- out of stock product view
CREATE VIEW out_of_stock_product_view AS 
SELECT * FROM products p
JOIN
inventory i ON p.id = i.product_id
WHERE
i.quantity = 0;

DESC out_of_stock_product_view;

SELECT * FROM product_in_stock_view;

SELECT * FROM out_of_stock_product_view;

SELECT * FROM product_in_stock_view WHERE category = "mobile";

SELECT * FROM out_of_stock_product_view WHERE category = "mobile";

