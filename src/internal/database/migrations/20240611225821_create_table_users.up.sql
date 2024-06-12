CREATE TABLE IF NOT EXISTS users(
                                    id serial PRIMARY KEY,
                                    first_name VARCHAR (50) UNIQUE NOT NULL,
                                    last_name VARCHAR (50) UNIQUE NOT NULL,
                                    phone_number VARCHAR (10) UNIQUE NOT NULL,
                                    password VARCHAR (50) NOT NULL,
                                    email VARCHAR (300) UNIQUE NOT NULL,
                                    created_at TIMESTAMP NOT NULl,
                                    updated_at TIMESTAMP NOT NULl DEFAULT NOW()

);