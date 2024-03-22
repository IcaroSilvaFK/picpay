

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(40),
  email VARCHAR(40),
  password VARCHAR(255),
  identifier INTEGER,
  type  INTEGER
);

CREATE TABLE wallets (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  amount INTEGER,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  payer INTEGER,
  payee INTEGER,
  value INTEGER,
  FOREIGN KEY (payer) REFERENCES users(id),
  FOREIGN KEY (payee) REFERENCES users(id)
);