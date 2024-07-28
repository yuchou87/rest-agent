package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yuchou87/rest-agent/internal/cli/auth"
	"github.com/yuchou87/rest-agent/internal/cli/gen"
	"github.com/yuchou87/rest-agent/pkg/utils"
	"os"
	"path/filepath"
)

var (
	cfgFile string
	Version string
	Commit  string
	Date    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rest-agent",
	Short: "Rest debugging client powered by AI",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(v string, c string, d string) {
	Version = v
	Commit = c
	Date = d
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	checkFileAndDir()
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.AddCommand(gen.GenerateCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("Default config file (%s)", getConfigFilePath()))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// the config will belocated under `~/.config/rest-agent/.config.yaml` on linux
		configDir := utils.GetConfigDir()

		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".config")

		_ = viper.SafeWriteConfig()
	}

	viper.SetEnvPrefix("REST_AGENT")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_ = 1
		//	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func getConfigFilePath() string {
	return filepath.Join(utils.GetConfigDir(), ".config.yaml")
}

func checkFileAndDir() {
	config := getConfigFilePath()
	_, err := utils.FileExists(config)
	cobra.CheckErr(err)

	configDir := filepath.Dir(config)
	err = utils.EnsureDirExists(configDir)
	cobra.CheckErr(err)
}
