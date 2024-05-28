CREATE TABLE CartItems
(
    Id       serial PRIMARY KEY NOT NULL,
    CartId   integer REFERENCES Carts(Id) NOT NULL,
    Product  VARCHAR(50) NOT NULL UNIQUE,
    Quantity integer NOT NULL
);

CREATE TABLE Carts
(
    Id serial PRIMARY KEY NOT NULL
);