CREATE TYPE categories as ENUM('Clothing', 'Accessories', 'Flower', 'Beverages');
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category categories,
    imageUrl VARCHAR(255),
    notes   VARCHAR(100),
    price   INTEGER,
    location VARCHAR(255),
    isAvailable BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);