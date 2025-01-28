package company

import (
	"fmt"
	"time"

	"github.com/EdoRguez/business-manager-admin/pkg/db"
	"github.com/EdoRguez/business-manager-admin/pkg/models"
	"github.com/spf13/cobra"
)

const (
	createAction = "create"
	editAction   = "edit"
)

var (
	action     = ""
	CompanyCmd = &cobra.Command{
		Use:   "company",
		Short: "Company Features",
		Long:  `This subcommand starts company features`,
		Run: func(cmd *cobra.Command, args []string) {
			switch action {
			case createAction:
				CreateCompany()
			}
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

func CreateCompany() {
	var model models.Company
	model.CreatedAt = time.Now()
	model.ModifiedAt = time.Now()

	fmt.Println()
	fmt.Print(" > Company name: ")
	_, _ = fmt.Scanf("%s", &model.Name)

	fmt.Print(" > Name Format URL: ")
	_, _ = fmt.Scanf("%s", &model.NameFormatUrl)

	fmt.Print(" > Plan ID (Basic = 1 / Pro = 2): ")
	_, _ = fmt.Scanf("%d", &model.PlanID)

	fmt.Print(" > Last Payment Date (YYYY-MM-DD): ")
	var lastPaymentDateInput string
	_, _ = fmt.Scanln(&lastPaymentDateInput)
	model.LastPaymentDate, _ = time.Parse("2006-01-02", lastPaymentDateInput)

	if err := db.CreateCompany(&model); err != nil {
		fmt.Println("Error Creating Company")
		fmt.Println(err.Error())
	} else {
		fmt.Printf("\n- Company Created ! ID = %d", model.ID)
	}
}
