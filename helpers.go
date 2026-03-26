package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/api-service/api-service/models"
	"github.com/api-service/api-service/utils"
	"github.com/gorilla/sessions"
)

func ParseQueryToMap(r *http.Request) map[string]string {
	query := r.URL.Query()
	mapQuery := make(map[string]string)

	for k, v := range query {
		mapQuery[k] = strings.Join(v, ",")
	}

	return mapQuery
}

func GetErrorStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrInvalidData:
		return http.StatusBadRequest
	case models.ErrInternalError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func GetErrorResponse(err error) (int, interface{}) {
	status := GetErrorStatus(err)
	var message string
	var data interface{}

	switch err {
	case models.ErrNotFound:
		message = "Resource not found"
	case models.ErrInvalidData:
		message = "Invalid data"
	case models.ErrInternalError:
		message = "Internal server error"
	default:
		message = "Internal server error"
	}

	return status, map[string]string{"error": message}
}

func GetPaginationResponse(total int, limit int, offset int) (int, int, int) {
	totalPages := (total - 1) / limit
	if totalPages < 0 {
		totalPages = 0
	}

	return totalPages, offset, offset + limit
}

func GetJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}

func GetValidatedData(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}

	return nil
}

func GetSessionStore(r *http.Request) *sessions.Store {
	sessionStore := sessions.NewCookieStore([]byte("secret"))
	sessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}

	return sessionStore
}

func GetUserIDFromSession(r *http.Request) (string, error) {
	sessionStore := GetSessionStore(r)
	session, err := sessionStore.Get(r, "session")
	if err != nil {
		return "", err
	}

	id, err := session.Get("user_id")
	if err != nil {
		return "", err
	}

	return id.(string), nil
}

func GetOffsetFromQuery(r *http.Request) (int, error) {
	offset := r.URL.Query().Get("offset")
	if offset == "" {
		return 0, nil
	}

	offsetInt, err := strconv.Atoi(offset)
	return offsetInt, err
}