package tests

import (
	"bytes"
	"encoding/json"
	"helpdesk/app/user"
	"helpdesk/database"
	"helpdesk/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db = database.Init()

func TestUserIndex(t *testing.T) {
	existingUser := user.User{
		FullName:   "Alice",
		Email:      "alice@example.com",
		Password:   "password",
		IsActive:   true,
		IsArchived: false,
	}

	db.Create(&existingUser)

	resp, err := http.Get("http://localhost:3000/api/v1/users")

	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

	defer resp.Body.Close()

	var data utils.ApiResponseWithData[[]user.User]

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Users retrieved successfully", data.Message)
	assert.Equal(t, existingUser.FullName, data.Data[0].FullName)
	assert.Equal(t, existingUser.Email, data.Data[0].Email)
	assert.Empty(t, data.Data[0].Password)
	assert.Equal(t, existingUser.IsActive, data.Data[0].IsActive)
	assert.Equal(t, existingUser.IsArchived, data.Data[0].IsArchived)
}

func TestUserShow(t *testing.T) {
	existingUser := user.User{
		FullName:   "Bob",
		Email:      "bob@example.com",
		Password:   "password",
		IsActive:   true,
		IsArchived: false,
	}

	db.Create(&existingUser)

	resp, err := http.Get("http://localhost:3000/api/v1/users/" + existingUser.Id)

	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

	defer resp.Body.Close()

	var data utils.ApiResponseWithData[user.User]

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "User retrieved successfully", data.Message)
	assert.Equal(t, existingUser.Id, data.Data.Id)
	assert.Equal(t, existingUser.FullName, data.Data.FullName)
	assert.Equal(t, existingUser.Email, data.Data.Email)
	assert.Empty(t, data.Data.Password)
	assert.Equal(t, existingUser.IsActive, data.Data.IsActive)
	assert.Equal(t, existingUser.IsArchived, data.Data.IsArchived)
}

func TestUserStore(t *testing.T) {
	newUser := user.User{
		FullName:   "Charlie",
		Email:      "charlie@example.com",
		Password:   "password",
		IsActive:   true,
		IsArchived: false,
	}

	payload, err := json.Marshal(newUser)

	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

	resp, err := http.Post("http://localhost:3000/api/v1/users", "application/json", bytes.NewBuffer(payload))

	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

	defer resp.Body.Close()

	var data utils.ApiResponseWithData[user.User]

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, "User created successfully", data.Message)
	assert.NotEmpty(t, data.Data.Id)
	assert.Equal(t, newUser.FullName, data.Data.FullName)
	assert.Equal(t, newUser.Email, data.Data.Email)
	assert.Empty(t, data.Data.Password)
	assert.Equal(t, newUser.IsActive, data.Data.IsActive)
	assert.Equal(t, newUser.IsArchived, data.Data.IsArchived)
}

func TestUserUpdate(t *testing.T) {
	existingUser := user.User{
		FullName:   "David",
		Email:      "david@example.com",
		Password:   "password",
		IsActive:   true,
		IsArchived: false,
	}

	db.Create(&existingUser)

	updatedUser := map[string]interface{}{
		"FullName": "David Jr.",
	}

	payload, err := json.Marshal(updatedUser)

	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

	req, err := http.NewRequest(http.MethodPatch, "http://localhost:3000/api/v1/users/"+existingUser.Id, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

	defer resp.Body.Close()

	var data utils.ApiResponseWithData[user.User]

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		t.Error("Expected error to be nil, got", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "User updated successfully", data.Message)
	assert.Equal(t, existingUser.Id, data.Data.Id)
	assert.Equal(t, updatedUser["FullName"].(string), data.Data.FullName)
	assert.Equal(t, existingUser.Email, data.Data.Email)
	assert.Empty(t, data.Data.Password)
	assert.Equal(t, existingUser.IsActive, data.Data.IsActive)
	assert.Equal(t, existingUser.IsArchived, data.Data.IsArchived)
}
