package cmd

import (
	"github.com/byyjoww/leaderboard/config"
	"github.com/byyjoww/leaderboard/logging"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts a new server",
		Run: func(cmd *cobra.Command, _ []string) {
			var (
				cfg    = config.Build()
				logger = logging.NewLoggr(cfg.Logging)
			)

			logger.Info("Initializing app")

			switch startFlags.serverType {
			case serverTypeApi:
				startAPI(logger, cfg)
			}
		},
	}

	startFlags = struct {
		serverType string
	}{}

	serverTypeApi string = "api"
)

func init() {
	RootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&startFlags.serverType, "type", serverTypeApi, "The server type to initialize")
}
