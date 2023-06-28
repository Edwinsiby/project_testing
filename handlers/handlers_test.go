package handlers_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"test/domain"
// 	"test/handlers"
// 	"test/repository"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAddProduct_Success(t *testing.T) {
// 	// Mock the repository
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error creating mock database: %v", err)
// 	}
// 	defer db.Close()
// 	repository.SetDB(db)

// 	// Create a new Gin router
// 	router := gin.Default()
// 	router.POST("/addproduct", handlers.AddProduct)

// 	// Prepare the request payload
// 	payload := domain.Product{
// 		ProductName: "Test Product",
// 		Price:       9.99,
// 		Catergory:   "Test Category",
// 	}
// 	requestBody := gin.H{"json": payload}

// 	// Set the expectations on the mock DB
// 	mock.ExpectQuery("SELECT COUNT(.+) FROM products").WithArgs(payload.ProductName).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
// 	mock.ExpectExec("INSERT INTO products").WithArgs(payload.ProductName, payload.Price, payload.Catergory).WillReturnResult(sqlmock.NewResult(0, 1))

// 	// Perform the HTTP request
// 	w := performRequest(router, "POST", "/addproduct", requestBody)

// 	// Assert the response
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{"success":"product created successfully"}`, w.Body.String())

// 	// Assert the expectations
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// func TestAddProduct_ProductAlreadyExists(t *testing.T) {
// 	// Mock the repository
// 	db, mock, err := repository.ConnectTestDB()
// 	if err != nil {
// 		t.Fatalf("Failed to connect to the test database: %v", err)
// 	}
// 	repository.SetDB(db)

// 	// Create a new Gin router
// 	router := gin.Default()
// 	router.POST("/addproduct", handlers.AddProduct)

// 	// Prepare the request payload
// 	payload := domain.Product{
// 		ProductName: "Test Product",
// 		Price:       9.99,
// 		Catergory:   "Test Category",
// 	}
// 	requestBody := gin.H{"json": payload}

// 	// Set the expectations on the mock DB
// 	mock.ExpectQuery("SELECT COUNT(.+) FROM products").WithArgs(payload.ProductName).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

// 	// Perform the HTTP request
// 	w := performRequest(router, "POST", "/addproduct", requestBody)

// 	// Assert the response
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// 	assert.JSONEq(t, `{"error":"product already exists"}`, w.Body.String())

// 	// Assert the expectations
// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// // Helper function to perform an HTTP request
// func performRequest(router *gin.Engine, method, path string, body gin.H) *httptest.ResponseRecorder {
// 	requestBody := gin.H{"json": body}
// 	jsonBody := ""

// 	if body != nil {
// 		jsonBody = requestBody["json"].(string)
// 	}

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(method, path, strings.NewReader(jsonBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	router.ServeHTTP(w, req)
// 	return w
// }
