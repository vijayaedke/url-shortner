package cmd

import (
	"fmt"
	"os"

	"url-shortner/internal/config"
	"url-shortner/internal/wire"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "url-shortner",
	Short: "URL shortner generates and stores long url",
	Long: `A URL shortner which takes a long URL and generates a
			short URL. The short URL is not more than 16 characters`,
	Run: func(cmd *cobra.Command, args []string) {
		urlShortenerServer, _, err := wire.InitializeServer()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't initialize the server. Got this error: %v", err)
			os.Exit(-1)
		}

		urlShortenerServer.StartServer()
	},
}

func init() {
	viper.BindEnv("redis.password", "REDIS_PASSWORD")
	viper.BindEnv("redis.host", "REDIS_HOST")
	viper.BindEnv("redis.port", "REDIS_PORT")
	viper.BindEnv("redis.database", "REDIS_DATABASE")
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config.InitConfig(cfgFile)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("url shortner app init failed", err)
	}
}
