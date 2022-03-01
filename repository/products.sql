CREATE TABLE products (
    id serial,
    sku varchar(10),
    name varchar(100),
    price numeric DEFAULT 9.99,
    qty int
);