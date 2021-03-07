package core

import (
	"encoding/json"
	"log"
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// AppHandler is used to have a custom signature for the endpoint handlers
type AppHandler func(http.ResponseWriter, *http.Request) *ServerError

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		logger, e := createLogger()
		if e != nil {
			log.Fatal("Unable to create logger")
		}
		defer logger.Sync()

		logger.Error(err.Message, zap.String("error", err.Error.Error()))

		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
	}
}

func createLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return config.Build()
}
