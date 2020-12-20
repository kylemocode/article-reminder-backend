package handlers

import (
	"article-reminder/domain"
	"fmt"
	"net/http"
)

func (s *Server) registerUser() http.HandlerFunc {
	var payload domain.RegisterPayload
	return validatePayload(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("payload", payload)
	}, &payload)
}
