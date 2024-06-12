CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO roles (name, created_at) VALUES ('admin', NOW());
INSERT INTO roles (name, created_at) VALUES ('user', NOW());
