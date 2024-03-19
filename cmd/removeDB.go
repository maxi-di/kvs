package cmd

import "github.com/spf13/cobra"

func NewRemoveDBCmd(props *Props) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-db [name]",
		Short: "remove DB",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := props.storage.RemoveDB(args[0]); err != nil {
				props.logger.Fatal(err)
			}
		},
	}
}
