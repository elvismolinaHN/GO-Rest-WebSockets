package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/segmentio/ksuid"
	"platzi.com/go/rest-ws/models"
	"platzi.com/go/rest-ws/repository"
	"platzi.com/go/rest-ws/server"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id    string `json:id`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request) // Le envia el cuerpo de la peticion
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // Este error especifica que fue culpa del cliente
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // Este error indica que fue culpa del servidor
			return
		}

		var user = models.User{
			Email:    request.Email,
			Password: request.Password,
			Id:       id.String(),
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}
}
