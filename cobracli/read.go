package cobracli

import (
	"learn_1/repository"

	"github.com/spf13/cobra"
)

func CreateReadCommand(userrepo repository.UserRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "Read [id] ",
		Short: "Read a created user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			userrepo.Read(id)
		},
	}
}

func CreateReadallCommand(userrepo repository.UserRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "Read_all",
		Short: "Read all users",
		Run: func(cmd *cobra.Command, args []string) {
			userrepo.Read_all()
		},
	}
}
