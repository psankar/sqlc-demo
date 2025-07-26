BEGIN;

-- 1. Create the new `authors` table with unique email addresses
CREATE TABLE authors (
    author_id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE
);

-- 2. Add a temporary column to store `author_id` in `posts`
ALTER TABLE posts ADD COLUMN author_id INTEGER;

-- 3. Insert unique emails into `authors` and update `posts` with corresponding `author_id`
INSERT INTO authors (email)
SELECT DISTINCT author_email FROM posts;

-- 4. Update posts.author_id based on author_email
UPDATE posts
SET author_id = authors.author_id
FROM authors
WHERE posts.author_email = authors.email;

-- 5. Add NOT NULL constraint and FOREIGN KEY constraint
ALTER TABLE posts
    ALTER COLUMN author_id SET NOT NULL,
    ADD CONSTRAINT posts_author_id_fkey FOREIGN KEY (author_id) REFERENCES authors(author_id);

-- 6. Drop the old `author_email` column from `posts`
ALTER TABLE posts DROP COLUMN author_email;

COMMIT;