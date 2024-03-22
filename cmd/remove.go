package cmd

import (
	"github.com/spf13/cobra"
)

func NewRemoveCmd(props *Props) *cobra.Command {
	dbName := ""
	key := ""
	c := &cobra.Command{
		Use:     "remove",
		Aliases: []string{"delete"},
		Short:   "remove record from DB",
		Run: func(cmd *cobra.Command, args []string) {
			err := openDB(props.storage, dbName)
			if err != nil {
				props.logger.Fatal(err)
			}
			props.logger.Infof("removing from db '%s'", dbName)

			if key == "" {
				keys := props.storage.GetKeys()
				key, _, err = fuzzy(keys, "Select for delete...")
				if err != nil {
					props.logger.Fatal(err)
				}
			}
			if key == "" {
				return
			}

			err = props.storage.RemoveValue(key)
			if err != nil {
				props.logger.Fatal(err)
			}
		},
	}
	c.Flags().StringVar(&dbName, "db", "", "db name (with ext)")
	c.Flags().StringVar(&key, "key", "", "key name")

	return c
}
