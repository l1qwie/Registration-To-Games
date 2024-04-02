CREATE SEQUENCE mediarepository_id_seq
    INCREMENT 1
    START 1
;

CREATE SEQUENCE gameswithusers_id_seq
    INCREMENT 1
    START 1
;

CREATE SEQUENCE gameid_seq
    INCREMENT 1
    START 1
;

CREATE TABLE users
(
    userid bigint NOT NULL DEFAULT 0 PRIMARY KEY,
    name text DEFAULT '',
    lastname text DEFAULT '',
    username text DEFAULT '',
    whence text DEFAULT '',
    language text DEFAULT '',
    useradmin boolean DEFAULT false,
    action text DEFAULT '',
    setupreg text DEFAULT '',
    actionwithmes text DEFAULT 'DEL',
    level integer DEFAULT 0,
    launchpoint integer DEFAULT 0, 
    seats integer DEFAULT 0,
    payment text  DEFAULT '',
    timeinterval text DEFAULT '',
    direction text DEFAULT '',
    mediagoupcounter integer DEFAULT 0,
    changeable text DEFAULT '',
    actgame text DEFAULT '',
    willchangeable text DEFAULT '',
    newpay text DEFAULT '',
    notifgameid integer DEFAULT 0,
    lasttimeuse timestamp without time zone,
    customlanguage boolean DEFAULT false,
    phonenumber bigint DEFAULT 0,
    status integer DEFAULT 1,
    gameid integer DEFAULT 0,
    mediagroupid text DEFAULT '',
    exmessageid integer DEFAULT 0,
    mlimit integer DEFAULT 7
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

CREATE TABLE gameswithusers
(   
    id integer NOT NULL DEFAULT 0 PRIMARY KEY,
    userid bigint ,
    gameid integer,
    seats integer,
    payment text,
    statuspayment integer DEFAULT 0,
    status integer DEFAULT 0
)