CREATE TABLE users (
      id SERIAL PRIMARY KEY,
      user_id VARCHAR(20) NOT NULL,
      username VARCHAR(20) NOT NULL,
      purchased Boolean default false,
      language_code VARCHAR(10) DEFAULT 'en' not null ,
      chosen_package varchar(255) default 'no' not null,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      UNIQUE (user_id, username)
);
