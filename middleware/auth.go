package middleware

import (
	"context"
	"fmt"

	"os"
	"strings"

	"net/http"

	"github.com/golang-jwt/jwt"
)

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				next.ServeHTTP(w, r)
				return
			}
			bearerToken := strings.Split(auth, " ")[1]

			// ルートの公開鍵を読み込む
			publicKeyData, err := os.ReadFile("public_key.pem")
			if err != nil {
				http.Error(w, "Failed to load public key error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
			if err != nil {
				http.Error(w, "Public key parsing failed error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// トークンを検証する
			token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
				return publicKey, nil
			})
			if err != nil {
				if err.Error() == "Token is expired" {
					next.ServeHTTP(w, r)
					return // トークンが期限切れの場合は、次のハンドラに進む
				}
				http.Error(w, "Token parsing failed error: "+err.Error(), http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := addIDAndNameToContext(r.Context(), claims["sub"].(string), claims["preferred_username"].(string))
				r = r.WithContext(ctx)
			} else {
				http.Error(w, "Token validation failed error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func addIDAndNameToContext(ctx context.Context, id, name string) context.Context {
	ctx = context.WithValue(ctx, "id", id)
	ctx = context.WithValue(ctx, "name", name)
	return ctx
}

func GetUserIDFromContext(ctx context.Context) (string, error) {
	id := ctx.Value("id")
	if id == nil {
		return "", fmt.Errorf("id is nil")
	}
	return id.(string), nil
}

func GetNameFromContext(ctx context.Context) (string, error) {
	name := ctx.Value("name")
	if name == nil {
		return "", fmt.Errorf("name is nil")
	}
	return name.(string), nil
}
