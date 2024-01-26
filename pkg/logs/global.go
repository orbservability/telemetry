package logs

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Sets up the global logging configuration.
func Init() error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := version(); err != nil {
		return err
	}
	if err := service(); err != nil {
		return err
	}
	if err := env(); err != nil {
		return err
	}

	log.Info().Msg("Logger configured!")

	return nil
}

func version() error {
	version := os.Getenv("VERSION")

	log.Logger = log.With().
		Str("version", version).
		Caller().
		Logger()

	return nil
}

func service() error {
	service := os.Getenv("SERVICE")

	if service == "" {
		return fmt.Errorf("SERVICE environment variable is missing")
	}

	log.Logger = log.With().
		Str("service", service).
		Logger()

	return nil
}

func env() error {
	env := os.Getenv("ENV")

	switch env {
	case "production":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "test", "development":
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	default:
		return fmt.Errorf("ENV set to '%s'. Must be set to 'test', 'development', or 'production'", env)
	}

	log.Logger = log.With().
		Str("env", env).
		Logger()

	return nil
}
