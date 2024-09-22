CREATE DATABASE my_commerce;

USE my_commerce;

CREATE TABLE IF NOT EXISTS users (
    id CHAR(30) PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    full_name VARCHAR(200) GENERATED ALWAYS AS (CONCAT(first_name, ' ', last_name)),
    user_email VARCHAR(100) NOT NULL UNIQUE,
    user_password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW())
);

CREATE TABLE IF NOT EXISTS products (
    id CHAR(30) PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    product_description TEXT NOT NULL,
    price FLOAT NOT NULL,
    image_url VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS inventory (
    product_id CHAR(30) NOT NULL, 
    quantity INT NOT NULL
);

CREATE TABLE IF NOT EXISTS carts (
    id CHAR(30) PRIMARY KEY
    user_id CHAR(30) NOT NULL FOREIGN KEY REFERENCES users(id),
    cart_state ENUM('ACTIVE', 'INACTIVE') DEFAULT('ACTIVE'),
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW())
);

CREATE TABLE IF NOT EXISTS cart_items (
    id CHAR(30) PRIMARY KEY,
    cart_id CHAR(30) NOT NULL FOREIGN KEY REFERENCES carts(id),
    product_id CHAR(30) NOT NULL FOREIGN KEY REFERENCES products(id),
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW())
);

CREATE TABLE IF NOT EXISTS orders (
    id CHAR(30) PRIMARY KEY,
    user_id CHAR(30) NOT NULL FOREIGN KEY REFERENCES users(id), 
    cart_id CHAR(30) NOT NULL FOREIGN KEY REFERENCES carts(id),
    total_price FLOAT NOT NULL,
    order_state ENUM('CREATED', 'COMPLETED', 'CANCELLED') DEFAULT('CREATED'),
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW())
);

CREATE TABLE IF NOT EXISTS payments (
    id CHAR(30) PRIMARY KEY,
    order_id INT NOT NULL FOREIGN KEY REFERENCES orders(id),
    payment_method ENUM('STRIPE', 'RAZORPAY', 'PHONEPAY', 'PAYPAL') NOT NULL, -- For now we only have stripe
    payment_status ENUM('PENDING', 'COMPLETED', 'FAILED') DEFAULT('PENDING'),
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW())
);


INSERT INTO products 
    (p_name, p_description, price) 
VALUES
    ("Tomato", "Round red vegetable", 10),
    ("Potato", "Brown vegetable", 12);

SELECT * FROM products p 
JOIN 
inventory i ON p.id = i.product_id 
WHERE 
i.quantity > 0 AND p.price > 100;



