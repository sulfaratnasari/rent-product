CREATE TABLE rp_product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE rp_stock_item (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES rp_product(id)
);

CREATE TABLE rp_order (
    id SERIAL PRIMARY KEY,
    stock_item_id INT,
    start_date DATE,
    end_date DATE,
    FOREIGN KEY (stock_item_id) REFERENCES rp_stock_item(id)
);