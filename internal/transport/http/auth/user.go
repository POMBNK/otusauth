package auth

import (
	"encoding/json"
	"fmt"
	"github.com/POMBNK/gateway/pkg/jwt"
	"net/http"
	"os"
	"strconv"
)

func (s *Server) FindUserByUsername(w http.ResponseWriter, r *http.Request, username string) {
	if err := accessStatus(r, nil, &username); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := s.authService.FindUserByUsername(r.Context(), username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	resp := user.ToResponse()

	userBytes, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userBytes)
}
func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request, userId int64) {
	if err := accessStatus(r, &userId, nil); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}

	err := s.authService.DeleteUser(r.Context(), int(userId))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("deleted successfully"))
}

func (s *Server) FindUserById(w http.ResponseWriter, r *http.Request, userId int64) {
	if err := accessStatus(r, &userId, nil); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}
	user, err := s.authService.FindUserBydID(r.Context(), int(userId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(userBytes)
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request, userId int64) {
	if err := accessStatus(r, &userId, nil); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}

	var body UpdateUserJSONRequestBody
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

	err = s.authService.UpdateUser(r.Context(), int(userId), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	updatedUser, err := s.authService.FindUserBydID(r.Context(), int(userId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	tokenizer := jwt.NewTokenizer(os.Getenv("SECRET"))
	pair, err := tokenizer.GeneratePair(strconv.Itoa(updatedUser.ID), updatedUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	acook, rcook := tokenizer.PrepareCookies(pair)
	http.SetCookie(w, acook)
	http.SetCookie(w, rcook)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("updated successfully"))
}

func accessStatus(r *http.Request, userId *int64, username *string) error {
	userID, err := strconv.ParseInt(r.Context().Value("user_id").(string), 10, 64)
	if err != nil {
		return fmt.Errorf("you are not allowed to perform this action")
	}

	if userId == nil && username == nil {
		return fmt.Errorf("you are not allowed to perform this action")
	}

	usernamectx := r.Context().Value("username").(string)
	if username != nil {
		if usernamectx != *username {
			return fmt.Errorf("you are not allowed to perform this action")
		}
	}

	if userId != nil {
		if userID != *userId {
			return fmt.Errorf("you are not allowed to perform this action")
		}

	}

	return nil
}
