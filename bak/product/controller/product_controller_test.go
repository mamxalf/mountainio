package controller

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mountainio/domain/entity"
	model2 "mountainio/domain/model"
	"net/http/httptest"
	"testing"
)

func TestProductController_Create(t *testing.T) {
	productRepository.DeleteAll()
	createProductRequest := model2.CreateProductRequest{
		Name:     "Test Product",
		Price:    1000,
		Quantity: 1000,
	}
	requestBody, _ := json.Marshal(createProductRequest)

	request := httptest.NewRequest("POST", "/api/products", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model2.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	createProductResponse := model2.CreateProductResponse{}
	json.Unmarshal(jsonData, &createProductResponse)
	assert.NotNil(t, createProductResponse.Id)
	assert.Equal(t, createProductRequest.Name, createProductResponse.Name)
	assert.Equal(t, createProductRequest.Price, createProductResponse.Price)
	assert.Equal(t, createProductRequest.Quantity, createProductResponse.Quantity)
}

func TestProductController_List(t *testing.T) {
	productRepository.DeleteAll()
	product := entity.Product{
		Id:       uuid.New().String(),
		Name:     "Sample Product",
		Price:    1000,
		Quantity: 1000,
	}
	productRepository.Insert(product)

	request := httptest.NewRequest("GET", "/api/products", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model2.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.([]interface{})
	containsProduct := false

	for _, data := range list {
		jsonData, _ := json.Marshal(data)
		getProductResponse := model2.GetProductResponse{}
		json.Unmarshal(jsonData, &getProductResponse)
		if getProductResponse.Id == product.Id {
			containsProduct = true
		}
	}

	assert.True(t, containsProduct)
}
