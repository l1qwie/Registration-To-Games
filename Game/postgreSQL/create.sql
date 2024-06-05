CREATE TABLE Schedule (
    gameId serial PRIMARY KEY,
    sport text DEFAULT '',
    date integer DEFAULT 0,
    time integer DEFAULT 0,
    seats integer DEFAULT 0,
    price integer DEFAULT 0,
    currency text DEFAULT '',
    link text DEFAULT '',
    address text DEFAULT '',
    status integer DEFAULT 0
);