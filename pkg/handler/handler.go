package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/Tom-Challenger/go-todo/pkg/service"

	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/Tom-Challenger/go-todo/docs"
)

type Handler struct {
	service *service.Service
}

func NewHendler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		// URI для регистрации
		// /auth/sing-up
		auth.POST("/sign-up", h.singUp)
		// Маршрут для авторизации
		auth.POST("/sign-in", h.singIn)
	}

	api := router.Group("/api", h.userIdentify)
	{
		lists := api.Group("/lists")
		{
			// URI для получения всех списков
			// /api/lists/
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group("/:id/items")
			{
				// Получение всех задач списка
				// /api/lists/1/items/
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
