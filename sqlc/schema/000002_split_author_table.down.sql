BEGIN;

-- 1. Add back the author_email column to `posts`
ALTER TABLE posts ADD COLUMN author_email TEXT;

-- 2. Populate author_email from authors table using author_id
UPDATE posts
SET author_email = authors.email
FROM authors
WHERE posts.author_id = authors.author_id;

-- 3. Drop the foreign key constraint and author_id column
ALTER TABLE posts DROP CONSTRAINT posts_author_id_fkey;
ALTER TABLE posts DROP COLUMN author_id;

-- 4. Drop the `authors` table
DROP TABLE authors;

COMMIT;
