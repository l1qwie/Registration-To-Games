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
    payment varchar,
    statuspayment integer DEFAULT 0,
    status integer DEFAULT 0
);

CREATE TABLE schedule
(
    gameid integer NOT NULL DEFAULT 0 PRIMARY KEY,
    sport varchar ,
    date integer,
    "time" integer,
    seats integer,
    link varchar,
    address varchar ,
    price integer,
    currency varchar ,
    status integer DEFAULT 0
);