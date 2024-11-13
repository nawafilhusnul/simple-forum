ALTER TABLE users ADD COLUMN user_name VARCHAR(255) NOT NULL;
ALTER TABLE users ADD CONSTRAINT user_name_unique UNIQUE (user_name);