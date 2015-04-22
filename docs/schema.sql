CREATE TABLE users (
    id serial PRIMARY KEY,
    username varying(255) NOT NULL,
    hash varying(256) NOT NULL,
    salt character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE expense (
    id serial PRIMARY KEY,
    userid integer,
    amount double precision,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE tags (
    id serial PRIMARY KEY,
    parent integer,
    userid integer,
    tag varying(20),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE expense_tags (
    id serial PRIMARY KEY,
    expenseid integer,
    tagid integer,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);
