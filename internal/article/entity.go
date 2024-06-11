package article

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Article struct {
	ID           string `gorm:"primaryKey;type:varchar(7)"`
	Title        string `gorm:"type:varchar(255)"`
	Description  string `gorm:"type:text"`
	ThumbnailURL string `gorm:"type:varchar(255)"`
	AuthorID     string

	Categories []ArticleCategories
	Sections   []ArticleSection

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type WasteCategory struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(50);unique;not null"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type VideoCategory struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(50);unique;not null"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type ArticleCategories struct {
	ID                uint   `gorm:"primaryKey"`
	ArticleID         string `gorm:"type:varchar(7)"`
	WasteCategoryID   uint
	ContentCategoryID int

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type ArticleSection struct {
	ID          uint   `json:"-" gorm:"primaryKey"`
	ArticleID   string `json:"-" gorm:"type:varchar(7)"`
	Title       string `json:"title" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
	ImageURL    string `json:"image_url" gorm:"type:varchar(255)"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type ArticleRepository interface {
	// Article Repository
	Create(article Article) (*Article, error)
	FindByID(articleID string) (*Article, error)
	FindAll(page, limit uint) (*[]Article, int64, error)
	FindLastID() (string, error)
	FindByKeyword(keyword string) (*[]Article, error)
	FindByCategory(categoryName string, categoryType string) (*[]Article, error)
	Update(article Article) error
	Delete(articleID string) error

	// Category Repository
	FindCategories(articleID string) (*[]WasteCategory, *[]VideoCategory, error)
	FindCategoryByName(categoryName, categoryType string) (uint, error)

	// Article Section Repository
	CreateSection(section ArticleSection) error
	UpdateSection(section ArticleSection) error
	DeleteSection(sectionID uint) error
	DeleteAllSection(articleID string) error

	// Article Categories Repository
	CreateArticleCategory(categories ArticleCategories) error
	UpdateArticleCategory(categories ArticleCategories) error
	DeleteAllArticleCategory(articleID string) error
}

type ArticleUsecase interface {
	NewArticle(article ArticleInput, authorId string) (*ArticleDetail, error)
	GetArticleByID(articleID string) (*ArticleDetail, error)
	GetAllArticle(page, limit int) (*ArticleResponsePagination, error)
	GetArticleByKeyword(keyword string) (*[]ArticleDetail, error)
	GetArticleByCategory(categoryName string, categoryType string) (*[]ArticleDetail, error)
	Update(articleID string, article ArticleInput) error
	Delete(articleID string) error

	GetArticleDetail(article Article) *ArticleDetail
	GetDetailAuthor(authorID string) (*AdminDetail, error)

	// add usecase yang belum saya buat dibawah
}

type ArticleHandler interface {
	GetAllArticle(c echo.Context) error
	GetArticleByKeyword(c echo.Context) error
	GetArticleByCategory(c echo.Context) error
	GetArticleByID(c echo.Context) error
}