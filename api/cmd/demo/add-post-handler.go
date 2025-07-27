package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/psankar/sqlc-demo/sqlc/db"
)

type AddPostRequest struct {
	Post        string `json:"post"`
	AuthorEmail string `json:"author_email"`
}

type AddPostResponse struct {
	PostID int32 `json:"post_id"`
}

func addPostHandler(w http.ResponseWriter, r *http.Request) {
	var addPostRequest AddPostRequest
	err := json.NewDecoder(r.Body).Decode(&addPostRequest)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := queries.CreatePostWithAuthorEmail(
		r.Context(),
		db.CreatePostWithAuthorEmailParams{
			Email: addPostRequest.AuthorEmail,
			Post:  pgtype.Text{String: addPostRequest.Post, Valid: true},
		})
	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.NewEncoder(w).Encode(AddPostResponse{
		PostID: post.PostID,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}
