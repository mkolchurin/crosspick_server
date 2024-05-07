drop table maps; 
CREATE TABLE maps (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    mode int NOT NULL,
    icon TEXT NOT NULL,
    description TEXT,
    author VARCHAR(255),
    created_at BIGINT
);