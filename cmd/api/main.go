package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/natefinch/lumberjack"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"jax.hoangdv99/internal/data"
	"jax.hoangdv99/internal/mailer"
)

var (
	version string
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type application struct {
	config config
	logger zerolog.Logger
	models data.Models
	wg     sync.WaitGroup
	mailer mailer.Mailer
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("error loading .env file")
	}

	var cfg config
	cfg.port = getEnvAsInt("PORT", 4000)
	cfg.env = getEnv("ENVIRONMENT", "development")
	cfg.db.dsn = fmt.Sprintf("%s:%s@/%s?parseTime=true", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	cfg.db.maxOpenConns = getEnvAsInt("DB_MAX_OPEN_CONNS", 25)
	cfg.db.maxIdleConns = getEnvAsInt("DB_MAX_IDLE_CONNS", 25)
	cfg.db.maxIdleTime = getEnv("DB_MAX_IDLE_TIME", "15m")

	cfg.smtp.host = getEnv("SMTP_HOST", "")
	cfg.smtp.port = getEnvAsInt("SMTP_PORT", 2525)
	cfg.smtp.username = getEnv("SMTP_USERNAME", "")
	cfg.smtp.password = getEnv("SMTP_PASSWORD", "")
	cfg.smtp.sender = getEnv("SMTP_SENDER", "")

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("error opening database connection")
	}
	defer db.Close()

	log.Info().Msg("database connection established")

	app := &application{
		config: cfg,
		logger: initLogger(),
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	err = app.serve()
	if err != nil {
		log.Fatal().Err(err).Msg("error starting server")
	}
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.config.port), Handler: app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		app.logger.Info().Str("signal", s.String()).Msg("shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		app.logger.Info().Str("addr", srv.Addr).Msg("completing background tasks")

		app.wg.Wait()
		shutdownError <- nil
	}()

	app.logger.Info().Str("addr", srv.Addr).Str("env", app.config.env).Msg("starting server")

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	app.logger.Info().Str("addr", srv.Addr).Msg("stopped server")

	return nil
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initLogger() zerolog.Logger {
	logFileName := "tmp/logs/" + time.Now().Format("2006-01-02") + ".log"

	if _, err := os.Stat("tmp/logs"); os.IsNotExist(err) {
		_ = os.Mkdir("tmp/logs", 0755)
	}

	logFile := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    10,
		MaxBackups: 7,
		MaxAge:     30,
		Compress:   true,
	}

	multi := io.MultiWriter(os.Stdout, logFile)

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := zerolog.New(multi)

	return logger
}
