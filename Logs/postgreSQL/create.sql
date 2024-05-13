CREATE SEQUENCE action_id_seq
    INCREMENT 1
    START 1
;

CREATE TABLE Actions
(
    id integer NOT NULL DEFAULT 0 PRIMARY KEY,
    userId bigint NOT NULL DEFAULT 0,
    act text DEFAULT '',
    eventTime TIMESTAMP
);