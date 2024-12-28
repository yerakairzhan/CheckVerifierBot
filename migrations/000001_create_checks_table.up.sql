CREATE TABLE users (
      id SERIAL PRIMARY KEY,
      user_id VARCHAR(20) NOT NULL,
      username VARCHAR(20) NOT NULL,
      purchased Boolean default false,
      language_code VARCHAR(10) DEFAULT 'en' not null ,
      UNIQUE (user_id, username)
);
