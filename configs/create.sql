CREATE TABLE CartItems
(
    Id       serial PRIMARY KEY NOT NULL,
    CartId   integer NOT NULL,
    Product  VARCHAR(50) NOT NULL,
    Quantity integer NOT NULL
);