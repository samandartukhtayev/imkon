package api

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	_ "github.com/samandartukhtayev/imkon/api/docs" // for swagger
	v1 "github.com/samandartukhtayev/imkon/api/v1"
	"github.com/samandartukhtayev/imkon/config"
	"github.com/samandartukhtayev/imkon/storage"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type RouterOptions struct {
	Cfg      *config.Config
	Storage  storage.StorageI
	InMemory storage.InMemoryStorageI
}

// @title           Swagger for Imkon
// @version         1.0
// @description     This is a Imkon project api
// @BasePath  		/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "*")
	router.Use(cors.New(corsConfig))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:      opt.Cfg,
		Storage:  opt.Storage,
		InMemory: opt.InMemory,
	})
	router.Static("/media", "./media")
	apiV1 := router.Group("/v1")

	// User
	apiV1.POST("/users", handlerV1.CreateUser)
	apiV1.GET("/users/:id", handlerV1.GetUser)
	apiV1.GET("/users", handlerV1.GetAllUsers)
	apiV1.GET("/users-accepted", handlerV1.GetAllAcceptedUsers)
	apiV1.PUT("/users/:id", handlerV1.UpdateUser)
	apiV1.DELETE("/users/:id", handlerV1.DeleteUser)
	apiV1.POST("/users/portfolia-upload/:id", handlerV1.UsersPortfoliaUpload)
	apiV1.POST("/users/image-upload/:id", handlerV1.UsersImageUpload)

	// Business
	apiV1.POST("/businesses", handlerV1.CreateBusiness)
	apiV1.GET("/businesses/:id", handlerV1.GetBusiness)
	apiV1.GET("/businesses", handlerV1.GetAllBusinesses)
	apiV1.PUT("/businesses/:id", handlerV1.UpdateBusiness)
	apiV1.DELETE("/businesses/:id", handlerV1.DeleteBusiness)
	apiV1.POST("/businesses/image-upload/:id", handlerV1.BusinessesImageUpload)

	apiV1.POST("/vacancies", handlerV1.CreateVacancy)
	apiV1.GET("/vacancies/:id", handlerV1.GetVacancy)
	apiV1.GET("/vacancies", handlerV1.GetAllVacancies)
	apiV1.PUT("/vacancies/:id", handlerV1.UpdateVacancy)
	apiV1.DELETE("/vacancies/:id", handlerV1.DeleteVacancy)
	apiV1.POST("/vacancies/image-upload/:id", handlerV1.VacanciesImageUpload)

	apiV1.POST("/courses", handlerV1.CreateCourse)
	apiV1.GET("/courses/:id", handlerV1.GetCourse)
	apiV1.GET("/courses", handlerV1.GetAllCourses)
	apiV1.PUT("/courses/:id", handlerV1.UpdateCourse)
	apiV1.DELETE("/courses/:id", handlerV1.DeleteCourse)
	apiV1.POST("/courses/image-upload/:id", handlerV1.CoursesImageUpload)

	apiV1.POST("/accepted-vacancies", handlerV1.AcceptVacancy)
	apiV1.GET("/accepted-vacancies-by-user/:id", handlerV1.GetAcceptedVacanciesByUserId)
	apiV1.GET("/accepted-vacancies-by-business/:id", handlerV1.GetAcceptedVacanciesByBusinessId)
	apiV1.GET("/accepted-vacancies-by-id/:id", handlerV1.GetAcceptedVacanciesById)
	apiV1.DELETE("/accepted-vacancy/:id", handlerV1.DeleteAcceptedVacancy)

	apiV1.POST("/categories", handlerV1.CreateCategory)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
