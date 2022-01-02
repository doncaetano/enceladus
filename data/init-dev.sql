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
    email  VARCHAR(50) UNIQUE,
    first_name VARCHAR(50),
    last_name  VARCHAR(50),
  	is_active  BOOL NOT NULL DEFAULT TRUE,
  	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  	updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TRIGGER update_user_timestamp
BEFORE UPDATE ON "user"
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();

INSERT INTO "user" (first_name, last_name, email) VALUES
  ('Tony', 'Stark', 'tonystark@email.com'),
  ('Hulk', 'Smash', 'hulksmash@email.com');


