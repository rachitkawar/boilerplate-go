CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    role_id INT,
    created_at TIMESTAMP NOT NULL ,
    updated_at TIMESTAMP DEFAULT now()
);

ALTER TABLE permissions
    ADD CONSTRAINT permissions_name_unique UNIQUE (name),
    ADD CONSTRAINT permissions_role_id_fkey FOREIGN KEY (role_id) REFERENCES roles (id) ON UPDATE CASCADE ON DELETE CASCADE;

CREATE INDEX IF NOT EXISTS permissions_role_id_idx ON permissions (role_id);


-- Seeding Permissions --

INSERT INTO permissions (name, description, role_id, created_at)
VALUES ('create_user', 'Create a new user', 1, now()),
       ('update_user', 'Update an existing user', 1, now()),
       ('delete_user', 'Delete a user', 1, now() );

INSERT INTO permissions (name, description, role_id, created_at)
VALUES ('create_role', 'Create a new role', 1, now()),
       ('update_role', 'Update an existing role', 1, now()),
       ('delete_role', 'Delete a role', 1, now() );


INSERT INTO permissions (name, description, role_id, created_at)
VALUES ('create_permission', 'Create a new permission', 1, now()),
       ('update_permission', 'Update an existing permission', 1, now()),
       ('delete_permission', 'Delete a permission', 1, now() );