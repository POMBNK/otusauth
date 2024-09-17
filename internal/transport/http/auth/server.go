package auth

import (
	"encoding/json"
	"fmt"
	"github.com/POMBNK/gateway/internal/service/auth"
	"github.com/POMBNK/gateway/pkg/jwt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	authService *auth.Service
}

func NewServer(authService *auth.Service) *Server {
	return &Server{
		authService: authService,
	}
}

func (s *Server) Register(baseURL string) http.Handler {
	return HandlerWithOptions(s, StdHTTPServerOptions{
		BaseURL: "/api/v1",
	})
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {

	var body SignInJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	user, err := body.ToUser()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	newUser, err := s.authService.SignIn(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	tokenizer := jwt.NewTokenizer(os.Getenv("SECRET"))
	pair, err := tokenizer.GeneratePair(strconv.Itoa(newUser.ID), newUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	acook, rcook := tokenizer.PrepareCookies(pair)

	http.SetCookie(w, acook)
	http.SetCookie(w, rcook)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("signed in successfully. Welcome %s", newUser.Username)))
}

// todo: перетащить в сервис пользователей, тут остается только validate
func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {

	var body SignUpJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := body.ToUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = s.authService.SingUp(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("signed up successfully"))
}

func (s *Server) Validate(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("validated successfully. Welcome %s", username)))
}
