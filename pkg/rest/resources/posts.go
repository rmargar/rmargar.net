package resources

import (
	"time"

	"github.com/rmargar/website/pkg/domain"
)

type PostPayloadJSONApi struct {
	Title   string   `json:"title" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Tags    []string `json:"tags"`
}

type PostJSONApi struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Tags     []string  `json:"tags"`
	Author   string    `json:"author"`
	Added    time.Time `json:"added"`
	Modified time.Time `json:"modified"`
}

type PostCreatedJSONApi struct {
	Message string      `json:"msg"`
	Data    PostJSONApi `json:"data"`
}

func BuildCreatedResponse(post *domain.Post) *PostCreatedJSONApi {
	return &PostCreatedJSONApi{
		Message: "Created",
		Data: PostJSONApi{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			Author:   post.Author,
			Tags:     post.Tags,
			Added:    post.Added,
			Modified: post.Modified,
		},
	}
}
