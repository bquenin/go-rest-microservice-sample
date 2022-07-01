package api

import (
	"context"
	"net/http"

	"github.com/bquenin/microservice/internal/database"
	"github.com/gin-gonic/gin"
)

type AuthorService struct {
	queries *database.Queries
}

func NewAuthorService(postgres *database.Postgres) *AuthorService {
	return &AuthorService{queries: database.New(postgres.DB)}
}

type CreateAuthorRequest struct {
	Name string
	Bio  string
}

func (as *AuthorService) Create(c *gin.Context) {
	// Parse request
	var request CreateAuthorRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create author
	params := database.CreateAuthorParams{
		Name: request.Name,
		Bio:  request.Bio,
	}
	author, err := as.queries.CreateAuthor(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, author)
}

func (as *AuthorService) List(c *gin.Context) {
	authors, err := as.queries.ListAuthors(context.Background())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	if authors == nil {
		authors = []database.Author{}
	}
	c.IndentedJSON(http.StatusOK, authors)
}
