package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"gogo-form/database"
	"gogo-form/domain"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

var formHandler FormHandler

var formId string
var payloadForm = domain.Form{
	Name:        "Test Form",
	Description: "This is a test form for unit testing",
	Questions: []domain.Question{
		{
			Text: "Do you like Go programming language?",
			Type: "boolean",
		},
	},
}

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	gin.SetMode("test")
	database.InitDB()
	formHandler = NewFormHandler()

	collection := database.GetCollection("forms")

	res, _ := collection.InsertOne(context.Background(), payloadForm)
	formId = res.InsertedID.(primitive.ObjectID).Hex()

	os.Exit(m.Run())
}

func setupTestContext(method string, path string, data []byte, params ...gin.Param) (*gin.Context, *httptest.ResponseRecorder) {
	var req *http.Request

	if len(data) != 0 {
		req, _ = http.NewRequest(method, path, bytes.NewBuffer(data))
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	if len(params) > 0 {
		ctx.Params = params
	}

	return ctx, w
}

// Common flow //

func TestCreateForm(t *testing.T) {
	form := payloadForm
	form.Name = "Created form"

	data, _ := json.Marshal(form)
	ctx, w := setupTestContext("POST", "/form", data)

	formHandler.Create(ctx)

	assert.Equal(t, 201, w.Code)
}

func TestGetAllForms(t *testing.T) {
	ctx, w := setupTestContext("GET", "/form", nil)

	formHandler.GetAll(ctx)

	assert.Equal(t, 200, w.Code)
}

func TestGetOneForm(t *testing.T) {
	ctx, w := setupTestContext("GET", "/form/"+formId, nil, gin.Param{Key: "id", Value: formId})

	formHandler.GetOne(ctx)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateForm(t *testing.T) {
	form := payloadForm
	form.Name = "Updated form"

	data, _ := json.Marshal(form)
	ctx, w := setupTestContext("PUT", "/form/"+formId, data, gin.Param{Key: "id", Value: formId})

	formHandler.Update(ctx)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteForm(t *testing.T) {
	ctx, w := setupTestContext("DELETE", "/form/"+formId, nil, gin.Param{Key: "id", Value: formId})

	formHandler.Delete(ctx)

	assert.Equal(t, 200, w.Code)
}

// User error flow //

func TestCreateInvalidForm(t *testing.T) {
	data, _ := json.Marshal(domain.Form{})
	ctx, w := setupTestContext("POST", "/form", data)

	formHandler.Create(ctx)

	assert.Equal(t, 422, w.Code)
}

func testUpdateInvalidForm(t *testing.T) {
	data, _ := json.Marshal(domain.Form{})
	ctx, w := setupTestContext("PUT", "/form/"+formId, data, gin.Param{Key: "id", Value: formId})

	formHandler.Update(ctx)

	assert.Equal(t, 422, w.Code)
}

func TestCreateFormNoJson(t *testing.T) {
	data := []byte("no json")
	ctx, w := setupTestContext("POST", "/form", data)

	formHandler.Create(ctx)

	assert.Equal(t, 400, w.Code)
}

func testUpdateFormNoJson(t *testing.T) {
	data := []byte("no json")
	ctx, w := setupTestContext("PUT", "/form/"+formId, data, gin.Param{Key: "id", Value: formId})

	formHandler.Update(ctx)

	assert.Equal(t, 400, w.Code)
}

func TestGetOneUnexistingForm(t *testing.T) {
	ctx, w := setupTestContext("GET", "/form/000", nil, gin.Param{Key: "id", Value: "000"})

	formHandler.GetOne(ctx)

	assert.Equal(t, 404, w.Code)
}

func TestUpdateUnexistingForm(t *testing.T) {
	form := payloadForm
	form.Name = "Updated form"

	data, _ := json.Marshal(form)
	ctx, w := setupTestContext("PUT", "/form/000", data, gin.Param{Key: "id", Value: "000"})

	formHandler.Update(ctx)

	assert.Equal(t, 404, w.Code)
}

func TestDeleteUnexistingForm(t *testing.T) {
	ctx, w := setupTestContext("DELETE", "/form/000", nil, gin.Param{Key: "id", Value: "000"})

	formHandler.Delete(ctx)

	assert.Equal(t, 404, w.Code)
}
