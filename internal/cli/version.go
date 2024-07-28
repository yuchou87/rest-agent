package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime/debug"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of rest-agent",
	Long:  `All software has versions. This is rest-agent's`,
	Run: func(cmd *cobra.Command, args []string) {
		if Version == "dev" {
			details, ok := debug.ReadBuildInfo()
			if ok && details.Main.Version != "" && details.Main.Version != "(devel)" {
				Version = details.Main.Version
				for _, i := range details.Settings {
					if i.Key == "vcs.time" {
						Date = i.Value
					}
					if i.Key == "vcs.revision" {
						Commit = i.Value
					}
				}
			}
		}
		fmt.Printf("rest-agent: %s (%s), built at: %s\n", Version, Commit, Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
