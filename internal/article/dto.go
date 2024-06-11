package article

import (
	"time"
)

type ArticleInput struct {
	Title             string           `json:"title"`
	Description       string           `json:"description"`
	ThumbnailURL      string           `json:"thumbnail_url"`
	WasteCategories   []string         `json:"waste_categories"`
	ContentCategories []string         `json:"content_categories"`
	Sections          []ArticleSection `json:"sections"`
}

type ArticleSectionInput struct {
	ArticleID   string `json:"article_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type AdminDetail struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type ArticleDetail struct {
	ID           string      `json:"id"`
	AuthorID     AdminDetail `json:"author"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	ThumbnailURL string      `json:"thumbnail_url"`
	CreatedAt    time.Time   `json:"created_at"`

	WasteCategories   []WasteCategory  `json:"waste_categories"`
	ContentCategories []VideoCategory  `json:"content_categories"`
	Sections          []ArticleSection `json:"sections"`
}

type ArticleResponsePagination struct {
	Total    int64           `json:"total"`
	Page     uint            `json:"page"`
	Limit    uint            `json:"limit"`
	Articles []ArticleDetail `json:"articles"`
}