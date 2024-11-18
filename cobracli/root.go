package cobracli

import (
	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:   "CRUD",
	Short: "My workbench for user repository with redis and mysql ",
}
