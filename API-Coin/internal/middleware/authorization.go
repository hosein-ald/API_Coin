package middleware

import (
	"errors"
	"go_api/api"
	"net/http"

	"go_api/internal/tools"

	log "github.com/sirupsen/logrus"
)

var UnAthorizedError = errors.New("Invalid Username or Token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("Username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAthorizedError)
			api.RequestErrorHandler(w, UnAthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAthorizedError)
			api.RequestErrorHandler(w, UnAthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
