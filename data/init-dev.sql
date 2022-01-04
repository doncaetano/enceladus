CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS "user"
(
    id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    email  VARCHAR(50) NOT NULL UNIQUE,
    password  VARCHAR(60) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name  VARCHAR(50) NOT NULL,
  	is_active  BOOL NOT NULL DEFAULT TRUE,
  	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

DROP TRIGGER IF EXISTS update_user_timestamp ON "user";
CREATE TRIGGER update_user_timestamp
BEFORE UPDATE ON "user"
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();

CREATE TABLE IF NOT EXISTS refresh_token
(
    id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id  uuid NOT NULL,
  	is_active  BOOL NOT NULL DEFAULT TRUE,
  	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  	FOREIGN KEY (user_id) REFERENCES "user"(id)
);

DROP TRIGGER IF EXISTS update_user_timestamp ON refresh_token;
CREATE TRIGGER update_user_timestamp
BEFORE UPDATE ON refresh_token
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();

CREATE TABLE IF NOT EXISTS access_token
(
    id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id  uuid NOT NULL,
  	refresh_token_id uuid NOT NULL,
  	expires_at TIMESTAMP WITH TIME ZONE,
  	is_active  BOOL NOT NULL DEFAULT TRUE,
  	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  	FOREIGN KEY (user_id) REFERENCES "user"(id),
  	FOREIGN KEY (refresh_token_id) REFERENCES refresh_token(id)
);

DROP TRIGGER IF EXISTS update_user_timestamp ON access_token;
CREATE TRIGGER update_user_timestamp
BEFORE UPDATE ON access_token
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();

INSERT INTO "user" (first_name, last_name, email, password) VALUES
  ('Tony', 'Stark', 'tonystark@email.com', '$2a$10$98DlxWcup8dTCnNWRMJs4eHi0yhtUCedBJ6RF205af246M/rBxQ8C'),
  ('Hulk', 'Smash', 'hulksmash@email.com', '$2a$10$98DlxWcup8dTCnNWRMJs4eHi0yhtUCedBJ6RF205af246M/rBxQ8C');


