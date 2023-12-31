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

INSERT INTO rp_product  (name, description) VALUES ('Monitor A', 'This is Monitor A');

INSERT INTO rp_stock_item  (product_id) VALUES (1);
INSERT INTO rp_stock_item  (product_id) VALUES (1);
INSERT INTO rp_stock_item  (product_id) VALUES (1);
INSERT INTO rp_stock_item  (product_id) VALUES (1);

INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (1, '02-10-2023', '10-10-2023');
INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (2, '02-10-2023', '10-10-2023');
INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (3, '02-10-2023', '10-10-2023');
INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (4, '02-10-2023', '10-10-2023');
INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (2, '12-10-2023', '15-10-2023');


INSERT INTO rp_product  (name, description) VALUES ('Monitor B', 'This is Monitor B');

INSERT INTO rp_stock_item  (product_id) VALUES (2);
INSERT INTO rp_stock_item  (product_id) VALUES (2);
INSERT INTO rp_stock_item  (product_id) VALUES (2);
INSERT INTO rp_stock_item  (product_id) VALUES (2);

INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (5, '02-11-2023', '10-11-2023');
INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (6, '02-11-2023', '10-11-2023');
INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (7, '02-11-2023', '10-11-2023');
INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (8, '02-11-2023', '10-11-2023');
INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (6, '12-11-2023', '15-11-2023');