package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/upinmcSE/goshop/internal/config"
	"github.com/upinmcSE/goshop/internal/routes"
	"github.com/upinmcSE/goshop/internal/validation"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config *config.Config
	router *gin.Engine
	modules []Module
}

func NewApplication(cfg *config.Config) *Application {
	if err := validation.InitValidator(); err != nil {
		log.Fatalf("Validator init failed %v", err)
	}

	loadEnv()

	r := gin.Default()

	modules := []Module{
		NewUserModule(),
	}

	routes.RegisterRoutes(r, getModulRoutes(modules)...)

	return &Application{
		config: cfg,
		router: r,
		modules: modules,
	}
}

func (a *Application) Run() error {
	// return a.router.Run(a.config.ServerAddress)
	srv := &http.Server{
		Addr: a.config.ServerAddress,
		Handler: a.router,
	}

	quit := make(chan os.Signal, 1)
	// syscall.SIGINT -> Ctrl + C
	// syscall.SIGTERM -> Kill service
	// syscall.SIGHUP -> Reaload service
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func(){
		if err := srv.ListenAndServe(); err != http.ErrServerClosed{
			log.Fatalf("Faild to start server: %v", err)
		}
	}()

	// Lệnh này sẽ block cho đến khi nhận được tín hiệu
	// Đây là channel receive operation trong Go.
	// Nó sẽ chặn (block) chương trình lại cho đến khi channel quit có ít nhất 1 giá trị được gửi vào.
	// Khi có giá trị, nó sẽ lấy giá trị đó ra và tiếp tục chạy code tiếp theo.
	<- quit

	ctx, cancel :=context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	return nil
}

func getModulRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}

func loadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found")
	}
}
