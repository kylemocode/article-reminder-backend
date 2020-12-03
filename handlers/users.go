package handlers

import (
	"article-reminder/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) registerUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := domain.RegisterPayload{}

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		fmt.Println("payload", payload)
		// user, err := s.domain.Register()
	}
}
