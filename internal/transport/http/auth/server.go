package auth

import (
	"github.com/POMBNK/gateway/internal/service/auth"
	"github.com/POMBNK/gateway/pkg/jwt"
	"net/http"
	"os"
)

type Server struct {
	authService *auth.Service
}

func NewServer(authService *auth.Service) *Server {
	return &Server{
		authService: authService,
	}
}

func (s *Server) Register(mux *http.ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(s, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: mux,
		Middlewares: []MiddlewareFunc{
			MiddlewareFunc(jwt.MiddlewareOAPI(os.Getenv("SECRET"))),
		},
	})
}
