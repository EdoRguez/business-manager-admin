package user

import (
	"fmt"
	"time"

	"github.com/EdoRguez/business-manager-admin/pkg/db"
	"github.com/EdoRguez/business-manager-admin/pkg/models"
	"github.com/EdoRguez/business-manager-admin/pkg/util/password_hash"
	"github.com/spf13/cobra"
)

const (
	createAction = "create"
	editAction   = "edit"
)

var (
	action  = ""
	UserCmd = &cobra.Command{
		Use:   "user",
		Short: "User Features",
		Long:  `This subcommand starts user features`,
		Run: func(cmd *cobra.Command, args []string) {
			switch action {
			case createAction:
				CreateUser()
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
	UserCmd.Flags().StringVarP(&action, "action", "a", "", "Action  (create/edit)")
	UserCmd.MarkFlagRequired("action")
}

func CreateUser() {
	var model models.User
	model.CreatedAt = time.Now()
	model.ModifiedAt = time.Now()

	fmt.Println()
	fmt.Print(" > Company ID: ")
	_, _ = fmt.Scanln(&model.CompanyID)

	fmt.Print(" > Role ID (SuperAdmin (NO) = 1 / Admin = 2 / Regular = 3): ")
	_, _ = fmt.Scanln(&model.RoleID)

	fmt.Print(" > Email: ")
	_, _ = fmt.Scanln(&model.Email)

	fmt.Print(" > Password: ")
	_, _ = fmt.Scanln(&model.PasswordHash)
	model.PasswordHash = password_hash.HashPassword(model.PasswordHash)

	if err := db.CreateUser(&model); err != nil {
		fmt.Println("Error Creating User")
		fmt.Println(err.Error())
	} else {
		fmt.Printf("\n- User Created ! ID = %d", model.ID)
	}
}
