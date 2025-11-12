package middlewares

import (
	"fmt"
	"myGoLearn/week02_gin_gorm/internal/web"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginJWTMiddleware struct{}

func (login *LoginJWTMiddleware) CheckLogin() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			//不需要校验
			return
		}
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenSplit := strings.Split(auth, " ")
		if len(tokenSplit) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := tokenSplit[1]
		var uc web.UserClaim
		token, err := jwt.ParseWithClaims(
			tokenString,
			&uc,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(web.JWT_KEY), nil
			},
		)
		if uc.UserAgent != ctx.GetHeader("User-Agent") {
			//是攻击者来了！！！
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err != nil || !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//每10s刷新一下token，即token的剩余时间小于50s就需要刷新一下
		expiredAt := uc.ExpiresAt
		if expiredAt.Sub(time.Now()) < time.Second*50 {
			//uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			tokenStr, err := token.SignedString([]byte(web.JWT_KEY))
			fmt.Println(err)
			ctx.Header("x-jwt-token", tokenStr)
		}
		ctx.Set("userId", uc.UId)
	}
}
