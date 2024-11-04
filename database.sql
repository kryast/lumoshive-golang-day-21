CREATE TABLE users (
    ID SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    pass VARCHAR(100)
);

CREATE TABLE tasks (
    ID SERIAL PRIMARY KEY,
    description VARCHAR(256)
);