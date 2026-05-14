package provider

import (
	"github.com/RakaMurdiarta/online-shop-system/internal/config"
	"github.com/RakaMurdiarta/online-shop-system/internal/middlewares"
	"github.com/RakaMurdiarta/online-shop-system/internal/modules/users/handlers"
	"github.com/RakaMurdiarta/online-shop-system/internal/modules/users/repository"

	ImplService "github.com/RakaMurdiarta/online-shop-system/internal/modules/users/services/impl"
	"github.com/RakaMurdiarta/online-shop-system/pkg/database"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

func UserProvider(
	privateRoute *echo.Group,
	tx *database.TransactionManagerImpl,
	conf *config.Config,
	userRepo repository.UserRepository,

) {
	v := validator.New()

	userAddressService := ImplService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userAddressService, v)

	users := privateRoute.Group("/users")
	users.POST("", userHandler.CreateUser, middlewares.IsAdmin)
	users.GET("", userHandler.ListUsers, middlewares.IsAdmin)
	users.GET("/:id", userHandler.GetUserByID, middlewares.IsAdmin)
	users.DELETE("/:id", userHandler.DeleteUser, middlewares.IsAdmin)
}
