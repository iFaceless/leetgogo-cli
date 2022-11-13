package command

import (
	"fmt"
	"github.com/iFaceless/leetgogo-cli/pkg/workdir"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	initCmd := &cobra.Command{
		Use:   "init [path to workdir]",
		Short: "initialize your workdir, code templates will be generated in your workdir",
		Args:  cobra.MaximumNArgs(1),
		Run:   handleInitCommand,
	}
	rootCmd.AddCommand(initCmd)
}

func handleInitCommand(_ *cobra.Command, args []string) {
	_workdir := "."
	if len(args) == 1 {
		_workdir = args[0]
	}

	if _, err := os.Stat(_workdir); os.IsNotExist(err) {
		log.Fatalln(fmt.Sprintf("workdir `%s` is not exist", _workdir))
	}

	wd := workdir.New(_workdir)
	err := wd.Init()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("workdir initialized")
}
