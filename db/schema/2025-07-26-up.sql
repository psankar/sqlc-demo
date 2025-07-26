CREATE TABLE posts (
    post_id SERIAL PRIMARY KEY,
    author_email TEXT NOT NULL,
    post TEXT
);