INSERT INTO rp_product  (name, description) VALUES ('Monitor A', 'This is Monitor A');

INSERT INTO rp_stock_item  (product_id) VALUES (1);
INSERT INTO rp_stock_item  (product_id) VALUES (1);
INSERT INTO rp_stock_item  (product_id) VALUES (1);
INSERT INTO rp_stock_item  (product_id) VALUES (1);

INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (1, '02-10-2023', '10-10-2023');
INSERT INTO rp_order (stock_item_id, start_date, end_date) VALUES (2, '05-10-2023', '15-10-2023');