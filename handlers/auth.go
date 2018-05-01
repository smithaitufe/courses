package handlers

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	cc "github.com/smithaitufe/courses/context"
	ce "github.com/smithaitufe/courses/errors"
	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/services"
)

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var isAuthenticated bool
		var userID *string
		ctx := r.Context()
		token, err := validatePassedTokens(ctx, r)
		ce.LogOnError("No token was found", err)
		if token != nil {
			if token.Valid {
				userID = getUserIDFromToken(token)
				ctx = writeResponse(ctx, userID, isAuthenticated)
				h.ServeHTTP(w, r.WithContext(ctx))
			} else if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					ce.LogOnError("That's not even a token", err)
				} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
					ce.LogOnError("Timing is everything", err)
					userID = getUserIDFromToken(token)
					user, err := ctx.Value(ck.UserServiceKey).(*services.UserService).GetUser(userID)
					authService := ctx.Value(ck.AuthServiceKey).(*services.AuthService)
					newToken, err := authService.GenerateToken(user)
					ce.FailOnError("Could not generate tokens", err)
					newRefreshToken, err := authService.GenerateRefreshToken(user)
					ce.FailOnError("Could not generate tokens", err)
					ctx = writeResponse(ctx, userID, isAuthenticated)
					ctx = context.WithValue(ctx, ck.TokenKey, newToken)
					ctx = context.WithValue(ctx, ck.RefreshTokenKey, newRefreshToken)
					h.ServeHTTP(w, r.WithContext(ctx))
				} else {
					ce.LogOnError("Couldn't handle this token:", err)
				}
			}
		}

		ctx = writeResponse(ctx, userID, isAuthenticated)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
func writeResponse(ctx context.Context, userID *string, isAuthenticated bool) context.Context {
	ctx = context.WithValue(ctx, ck.UserIDKey, userID)
	ctx = context.WithValue(ctx, ck.IsAuthenticatedKey, isAuthenticated)
	return ctx
}
func getUserIDFromToken(token *jwt.Token) *string {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userIDByte, _ := base64.StdEncoding.DecodeString(claims["id"].(string))
		userID := string(userIDByte[:])
		return &userID
	}
	return nil
}
func validatePassedTokens(ctx context.Context, r *http.Request) (*jwt.Token, error) {
	var tokenString string
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Bearer" {
		return nil, errors.New(cc.TokenError)
	}
	tokenString = auth[1]
	token, err := ctx.Value(ck.AuthServiceKey).(*services.AuthService).ValidateToken(&tokenString)
	return token, err
}
