package web

import (
	"myGoLearn/week02_gin_gorm/internal/domain"
	"myGoLearn/week02_gin_gorm/internal/service"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	// EmailRegExp 邮箱校验正则
	EmailRegExp = "\\w+([-+.]\\w+)*@"
)

type UserHandler struct {
	EmailRegExp *regexp.Regexp
	service     *service.UserService
}

type UserClaim struct {
	UId       int64
	UserAgent string
	jwt.RegisteredClaims
}

const JWT_KEY = "zhuwei secret"

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		EmailRegExp: regexp.MustCompile(EmailRegExp),
		service:     userService,
	}
}

func (h *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "我是zhuzhu")
}
func (h *UserHandler) SignUp(ctx *gin.Context) {
	type UserReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req UserReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	isEmail := h.EmailRegExp.Match([]byte(req.Email))
	if !isEmail {
		ctx.String(http.StatusOK, "邮箱错了！！！")
		return
	}

	err := h.service.Signup(ctx, &domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	switch err {
	case nil:
		ctx.String(http.StatusOK, req.Email+"注册成功")
	case service.EmailDuplicateError:
		ctx.String(http.StatusOK, "邮箱冲突咯")
	default:
		ctx.String(http.StatusOK, "系统错误")

	}

}

func (h *UserHandler) LoginWithJWT(ctx *gin.Context) {
	type LoginReq struct {
		Email    string
		Password string
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := h.service.Login(ctx, req.Email, req.Password)
	switch err {
	case nil:
		uc := UserClaim{
			UId:       user.ID,
			UserAgent: ctx.GetHeader("User-Agent"),
			RegisteredClaims: jwt.RegisteredClaims{
				//30分钟过期
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)
		tokenStr, err := token.SignedString([]byte(JWT_KEY))
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
		}
		ctx.Header("x-jwt-token", tokenStr)

		ctx.String(http.StatusOK, user.Email+"登录成功")
	case service.InvalidEmailOrPassword:
		ctx.String(http.StatusOK, "邮箱或密码错误")
	default:
		ctx.String(http.StatusOK, "系统错误")

	}

}
func (h *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string
		Password string
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := h.service.Login(ctx, req.Email, req.Password)
	switch err {
	case nil:
		sess := sessions.Default(ctx)
		sess.Set("userId", user.ID)
		sess.Options(sessions.Options{
			//MaxAge: 60 * 15, //15分钟
			MaxAge: 30,
		})
		err := sess.Save()
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
			return
		}

		ctx.String(http.StatusOK, user.Email+"登录成功")
	case service.InvalidEmailOrPassword:
		ctx.String(http.StatusOK, "邮箱或密码错误")
	default:
		ctx.String(http.StatusOK, "系统错误")

	}

}

func (h *UserHandler) Edit(ctx *gin.Context) {
	type EditReq struct {
		AboutMe string
	}
	var req EditReq
	//绑定
	if err := ctx.Bind(&req); err != nil {
		return
	}

	//sess := sessions.Default(ctx)
	//userId := sess.Get("userId")

	userId := ctx.MustGet("userId").(int64)

	err := h.service.Edit(ctx, &domain.User{
		AboutMe: req.AboutMe,
		ID:      userId,
	})
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	ctx.String(http.StatusOK, "编辑成功")
}
func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	group := server.Group("/users")
	group.GET("/profile", h.Profile)
	group.POST("/signup", h.SignUp)
	//group.POST("/login", h.Login)
	group.POST("/login", h.LoginWithJWT)
	group.POST("/edit", h.Edit)

}
