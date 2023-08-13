package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func Ping(ctx *gin.Context) {
    ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    ctx.String(200, "pong")
}

func main() {
    gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Static("/static", "static")
	router.LoadHTMLGlob("index.html")
	router.GET("/", Home)
    router.GET("/ping", Ping)
    server := &http.Server{Addr: ":8080", Handler: router}
    go func() {
    if err := server.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatal(err.Error())}
    } ()
    var sig = make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT, os.Interrupt)
	<-sig
	err := server.Shutdown(context.Background())
	if err != nil {
        log.Fatalf("HTTP shutdown error: %v", err)
    }
    log.Print("Bye!")
 }