package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	postID := r.PathValue("post_id")
	if postID == "" {
		http.Error(w, "Post ID is required", http.StatusBadRequest)
		return
	}

	// Convert postID to int32
	var id int32
	_, err := fmt.Sscanf(postID, "%d", &id)
	if err != nil {
		log.Println("Error parsing post ID:", err)
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := queries.GetPost(r.Context(), id)
	if err != nil {
		log.Println("Error fetching post:", err)
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
