package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bquenin/microservice/cmd/microservice/config"
	"github.com/bquenin/microservice/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AuthorServiceTestSuite struct {
	suite.Suite
	router *gin.Engine
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *AuthorServiceTestSuite) SetupTest() {
	cfg, err := config.Read()
	suite.Require().NoError(err)

	postgres, err := database.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password)
	suite.Require().NoError(err)

	authorService := NewAuthorService(postgres)

	suite.router = gin.Default()
	suite.router.POST("/authors", authorService.Create)
	suite.router.GET("/authors", authorService.List)
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *AuthorServiceTestSuite) TestCreateAuthor() {
	request := CreateAuthorRequest{
		Name: "test author",
		Bio:  "A test bio",
	}
	var buffer bytes.Buffer
	suite.Require().NoError(json.NewEncoder(&buffer).Encode(request))

	req, err := http.NewRequest("POST", "/authors", &buffer)
	suite.Require().NoError(err)

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
	var author database.Author
	suite.Require().NoError(json.NewDecoder(rec.Result().Body).Decode(&author))
	suite.Require().NotZero(author.ID)
	suite.Require().Equal(request.Name, author.Name)
	suite.Require().Equal(request.Bio, author.Bio)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAuthorServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorServiceTestSuite))
}
