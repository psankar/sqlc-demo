-- name: CreatePostWithAuthorEmail :one
WITH ins_author AS (
  INSERT INTO authors (email)
  VALUES ($1)
  ON CONFLICT (email) DO NOTHING
  RETURNING author_id
),
sel_author AS (
  SELECT author_id FROM ins_author
  UNION
  SELECT author_id FROM authors WHERE email = $1
),
ins_post AS (
  INSERT INTO posts (author_id, post)
  SELECT author_id, $2 FROM sel_author
  RETURNING *
)
SELECT * FROM ins_post;

-- name: GetPost :one
SELECT * FROM posts WHERE post_id = $1;
