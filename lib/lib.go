package lib

type AddPostRequest struct {
	Post        string `json:"post"`
	AuthorEmail string `json:"author_email"`
}

type AddPostResponse struct {
	PostID int32 `json:"post_id"`
}
