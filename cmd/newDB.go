package cmd

import (
	"github.com/spf13/cobra"
)

func NewNewDBCmd(props *Props) *cobra.Command {
	return &cobra.Command{
		Use:   "new-db [name]",
		Short: "create new DB",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := props.storage.NewDB(args[0]); err != nil {
				props.logger.Fatal(err)
			}
		},
	}
}
