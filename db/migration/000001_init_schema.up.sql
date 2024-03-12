CREATE TABLE Account (
  id uuid PRIMARY KEY,
  role VARCHAR NOT NULL,
  username VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  is_email_verified bool NOT NULL DEFAULT false,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Added unique constraint to the username column
ALTER TABLE Account ADD CONSTRAINT unique_username UNIQUE (username);

CREATE TABLE Form (
  id uuid PRIMARY KEY,
  viewer_name VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  phone VARCHAR(20) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Post (
  id uuid PRIMARY KEY,
  title VARCHAR NOT NULL,
  article TEXT NOT NULL,
  picture BYTEA,
  -- username VARCHAR NOT NULL, 
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
);
  -- FOREIGN KEY (username) REFERENCES Account (username)

CREATE TABLE Session (
  id uuid PRIMARY KEY,
  username VARCHAR NOT NULL,
  role VARCHAR(5) NOT NULL,
  expires_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (username) REFERENCES Account(username)
);

CREATE TABLE Verify_email (
  username VARCHAR NOT NULL PRIMARY KEY,
  secret_code VARCHAR NOT NULL,
  is_used bool NOT NULL DEFAULT false,
  expires_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '15 minutes',
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Added foreign key constraint to Verify_email table
ALTER TABLE Verify_email ADD FOREIGN KEY (username) REFERENCES Account (username);

-- Added foreign key constraint to Post table
-- ALTER TABLE Post ADD FOREIGN KEY (username) REFERENCES Account (username);