package cobracli

import (
	"fmt"
	"learn_1/repository"
	"log"

	"github.com/spf13/cobra"
)

func CreateDeleteCommand(userrepo repository.UserRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id] ",
		Short: "delete a created user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			err := userrepo.DeleteUser(id)
			if err != nil {
				log.Fatalf("Error deleting user: %v", err)
			}
			fmt.Printf("deleted user : %v\n", id)
		},
	}
}
