package main

import (
  "context"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/gin-gonic/gin"

  "app/internal/config"
)

func main() {
  _ = config.Load()

  r := gin.New()
  r.GET("/health", func(c *gin.Context) {
    c.JSON(200, gin.H{"status": "ok"})
  })

  srv := &http.Server{
    Addr:    ":8080",
    Handler: r,
  }

  go func() {
    _ = srv.ListenAndServe()
  }()

  quit := make(chan os.Signal, 1)
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  _ = srv.Shutdown(ctx)
}
