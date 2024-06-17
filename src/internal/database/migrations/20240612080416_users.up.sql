CREATE TABLE IF NOT EXISTS users(
                                    id serial PRIMARY KEY,
                                    first_name VARCHAR (50)  NOT NULL,
                                    last_name VARCHAR (50)  NOT NULL,
                                    phone_number VARCHAR (10) NULL,
                                    password VARCHAR (50) NOT NULL,
                                    email VARCHAR (300) UNIQUE NOT NULL,
                                    role_id INT NOT NULL,
                                    created_at TIMESTAMP NOT NULl,
                                    updated_at TIMESTAMP NOT NULl DEFAULT NOW()

);

ALTER TABLE users ADD CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE ON UPDATE CASCADE;


INSERT INTO users  (first_name, last_name , phone_number, password, email, role_id, created_at)
VALUES ('admin', 'admin', '1234567890', 'admin', 'admin@example.com', 1, NOW());
