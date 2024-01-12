package middleware

import (
	"context"
	"log"

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

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				log.Println(claims["sub"])
				log.Println(claims["preferred_username"])
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
