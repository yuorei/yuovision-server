package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/yuorei/video-server/app/driver/firebase"
)

type AuthMiddleware struct {
	authClient *firebase.AuthClient
}

func NewAuthMiddleware(authClient *firebase.AuthClient) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: authClient,
	}
}

func (am *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// No authentication provided - pass through for public endpoints
			next.ServeHTTP(w, r)
			return
		}

		// Extract Bearer token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		
		// Verify Firebase ID token
		token, err := am.authClient.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add Firebase UID to context
		ctx := context.WithValue(r.Context(), "firebase_uid", token.UID)
		ctx = context.WithValue(ctx, "userID", token.UID) // For compatibility with existing helper

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}