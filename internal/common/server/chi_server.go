package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fvbock/endless"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

// 配合 oapi-codegen 使用, openapi 定义接口,然后使用 oapi-codegen 一键生成
// 示例:
// oapi-codegen -generate types -o (项目目录下)ports/openapi_types.gen.go -package ports (openapi接口文件)api/openapi/output.admin.json
// oapi-codegen -generate chi-server -o (项目目录下)ports/openapi_api.gen.go -package ports (openapi接口文件)api/openapi/output.admin.json
// 	server.RunHTTPServer("/", func(router chi.Router) http.Handler {
//		router.Use(ports.AuthMiddleware()) // 增加自定义中间件
//		return ports.HandlerFromMux(app.HttpServer, router)
//	})

// RunHTTPServer 启动
func RunHTTPServer(baseUrl string, createHandler func(router chi.Router) http.Handler) {

	RunHTTPServerOnAddr(fmt.Sprintf(":%s", os.Getenv("PORT")), baseUrl, createHandler)
}

func RunHTTPServerOnAddr(addr, baseUrl string, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	rootRouter.Mount(baseUrl, createHandler(apiRouter))

	logrus.Info("Starting HTTP server")

	err := endless.ListenAndServe(addr, rootRouter)
	if err != nil {
		logrus.Fatal(`cannot start admin server : `, err)
	}
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.NoCache)
}
