package main

import (
	"os"

	"github.com/devsendjin/url-shortener/internal/config"
	mwLogger "github.com/devsendjin/url-shortener/internal/http-server/middleware/logger"
	"github.com/devsendjin/url-shortener/internal/lib/logger/sl"
	"github.com/devsendjin/url-shortener/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

func main() {
	cfg := config.MustLoad()

	logger := sl.New(cfg.Env)

	logger.Info("Starting url-shortener...", slog.String("env", cfg.Env))
	logger.Debug("Debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		logger.Error("Failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// TODO: run server
}

// fmt.Printf("%+v\n", cfg)
// func prettyPrint(data interface{}) {
// 	var p []byte
// 	p, err := json.MarshalIndent(data, "", "\t")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("%s \n", p)
// }
