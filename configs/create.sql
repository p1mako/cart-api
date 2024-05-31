CREATE TABLE CartItems
(
    id       serial PRIMARY KEY NOT NULL,
    cart_id   integer REFERENCES Carts(id) NOT NULL,
    product  VARCHAR(50) NOT NULL,
    quantity integer NOT NULL,
    UNIQUE (cart_id, product)
);

CREATE TABLE Carts
(
    id serial PRIMARY KEY NOT NULL
);