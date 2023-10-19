package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/notblessy/config"
	"github.com/notblessy/db"
	"github.com/notblessy/delivery/http"
	"github.com/notblessy/internal/employee"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"

	"github.com/spf13/cobra"
)

var runHTTPServer = &cobra.Command{
	Use:   "httpsrv",
	Short: "run http server",
	Long:  `This subcommand is for starting the http server`,
	Run:   runHTTP,
}

func init() {
	rootCmd.AddCommand(runHTTPServer)
	config.LoadConfig()
}

func runHTTP(cmd *cobra.Command, args []string) {
	psql := db.NewPostgreSQL()
	trace := trace.NewNoopTracerProvider().Tracer("test")
	e := echo.New()

	employeeTransport := employee.NewTransport(psql, trace)

	httpSvc := http.NewHTTPService()
	httpSvc.RegisterEmployeeTransport(employeeTransport)

	httpSvc.Router(e)

	fmt.Println(config.HTTPPort())
	logrus.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTPPort())))
}
