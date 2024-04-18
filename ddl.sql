CREATE DATABASE p2_fp;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fullname VARCHAR(200) NOT NULL,
    email VARCHAR(200) NOT NULL,
    password VARCHAR(200) NOT NULL,
    deposit_amount DECIMAL(10, 2),
    role ENUM('admin', 'member') NOT NULL
);

CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tittle VARCHAR(200) NOT NULL,
    category VARCHAR(200) NOT NULL,
    stock INT,
    author_name VARCHAR(200) NOT NULL,
    rent_cost DECIMAL(10, 2)
);

CREATE TABLE rental_cart (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    book_id INT,
    quantity INT NOT NULL,
    total_price DECIMAL(10, 2),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_id) REFERENCES books(id)
);

CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    rental_id INT,
    rent_date DATE NOT NULL,
    due_date DATE NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (rental_id) REFERENCES rental_cart(id)
);

CREATE TABLE payments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_id INT,
    total_amount DECIMAL,
    payment_date DATE NOT NULL,
    status ENUM ('canceled', 'pending', 'complete'),
    FOREIGN KEY (transaction_id) REFERENCES transactions(id)
);
