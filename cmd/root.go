package cmd

import (
	"os"

	"github.com/notblessy/config"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "app root",
	Long:  `root for running console`,
}

type OpenTelemetryHook struct {
	tracer trace.Tracer
}

func newOpenTelemetryHook() *OpenTelemetryHook {
	tracer := otel.Tracer("logrus-hook")
	return &OpenTelemetryHook{
		tracer: tracer,
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setupLogger()
}

func setupLogger() {
	formatter := log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceColors:     true,
	}

	if config.Env() == "development" {
		formatter = log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceColors:     false,
		}
	}

	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)

	logLevel, err := log.ParseLevel(config.LogLevel())
	if err != nil {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)
}

func (h *OpenTelemetryHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *OpenTelemetryHook) Fire(entry *logrus.Entry) error {
	_, span := h.tracer.Start(entry.Dup().Context, entry.Caller.Function)
	defer span.End()

	for key, value := range entry.Data {
		var v string
		if str, ok := value.(string); ok {
			v = str
		}

		span.SetAttributes(attribute.String(key, v))
	}

	span.AddEvent(entry.Message)

	return nil
}
