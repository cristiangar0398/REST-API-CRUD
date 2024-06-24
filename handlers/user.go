package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cristiangar0398/REST-API-CRUD/models"
	"github.com/cristiangar0398/REST-API-CRUD/repository"
	"github.com/cristiangar0398/REST-API-CRUD/server"
	"github.com/segmentio/ksuid"
)

var (
	request  SignUpRequest
	response SignUpResponse
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		request, err := decodeSignUpRequest(r)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		userID, err := generateUserID()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := createUser(r.Context(), request, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}
}

func decodeSignUpRequest(r *http.Request) (SignUpRequest, error) {
	var request SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func generateUserID() (string, error) {
	id, err := ksuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func createUser(ctx context.Context, request SignUpRequest, userID string) (*models.User, error) {
	user := &models.User{
		Email:    request.Email,
		Password: request.Password,
		Id:       userID,
	}
	err := repository.InsertUser(ctx, user)
	return user, err
}
