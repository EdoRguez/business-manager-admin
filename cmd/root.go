package cmd

import (
	"os"

	"github.com/EdoRguez/business-manager-admin/cmd/company"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "business-manager-admin",
	Short: "Business Manager CLI to manage administrative things",
	Long:  `Business Manager CLI to manage administrative things`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(company.CompanyCmd)
}
