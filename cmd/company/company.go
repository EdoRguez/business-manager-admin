package company

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	action     = ""
	CompanyCmd = &cobra.Command{
		Use:   "company",
		Short: "Company Features",
		Long:  `This subcommand starts company features`,
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println()
			fmt.Print(" > Company name: ")
			name, _ := reader.ReadString('\n')

			fmt.Print(" > Plan ID (Basic = 1 / Pro = 2): ")
			planId, _ := reader.ReadString('\n')

			fmt.Print(" > Last Payment Date (YYYY-MM-DD): ")
			paymentDate, _ := reader.ReadString('\n')

			fmt.Printf("Hello %s, you are %s years old, %s.\n", name, planId, paymentDate)
		},
	}
)

func init() {
	// rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	CompanyCmd.Flags().StringVarP(&action, "action", "a", "", "Action  (create/edit)")
	CompanyCmd.MarkFlagRequired("action")
}
