package cmd

import (
	"fmt"
	"github.com/enkhalifapro/order-packs-calculator/api"
	"github.com/enkhalifapro/order-packs-calculator/internal/packing"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

const (
	appName = "order-packs-calculator"
)

var (
	environment string
	port        string

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "run order-packs-calculator",
		Long:  "run order-packs-calculator",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.SetFormatter(&logrus.JSONFormatter{})
			logger := logrus.WithFields(
				logrus.Fields{
					"environment": environment,
					"service":     appName,
					"operation":   "run",
				},
			)

			// Api registration and routing
			healthAPI := api.NewHealthAPI(logger)
			router := httprouter.New()
			router.GET("/", healthAPI.Index)
			router.GET("/health", healthAPI.Health)

			// Packing API
			packingManager := packing.NewManager(logger)
			packingAPI := api.NewPackingAPI(logger, packingManager)
			router.POST("/packing/calculate", packingAPI.Calculate)
			logger.Infof("%s app started", appName)
			r := cors.AllowAll().Handler(router)
			logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	viper.AutomaticEnv()
	getEnvVars()
	rootCmd.AddCommand(runCmd)
}

func getEnvVars() {
	environment = "dev"
	if viper.GetString("ENV") != "" {
		environment = viper.GetString("ENV")
	}
	port = "8090"
	if viper.GetString("PORT") != "" {
		port = viper.GetString("PORT")
	}
}
