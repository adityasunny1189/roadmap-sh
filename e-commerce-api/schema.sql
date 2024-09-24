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
    category VARCHAR(100) NOT NULL,
    price FLOAT NOT NULL,
    image_url VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS inventory (
    product_id CHAR(30) NOT NULL, 
    quantity INT NOT NULL
);

CREATE TABLE IF NOT EXISTS carts (
    id CHAR(30) PRIMARY KEY,
    user_id CHAR(30) NOT NULL,
    cart_state ENUM('CART_CREATED', 'CART_COMPLETED', 'CART_CANCELED') DEFAULT('CART_CREATED'),
    cart_amount FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW()),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS cart_items (
    id CHAR(30) PRIMARY KEY,
    cart_id CHAR(30) NOT NULL,
    product_id CHAR(30) NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW()),
    FOREIGN KEY (cart_id) REFERENCES carts(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS orders (
    id CHAR(30) PRIMARY KEY,
    user_id CHAR(30) NOT NULL, 
    cart_id CHAR(30) NOT NULL,
    total_price FLOAT NOT NULL,
    order_state ENUM('ORDER_CREATED', 'ORDER_COMPLETED', 'ORDER_CANCELLED') DEFAULT('ORDER_CREATED'),
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW()),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (cart_id) REFERENCES carts(id)
);

CREATE TABLE IF NOT EXISTS payments (
    id CHAR(30) PRIMARY KEY,
    order_id CHAR(30) NOT NULL,
    payment_method ENUM('STRIPE', 'RAZORPAY', 'PAYPAL') NOT NULL, -- For now we only have stripe
    payment_state ENUM('PAYMENT_PENDING', 'PAYMENT_SUCCESS', 'PAYMENT_FAILED', 'PAYMENT_CANCELED') DEFAULT('PAYMENT_PENDING'),
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW()),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

