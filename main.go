package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"project_sdu/api"
	"project_sdu/db"
	"project_sdu/middleware"
	"project_sdu/model"
	repo "project_sdu/repository"
	"project_sdu/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPIHandler    api.UserAPI
	StudentAPIHandler api.StudentAPI
	ParentAPIHandler  api.ParentAPI
	PostAPIHandler    api.PostAPI
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] \"%s %s %s\"\n",
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// --- CORS SETUP HERE ---
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CORS_ALLOWED_ORIGINS")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// --- END CORS ---

	// Get DATABASE_URL
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("DATABASE_URL tidak ditemukan.")
	}

	// Connect to DB
	database := db.NewDB()
	conn, err := database.ConnectURL(databaseURL)
	if err != nil {
		panic(err)
	}

	// Migration
	conn.AutoMigrate(&model.User{}, &model.Student{}, &model.Parent{}, &model.Post{})

	// Route
	router = RunServer(router, conn)

	// Get Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("✅ Server is running on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		panic(err)
	}

}

func RunServer(r *gin.Engine, conn interface{}) *gin.Engine {
	dbConn := conn.(*gorm.DB)

	userRepo := repo.NewUserRepo(dbConn)
	studentRepo := repo.NewStudentRepo(dbConn)
	parentRepo := repo.NewParentRepo(dbConn)
	postRepo := repo.NewPostRepo(dbConn)

	userService := service.NewUserService(userRepo)
	studentService := service.NewStudentService(studentRepo, parentRepo)
	parentService := service.NewParentService(parentRepo)
	postService := service.NewPostService(postRepo)

	userAPIHandler := api.NewUserAPI(userService)
	studentAPIHandler := api.NewStudentAPI(studentService)
	parentAPIHandler := api.NewParentAPI(parentService)
	postAPIHandler := api.NewPostAPI(postService)

	apiHandler := APIHandler{
		UserAPIHandler:    userAPIHandler,
		StudentAPIHandler: studentAPIHandler,
		ParentAPIHandler:  parentAPIHandler,
		PostAPIHandler:    postAPIHandler,
	}

	// ROUTES //

	// User routes
	user := r.Group("/user")
	{
		user.POST("/register", apiHandler.UserAPIHandler.Register)
		user.POST("/login", apiHandler.UserAPIHandler.Login)
		user.POST("/logout", apiHandler.UserAPIHandler.Logout)

		user.Use(middleware.Auth())
		user.GET("/profile", apiHandler.UserAPIHandler.GetUserProfile)
	}

	// Student routes
	student := r.Group("/student")
	{
		student.Use(middleware.Auth())
		student.POST("/add", apiHandler.StudentAPIHandler.CreateStudent)
		student.POST("/bulk-add", apiHandler.StudentAPIHandler.CreateManyStudents)
		student.GET("/get/:id", apiHandler.StudentAPIHandler.GetStudentByID)
		student.GET("/get-all", apiHandler.StudentAPIHandler.GetAllStudents)
		student.PUT("/update/:id", apiHandler.StudentAPIHandler.UpdateStudent)
		student.DELETE("/delete/:id", apiHandler.StudentAPIHandler.DeleteStudent)

	}

	// Parent routes
	parent := r.Group("/parent")
	{
		parent.Use(middleware.Auth())
		parent.POST("/add", apiHandler.ParentAPIHandler.CreateParent)
		parent.GET("/get-all", apiHandler.ParentAPIHandler.GetAllParents)
		parent.GET("/get/:id", apiHandler.ParentAPIHandler.GetParentByID)
		parent.PUT("/update/:id", apiHandler.ParentAPIHandler.UpdateParent)
		parent.DELETE("/delete/:id", apiHandler.ParentAPIHandler.DeleteParent)
	}

	// Post routes
	post := r.Group("/post")
	{
		post.GET("/get/:slug", apiHandler.PostAPIHandler.GetPostBySlug)
		post.GET("/get-all", apiHandler.PostAPIHandler.GetAllPosts)

		post.Use(middleware.Auth())
		post.POST("/add", apiHandler.PostAPIHandler.CreatePost)
		post.PUT("/update/:id", apiHandler.PostAPIHandler.UpdatePost)
		post.DELETE("/delete/:id", apiHandler.PostAPIHandler.DeletePost)
	}

	return r
}
