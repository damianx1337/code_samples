package main

import (
	// STD LIBS
	"os"
	"os/signal"
	"syscall"
	"bufio"
	"log"
	"context"
	"net/http"
	"time"

	// CUSTOM LIBS / MODULES
	"gin-skeleton/database"
	"gin-skeleton/responsecache"
	"gin-skeleton/models"
	"gin-skeleton/controllers"

	// GIN GONIC
	"github.com/gin-gonic/gin"

	// SWAGGER
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "<module>/docs"

	// METRICS - PROMETHEUS
	"github.com/prometheus/client_golang/prometheus/promhttp"

	// SECURITY - HEADERS
	//"github.com/gin-contrib/secure"
	// there is also gin-helm...

	// CACHE
	"github.com/chenyahui/gin-cache"
)

const (
	appHost = "0.0.0.0"
	appPort = "8080"
)

var (
	username = ""
	password = ""
)


// @title This is the title
// @version 1.0
// @description This is the description
// @contact.name Support
// @contact.mail @gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host <FQDN or IP>
// @BasePath /api
// @query.collection.format multi
func main () {
	file, err := os.Open("passwd")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		username = scanner.Text()
	}

	// INIT i.e., database and cache connection
	database.Init()
	responseCacheConfig := responsecache.Init()

	// INIT GIN ROUTER
	router := gin.Default()
	//router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// SECURITY HEADERS
	/*
	router.Use(secure.New(secure.Config{
		AllowedHosts: []string{""},
		SSLRedirect: false,
		SSLHost: "",
		STSSeconds: 315360000,
		STSIncludeSubdomains: true,
		FrameDeny: true,
		ContentTypeNosniff: true,
		BrowserXssFilter: true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen: true,
		ReferrerPolicy: "strict-origin-when-cross-origin",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	}))
	*/

	// HEALTHCHECK
	health := router.Group("/healthz")
	health.GET("/livez", controllers.Livez)
	health.GET("/ready", controllers.Ready)

	// METRICS
	metrics := router.Group("/metricz")
	metrics.GET("/prometheus", gin.WrapH(promhttp.Handler()))


	// API DOCS
	apiDocsPath := router.Group("/api/docs", gin.BasicAuth(gin.Accounts{
		username : password,
	}))
	apiDocsPath.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	// API V1
	apiV1 := router.Group("/api/v1", gin.BasicAuth(gin.Accounts{
		username : password,
	}))
	apiV1.GET("/", controllers.)
	// cache response
	apiV1.GET("/getAll", cache.CacheByRequestURI(cacheConfig.Store, cacheConfig.DefaultCacheTime), controllers.)
	apiV1.POST("/", controllers.)

	// START SVR (graceful shutdown)
	//router.Run(appHost + ":" + appPort) // original gin startup routine
	svr := &http.Server{
		Addr: appHost + ":" + appPort,
		Handler: router,
	}

	// Initializing server in dedicated go routine - no block graceful shutdown
	go func() {
		if err := svr.ListenAndServeTLS("server.crt","server.key"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// wait for interrupt signal to gracefully shutdown
	// server with timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// context is used to inform the server it has 5 seconds
	// to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting.")
}
