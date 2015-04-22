CREATE TABLE users (
    id serial PRIMARY KEY,
    username character varying(255) UNIQUE NOT NULL,
    hash character varying(256) NOT NULL,
    salt character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE expense (
    id serial PRIMARY KEY,
    userid integer,
    amount double precision,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    FOREIGN KEY(userid) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE tags (
    id serial PRIMARY KEY,
    parent integer,
    userid integer,
    tag character varying(20),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    FOREIGN KEY(userid) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE expense_tags (    
    expenseid integer REFERENCES expense(id),
    tagid integer REFERENCES tags(id)
);
