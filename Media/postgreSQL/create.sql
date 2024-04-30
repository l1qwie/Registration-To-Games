CREATE SEQUENCE mediarepository_id_seq
    INCREMENT 1
    START 1
;

CREATE TABLE schedule
(
    gameid integer NOT NULL DEFAULT 0 PRIMARY KEY,
    sport text ,
    date integer,
    "time" integer,
    status integer DEFAULT 0
);

CREATE TABLE mediarepository
(   
    id integer NOT NULL DEFAULT 0 PRIMARY KEY,
    fileid text DEFAULT '',
    type text DEFAULT '',
    userid integer DEFAULT 0,
    gameid integer DEFAULT 0,
    counter integer DEFAULT 0,
    status integer DEFAULT 0
);