CREATE TABLE public."user" (
    id VARCHAR NOT NULL,
    document VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    age INTEGER NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

ALTER TABLE public."user"
    ADD CONSTRAINT user_pk PRIMARY KEY (id);

ALTER TABLE public."user"
    ADD CONSTRAINT user_unique_document UNIQUE (document);

ALTER TABLE public."user"
    ADD CONSTRAINT user_unique_email UNIQUE (email);

ALTER TABLE public."user"
    ADD CONSTRAINT user_unique_password UNIQUE (password);
