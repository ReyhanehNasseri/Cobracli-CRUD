package main

import (
	"fmt"
	"learn_1/cobracli"
	"learn_1/repository"
	"log"
	"os"
)

func main() {
	userrepo := new(repository.CustomUserRepository)
	if connectionErr := userrepo.Connect(); connectionErr != nil {
		log.Fatal(connectionErr)
	}
	cobracli.Root.AddCommand(cobracli.CreateInsertCommand(userrepo))
	cobracli.Root.AddCommand(cobracli.CreateReadCommand(userrepo))
	cobracli.Root.AddCommand(cobracli.CreateReadallCommand(userrepo))
	cobracli.Root.AddCommand(cobracli.CreateDeleteCommand(userrepo))
	cobracli.Root.AddCommand(cobracli.CreateUpdateCommand(userrepo))
	execute()
}

func execute() {
	if err := cobracli.Root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
