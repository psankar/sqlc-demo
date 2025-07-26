package writer

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/psankar/sqlc-demo/lib"
	"github.com/psankar/sqlc-demo/sqlc/db"
)

var queries *db.Queries

func init() {
	conn, err := pgx.Connect(context.Background(), "")
	if err != nil {
		log.Fatal(err)
		return
	}

	queries = db.New(conn)
}

func main() {
	http.HandleFunc("/add-post/", func(w http.ResponseWriter, r *http.Request) {
		var addPostRequest lib.AddPostRequest
		err := json.NewDecoder(r.Body).Decode(&addPostRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		post, err := queries.CreatePost(
			r.Context(),
			db.CreatePostParams{
				AuthorEmail: addPostRequest.AuthorEmail,
				Post:        addPostRequest.Post,
			})
		if err != nil {
			log.Fatal(err)
			return
		}

		err = json.NewEncoder(w).Encode(lib.AddPostResponse{
			PostID: post.PostID,
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	})
}
