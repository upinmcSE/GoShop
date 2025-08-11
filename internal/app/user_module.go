package app

import (
	v1handler "github.com/upinmcSE/goshop/internal/handler/v1"
	"github.com/upinmcSE/goshop/internal/repository"
	"github.com/upinmcSE/goshop/internal/routes"
	v1routes "github.com/upinmcSE/goshop/internal/routes/v1"
	v1service "github.com/upinmcSE/goshop/internal/service/v1"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule() *UserModule {
	userRepo := repository.NewSqlUserRepository()
	userService := v1service.NewUserService(userRepo)
	userHandler := v1handler.NewUserHandler(userService)
	userRoutes := v1routes.NewUserRoutes(userHandler)
	return &UserModule{routes: userRoutes}
}

func (m *UserModule) Routes() routes.Route {
	return m.routes
}
