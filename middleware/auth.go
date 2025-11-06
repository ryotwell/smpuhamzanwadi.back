package middleware

import (
	"net/http"
	"project_sdu/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("session_token")
		if err != nil {
			// if ctx.GetHeader("Content-Type") == "application/json" {
			// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			// } else {
			// 	ctx.Redirect(http.StatusSeeOther, "/login")
			// }
			// ctx.Abort()
			// return

			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad requesto"})
			}
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("id", claims.UserID)

		ctx.Next()
		// TODO: answer here
	})
}
