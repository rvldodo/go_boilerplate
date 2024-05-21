package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/rvldodo/boilerplate/domain/store"
	"github.com/rvldodo/boilerplate/lib/jwt"
	"github.com/rvldodo/boilerplate/lib/log"
	"github.com/rvldodo/boilerplate/utils"
)

var UserKey = "UserID"

func AuthWithJWT(handleFunc http.HandlerFunc, store store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := jwt.GetTokenFromRequest(r)

		token, err := jwt.ValidateToken(tokenString)
		if err != nil {
			log.Errorf("Invalid token: %v", err)
			utils.WriteError(w, http.StatusForbidden, err)
			return
		}

		if !token.Valid {
			log.Error("Invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["userID"].(string)
		u, _ := uuid.Parse(userID)

		us, err := store.FindById(r.Context(), u)
		if err != nil {
			log.Errorf("Failed get user by id: %v", err)
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, us.ID)
		r.WithContext(ctx)

		handleFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}
	return userID
}
