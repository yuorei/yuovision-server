package middleware

import (
	"context"
	"log/slog"
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
		slog.Debug("attempting to verify Firebase ID token", "token_length", len(idToken))
		token, err := am.authClient.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			slog.Error("Firebase token verification failed", "error", err, "token_prefix", idToken[:min(10, len(idToken))])
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		slog.Debug("Firebase token verified successfully", "uid", token.UID)

		// Add Firebase UID to context
		ctx := context.WithValue(r.Context(), "firebase_uid", token.UID)
		ctx = context.WithValue(ctx, "userID", token.UID) // For compatibility with existing helper

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
