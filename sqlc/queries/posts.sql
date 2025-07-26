-- name: CreatePost :one
INSERT INTO posts (author_email, post) VALUES ($1, $2) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts WHERE post_id = $1;
