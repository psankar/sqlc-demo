package writer

import (
	"encoding/json"
	"net/http"

	"github.com/psankar/sqlc-demo/lib"
)

func main() {
	http.HandleFunc("/add-post/", func(w http.ResponseWriter, r *http.Request) {
		var addPostRequest lib.AddPostRequest
		err := json.NewDecoder(r.Body).Decode(&addPostRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	})
}
