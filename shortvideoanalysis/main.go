package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webtemp/shortvideoanalysis/db"
	"webtemp/shortvideoanalysis/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	// 设置ShortVideoCollection（顺便初始化了dbManager，db包下的init）
	db.DbManager.SetShortVideoCollection("short_video_url")
}

func main() {
	router := gin.Default()

	// 加载模板
	router.LoadHTMLGlob("templates/**/*")
	// 设置favicon
	// router.StaticFile("/favicon.ico", "xxxpath")
	// 设置路由
	routers.SetupRouter(router)

	// 8082端口启动
	// router.Run(":8082")

	srv := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 设置优雅退出
	gracefulExitWeb(srv)
}

func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-ch
	log.Println("Shutdown Server ...")

	// 关闭数据库连接
	db.DbManager.CloseDB()

	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")

	// 看看实际退出所耗费的时间
	fmt.Println("------exited--------", time.Since(now))
}
