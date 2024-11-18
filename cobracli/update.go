package cobracli

import (
	"fmt"
	"learn_1/repository"
	"log"

	"github.com/spf13/cobra"
)

func CreateUpdateCommand(userrepo repository.UserRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "update [id] [newfirstname] [newlastname]",
		Short: "update a created user",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			newfirstname := args[1]
			newlastname := args[2]
			id := args[0]
			err := userrepo.UpdateUser(id, newfirstname, newlastname)
			if err != nil {
				log.Fatalf("Error Updating user: %v", err)
			}
			fmt.Printf("Updated user !\n")
		},
	}
}
