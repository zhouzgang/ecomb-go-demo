package cmd

import (
	"ecomb-go-demo/web/cfg"
	"ecomb-go-demo/web/di"
	"ecomb-go-demo/web/svc"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
)

var envFile string

var rootCmd = &cobra.Command{
	Use: "ecomb-go-demo",
	Run: func(cmd *cobra.Command, args []string) {
		svc.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func init() {
	// todo what is this do
	log.Logger = log.With().Caller().Logger()
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	cobra.OnInitialize(initConfig, initDI)
	rootCmd.PersistentFlags().StringVarP(&envFile, "config", "c", "./web/dist/cfg/local.yml", "env config file")
}

func initConfig() {
	cfg.Load(envFile)
}

func initDI() {
	var container = di.New()

	container.Init()

	di.SetGlobalDI(container)
}