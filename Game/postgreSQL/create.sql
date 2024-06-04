CREATE SEQUENCE gameswithusers_id_seq
    INCREMENT 1
    START 1
;

CREATE TABLE Schedule (
    gameId integer NOT NULL PRIMARY KEY DEFAULT 0,
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
