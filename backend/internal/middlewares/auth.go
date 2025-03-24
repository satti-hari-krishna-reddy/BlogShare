package middlewares

import (
	"context"
	"net/http"
	"time"
	"fmt"

	"social-scribe/backend/internal/models"
	repo "social-scribe/backend/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type contextKey string

const UserIDKey contextKey = "userID"

// AuthMiddleware handles authentication and rate limiting
func AuthMiddleware(limit int, duration time.Duration, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate session
		cookie, err := r.Cookie("session_token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"success": false, "reason": "Unauthorized: Missing session token"}`))
			return
		}

		sessionData, exists := repo.GetCache(cookie.Value)
		if !exists {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"success": false, "reason": "Unauthorized: Invalid or expired session"}`))
			return
		}

		session, ok := sessionData.(models.CacheItem)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"success": false, "reason": "Unauthorized: Invalid session format"}`))
			return
		}

		if session.ExpiresAt.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"success": false, "reason": "Unauthorized: Session expired"}`))
			return
		}

		oid, ok := session.Value.(primitive.ObjectID)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"success": false, "reason": "Unauthorized: Invalid user ID format"}`))
			return
		}

		userID := oid.Hex()

		// checking for rate limiting
		if repo.IsRateLimited(userID, limit, duration) {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"success": false, "reason": "Too Many Requests"}`))
			return
		}

		// check for CSRF
		cacheKey := fmt.Sprintf("CSRF_%s", userID)
		tokenData, exists := repo.GetCache(cacheKey)
		if !exists {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"success": false, "reason": "CSRF token missing"}`))
			return
		}

		tokenInfo, _ := tokenData.(models.CacheItem)
		expectedToken := tokenInfo.Value

		// Get the token from the request header
		requestToken := r.Header.Get("X-CSRF-Token")
		if requestToken == "" || requestToken != expectedToken {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"success": false, "reason": "Invalid CSRF token"}`))
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		w.Header().Set("Content-Security-Policy", "frame-ancestors 'self'")
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=(), usb=()")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
