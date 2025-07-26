package reader

import (
	"net/http"
)

func main() {
	http.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
	})
}
