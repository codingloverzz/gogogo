package middlewares

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddleware struct{}

func (login *LoginMiddleware) CheckLogin() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			//不需要校验
			return
		}
		sess := sessions.Default(ctx)
		userId := sess.Get("userId")
		if userId == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//实现每分钟都去刷新一下session，避免用着用着就过期
		const updateTimeKey = "lastUpdateTime"
		now := time.Now()
		lastUpdateTime := sess.Get(updateTimeKey)
		if lastUpdateTime == nil || now.Sub(lastUpdateTime.(time.Time)) > time.Minute {
			sess.Set(updateTimeKey, now)
			sess.Set("userId", userId)
			sess.Save()
		}

	}
}
