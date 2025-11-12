package main

import (
	"myGoLearn/week02_gin_gorm/internal/repository"
	"myGoLearn/week02_gin_gorm/internal/repository/dao"
	"myGoLearn/week02_gin_gorm/internal/service"
	"myGoLearn/week02_gin_gorm/internal/web"
	"myGoLearn/week02_gin_gorm/internal/web/middlewares"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db := initDB()

	err := dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	server := initWebServer()
	initUser(server, db)

	err = server.Run(":8080")
	if err != nil {
		return
	}
}

func useSession(server *gin.Engine) {
	login := &middlewares.LoginMiddleware{}

	//基于cookie
	//store := cookie.NewStore([]byte("secret"))
	// 基于内存
	//store := memstore.NewStore([]byte(strings.Repeat("1", 32)))

	//基于redis实现
	store, err := redis.NewStore(12, "tcp", "localhost:6379", "", "", []byte(strings.Repeat("1", 32)))
	if err != nil {
		panic(err)
	}
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())

}
func useJWT(server *gin.Engine) {
	login := &middlewares.LoginJWTMiddleware{}
	server.Use(login.CheckLogin())
}
func initWebServer() *gin.Engine {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"x-jwt-token"},
	}))

	useJWT(server)
	//useSession(server)
	return server
}
func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	return db
}
func initUser(server *gin.Engine, db *gorm.DB) {
	userDao := dao.NewUserDao(db)
	userRepository := repository.NewUserRepository(userDao)
	userService := service.NewUserService(userRepository)
	userHandler := web.NewUserHandler(userService)

	userHandler.RegisterRoutes(server)

}
