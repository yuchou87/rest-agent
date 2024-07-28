package auth

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove provider(s)",
	Long:  "The command to remove AI backend provider(s)",
	PreRun: func(cmd *cobra.Command, args []string) {
		_ = cmd.MarkFlagRequired("backends")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if backend == "" {
			color.Red("Error: backends must be set.")
			_ = cmd.Help()
			return
		}
		inputBackends := strings.Split(backend, ",")

		err := viper.UnmarshalKey("ai", &configAI)
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		for _, b := range inputBackends {
			foundBackend := false
			for i, provider := range configAI.Providers {
				if b == provider.Name {
					foundBackend = true
					configAI.Providers = append(configAI.Providers[:i], configAI.Providers[i+1:]...)
					if configAI.DefaultProvider == b {
						configAI.DefaultProvider = "openai"
					}
					color.Green("%s deleted from the AI backend provider list", b)
					break
				}
			}
			if !foundBackend {
				color.Red("Error: %s does not exist in configuration file. Please use k8sgpt auth new.", b)
				os.Exit(1)
			}
		}

		viper.Set("ai", configAI)
		if err := viper.WriteConfig(); err != nil {
			color.Red("Error writing config file: %s", err.Error())
			os.Exit(1)
		}

	},
}

func init() {
	// add flag for backends
	removeCmd.Flags().StringVarP(&backend, "backends", "b", "", "Backend AI providers to remove (separated by a comma)")
}
