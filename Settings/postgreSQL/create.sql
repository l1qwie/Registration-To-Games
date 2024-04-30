
CREATE TABLE  users
(
    userid bigint NOT NULL DEFAULT 0 PRIMARY KEY,
    language text DEFAULT '',
    customlanguage boolean DEFAULT false
);

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
    price integer,
    currency text ,
    status integer DEFAULT 0
);