package auth

import (
	"github.com/POMBNK/gateway/internal/service/auth"
	"github.com/POMBNK/gateway/pkg/jwt"
	"github.com/rakyll/statik/fs"
	"log"
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

func (s *Server) Register(baseURL string) http.Handler {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/sign-in", s.SignIn)
	mux.HandleFunc("/api/v1/sign-up", s.SignUp)
	mux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/swagger/", http.FileServer(statikFS)).ServeHTTP(w, r)
	})

	return HandlerWithOptions(s, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: mux,
		Middlewares: []MiddlewareFunc{
			MiddlewareFunc(jwt.MiddlewareOAPI(os.Getenv("SECRET"))),
		},
	})
}
