package auth

import (
	"github.com/spf13/cobra"
	"github.com/yuchou87/rest-agent/pkg/ai"
)

var (
	backend        string
	password       string
	baseURL        string
	endpointName   string
	model          string
	engine         string
	temperature    float32
	providerRegion string
	providerId     string
	compartmentId  string
	topP           float32
	topK           int32
	maxTokens      int
	organizationId string
)

var configAI ai.Configuration

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with your chosen backend",
	Long:  `Provide the necessary credentials to authenticate with your chosen backend.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
	},
}

func init() {
	AuthCmd.AddCommand(addCmd)
	AuthCmd.AddCommand(defaultCmd)
	AuthCmd.AddCommand(updateCmd)
	AuthCmd.AddCommand(removeCmd)
}
