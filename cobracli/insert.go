package cobracli

import (
	"fmt"
	"learn_1/model"
	"learn_1/repository"
	"log"

	"github.com/spf13/cobra"
)

func CreateInsertCommand(usercrud repository.UserRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "insert [id] [firstname] [lastname]",
		Short: "Insert a new user",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			firstname := args[1]
			lastname := args[2]
			id := args[0]

			newuser := model.User{}
			newuser.ID = id
			newuser.Firstname = firstname
			newuser.Lastname = lastname
			saveduser, err := (usercrud).InsertUser(newuser)
			if err != nil {
				log.Fatalf("Error inserting user: %v", err)
			}
			fmt.Printf("Inserted user : %v\n", saveduser)
		},
	}
}
