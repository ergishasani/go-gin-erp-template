package config

import (
  "log"
  "os"

  "github.com/joho/godotenv"
)

type Config struct {
  Env       string
  Port      string
  DBUrl     string
  JWTSecret string
}

func Load() Config {
  _ = godotenv.Load()

  cfg := Config{
    Env:       get("APP_ENV", "dev"),
    Port:      get("PORT", "8080"),
    DBUrl:     must("DATABASE_URL"),
    JWTSecret: must("JWT_SECRET"),
  }

  return cfg
}

func get(k, d string) string {
  if v := os.Getenv(k); v != "" {
    return v
  }
  return d
}

func must(k string) string {
  if v := os.Getenv(k); v != "" {
    return v
  }
  log.Fatalf("missing env %s", k)
  return ""
}
