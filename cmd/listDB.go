package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewListDBCmd(props *Props) *cobra.Command {
	return &cobra.Command{
		Use:   "list-db",
		Short: "list of all DB's files",
		Run: func(cmd *cobra.Command, args []string) {
			for _, v := range props.storage.ListDB() {
				fmt.Println(v)
			}
		},
	}
}
