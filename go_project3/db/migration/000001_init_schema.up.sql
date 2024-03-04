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
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE UserRole AS ENUM ('guest', 'user', 'admin');

CREATE TABLE Account (
  id uuid PRIMARY KEY,
  role UserRole NOT NULL,
  username VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  is_email_verified bool NOT NULL DEFAULT false,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Session (
  id uuid PRIMARY KEY,
  username VARCHAR NOT NULL,
  -- token VARCHAR(500) NOT NULL,
  role VARCHAR(5) NOT NULL,
  expires_at TIMESTAMPTZ,
  -- DEFAULT CURRENT_TIMESTAMP + INTERVAL '1 hour',
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Verify_email (
  -- id uuid PRIMARY KEY,
  username varchar NOT NULL PRIMARY KEY,
  secret_code varchar NOT NULL,
  is_used bool NOT NULL DEFAULT false,
  expires_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '15 minutes',
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE Verify_email ADD FOREIGN KEY (username) REFERENCES Account (username);

-- ALTER TABLE Account ADD COLUMN is_email_verified bool NOT NULL DEFAULT false;

ALTER TABLE Session ADD FOREIGN KEY (username) REFERENCES Account(username);

ALTER TABLE Post ADD FOREIGN KEY (username) REFERENCES Account (username);

