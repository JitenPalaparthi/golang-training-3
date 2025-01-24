package main

import (
	"demo/database"
	"demo/handlers"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	dsn := os.Getenv("DB_CONN")
	//isFilebased := false
	if port == "" {
		flag.StringVar(&port, "port", "8086", "--port=8086 or -port=8086 or --port 8086 or -port 8086")
		//flag.BoolVar(&isFilebased, "file", false, "--file=true | -file=false")
		if dsn == "" {
			flag.StringVar(&dsn, "dsn", `host=localhost user=admin password=admin123 dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "-dsn=db connections string")
		}
		flag.Parse()
	}

	db, err := database.GetConnection(dsn)
	if err != nil {
		panic(err)
	}
	udb := database.NewUserDB(db)
	uhadler := handlers.NewUserHandler(udb)

	router := gin.Default()

	public := router.Group("/public")

	public.GET("/ping", handlers.Ping)

	public.GET("/health", handlers.Health)

	public.POST("/login", handlers.Login)

	privateUser := router.Group("/private/users")

	//privateUser.Use(handlers.Restrict)
	//privateUser.Use(handlers.AuthMiddleware())
	privateUser.POST("/", uhadler.Create)
	privateUser.GET("/:id", uhadler.GetByID)
	privateUser.DELETE("/:id", uhadler.DeleteByID)
	if err := router.Run(":" + port); err != nil {
		log.Println(err)
	}

}

func AuthenticateByUserNme(ctx *gin.Context) {
	username := ctx.GetHeader("username")
	password := ctx.GetHeader("password")
	if username == "" || password == "" {
		ctx.String(http.StatusUnauthorized, "invalid username or password")
		ctx.Abort()
	}
	if username == "jiten" && password == "jiten123" {
		ctx.Next()
	} else {
		ctx.String(http.StatusUnauthorized, "invalid username or password")
		ctx.Abort()
	}
}
