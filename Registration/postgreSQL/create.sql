CREATE SEQUENCE gameswithusers_id_seq
    INCREMENT 1
    START 1
;

CREATE TABLE gameswithusers
(   
    id integer NOT NULL DEFAULT 0 PRIMARY KEY,
    userid bigint ,
    gameid integer,
    seats integer,
    payment text,
    statuspayment integer DEFAULT 0,
    status integer DEFAULT 0
);

CREATE TABLE schedule
(
    gameid integer NOT NULL DEFAULT 0 PRIMARY KEY,
    sport text ,
    date integer,
    "time" integer,
    seats integer,
    latitude double precision,
    longitude double precision,
    address text ,
    price integer,
    currency text ,
    status integer DEFAULT 0
);