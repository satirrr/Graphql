CREATE TABLE users (
    id SERIAL PRIMARY KEY ,
    firstName TEXT NOT NULL ,
    lastName TEXT NOT NULL
);

CREATE TABLE todos (
    id SERIAL PRIMARY KEY ,
    userId SERIAL NOT NULL REFERENCES users(id) ,
    task TEXT NOT NULL ,
    done BOOLEAN NOT NULL
);