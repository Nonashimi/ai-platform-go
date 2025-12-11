package main

import (
	"log/slog"
	"net/http"
	"os"
	"project-go/internal/config"
	chatCreate "project-go/internal/http-server/handlers/chat/create"
	getbysessionid "project-go/internal/http-server/handlers/chat/getBySessionId"
	getAllSession "project-go/internal/http-server/handlers/session/getAll"
	test_create "project-go/internal/http-server/handlers/test/test"
	userCreate "project-go/internal/http-server/handlers/user/create"
	"project-go/internal/http-server/handlers/user/login"
	"project-go/internal/http-server/middleware/auth"
	"project-go/internal/http-server/middleware/logger"
	"project-go/internal/http-server/repository/store"
	"project-go/internal/lib/jwt"
	"project-go/internal/lib/logger/handlers/slogpretty"
	"project-go/internal/lib/logger/sl"
	"project-go/internal/storage/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log = log.With("env", cfg.Env)
	// log.Info("starting our project")
	// log.Debug("debug messages are enabled")
	// log.Error("error messages are enabled")
	db, err := postgres.New(cfg.Dsn)
	jwtService := jwt.NewJWTService(cfg.JWT_Key)
	authMiddleware := auth.AuthMiddleware([]byte(cfg.JWT_Key))
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	// router.Route("/url", func(r chi.Router) {
	// 	r.Use(middleware.BasicAuth("url-shortener", map[string]string{
	// 		cfg.HTTPServer.User: cfg.HTTPServer.Password,
	// 	}))
	// 	r.Post("/", save.New(log, storage, storage))
	// 	r.Delete("/{alias}", delete.New(log, storage))
	// })
	// router.Get("/{alias}", redirect.New(log, storage))
	store := store.NewStore(db)
	router.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Post("/chat", chatCreate.New(log, store.ChatRepo, store.SessionRepo))
		r.Get("/sessions", getAllSession.New(log, store.SessionRepo))
		r.Get("/sessions/{sessionId}", getbysessionid.New(log, store.ChatRepo))
		r.Post("/test", test_create.New(log, store.TestRepo))
	})
	router.Post("/register", userCreate.New(log, store.UserRepo))
	router.Post("/login", login.New(log, store.UserRepo, jwtService))
	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start and run server")
	}

	log.Error("server stopped")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
