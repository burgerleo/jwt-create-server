package main

import (
	"flag"
	"fmt"
	"jwt-generate-server/conf"
	"jwt-generate-server/router"
	"jwt-generate-server/service/token"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

var (
	confPath = flag.String("config", "./conf/app.dev.ini", "config location")
	confInfo *conf.Config
)

func Init() error {
	flag.Parse()

	// read config and pass variables ...
	var err error
	confInfo, err = conf.InitConfig(confPath)
	if err != nil {
		return fmt.Errorf("Init config err: %v", err)
	}

	err = token.InitToken(confInfo.JwtConf.JwtSecret, confInfo.JwtConf.JwtExpiredMinute)
	if err != nil {
		return fmt.Errorf("Init JWT err: %v", err)
	}
	return nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic err: %v", err)
		}
	}()

	err := Init()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	route := setupRouter()

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%v", confInfo.BaseConf.HttpPort),
		Handler: route,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(fmt.Sprintf("http listen : %v\n", err))
			panic(err)
		}
	}()

	gracefulShutdown()
}

func setupRouter() *gin.Engine {
	route := gin.Default()
	// 讓路徑大小寫通吃
	route.RedirectFixedPath = true
	router.ApiRouter(route)

	return route
}

// gracefulShutdown: handle the worker connection
func gracefulShutdown() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
