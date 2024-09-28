CREATE DATABASE my_commerce;

USE my_commerce;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    full_name VARCHAR(200) GENERATED ALWAYS AS (CONCAT(first_name, ' ', last_name)),
    user_email VARCHAR(100) NOT NULL UNIQUE,
    user_password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW())
);

CREATE TABLE IF NOT EXISTS categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    category_name VARCHAR(100) NOT NULL,
    category_description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_name VARCHAR(100) NOT NULL,
    product_description TEXT NOT NULL,
    price FLOAT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS inventory (
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT NOT NULL, 
    quantity INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS carts (
    id CHAR(30) PRIMARY KEY,
    user_id INT NOT NULL,
    cart_state ENUM('CART_CREATED', 'CART_COMPLETED', 'CART_CANCELED') DEFAULT('CART_CREATED'),
    cart_amount FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW()),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cart_items (
    id CHAR(30) PRIMARY KEY,
    cart_id CHAR(30) NOT NULL,
    product_id CHAR(30) NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW()),
    FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id INT PRIMARY KEY,
    user_id INT NOT NULL, 
    cart_id CHAR(30) NOT NULL,
    total_price FLOAT NOT NULL,
    order_state ENUM('ORDER_CREATED', 'ORDER_COMPLETED', 'ORDER_CANCELLED') DEFAULT('ORDER_CREATED'),
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW()),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS payments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT NOT NULL,
    payment_method ENUM('STRIPE', 'RAZORPAY', 'PAYPAL') NOT NULL, -- For now we only have stripe
    payment_state ENUM('PAYMENT_PENDING', 'PAYMENT_SUCCESS', 'PAYMENT_FAILED', 'PAYMENT_CANCELED') DEFAULT('PAYMENT_PENDING'),
    created_at TIMESTAMP DEFAULT(NOW()),
    updated_at TIMESTAMP DEFAULT(NOW()),
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
);

