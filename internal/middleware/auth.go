package middleware

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const UserEmailKey contextKey = "userEmail"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"success":false,"message":"missing authorization header"}`, http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			http.Error(w, `{"success":false,"message":"invalid authorization format"}`, http.StatusUnauthorized)
			return
		}

		// TODO: verify Firebase ID token
		// firebaseApp, err := firebase.NewApp(...)
		// authClient, err := firebaseApp.Auth(ctx)
		// decodedToken, err := authClient.VerifyIDToken(ctx, token)
		// email := decodedToken.Claims["email"]
		email := token

		ctx := context.WithValue(r.Context(), UserEmailKey, email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserEmail(ctx context.Context) string {
	email, _ := ctx.Value(UserEmailKey).(string)
	return email
}
