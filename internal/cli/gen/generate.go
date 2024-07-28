package gen

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/yuchou87/rest-agent/pkg/generate"
	"github.com/yuchou87/rest-agent/pkg/models"
	"github.com/yuchou87/rest-agent/pkg/utils"
	"log/slog"
	"os"
)

var (
	swaggerFile string
	outputType  string
)

var GenerateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate a test cases from a swagger file",
	Long:    "Generate a test cases from a swagger file",
	Run: func(cmd *cobra.Command, args []string) {
		isExist, _ := utils.FileExists(swaggerFile)
		if !isExist {
			color.Red("File is not exist, %s", swaggerFile)
			os.Exit(1)
		}

		content, err := os.ReadFile(swaggerFile)
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		vars := models.TestCaseGenerationPrompt{
			SwaggerFile: string(content),
			OutputType:  outputType,
		}

		var testCases models.TestCases

		color.Green("Test case generation started, just wait a moment")
		g, err := generate.NewGenerate(
			context.Background(),
			"test_case_generation_prompt",
			vars,
			outputType,
			&testCases,
		)

		if err != nil {
			slog.Error("Error: ", err)
			os.Exit(1)
		}

		if err := g.Generate(); err != nil {
			slog.Error("Error: ", err)
			os.Exit(1)
		}

		fmt.Printf("Results: %s\n", g.GetOutputResults())
		fmt.Printf("Results saved to: %s\n", g.GetOutputFile())
		color.Green("Test case generation completed")
	},
}

func init() {
	GenerateCmd.Flags().StringVarP(&swaggerFile, "swagger", "f", "", "Swagger file to generate")
	GenerateCmd.Flags().StringVarP(&outputType, "type", "t", "json", "output format for generate ")
}
