ALTER TABLE users
add username VARCHAR(100) NOT NULL;

ALTER TABLE users
ADD CONSTRAINT UNIQUE unique_username (username);