package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cristiangar0398/REST-API-CRUD/middleware"
	"github.com/cristiangar0398/REST-API-CRUD/models"
	"github.com/cristiangar0398/REST-API-CRUD/repository"
	"github.com/cristiangar0398/REST-API-CRUD/server"
	"github.com/segmentio/ksuid"
)

type InsertPostRequest struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	Id          string `json:"id`
	PostContent string `json:"post_content"`
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := middleware.TokenParseString(w, s, r)
		if err != nil {
			log.Fatal(err)
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var postRequest = InsertPostRequest{}
			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			post := models.Post{
				Id:           id.String(),
				Post_content: postRequest.PostContent,
				UserId:       claims.UserId,
			}
			err = repository.InsertPost(r.Context(), &post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("content-type", "aaplication/json")

			json.NewEncoder(w).Encode(PostResponse{
				Id:          post.Id,
				PostContent: post.Post_content,
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
