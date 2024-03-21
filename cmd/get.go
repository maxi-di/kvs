package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewGetCmd(props *Props) *cobra.Command {
	dbName := ""
	key := ""
	with_add := false
	c := &cobra.Command{
		Use:   "get",
		Short: "get Value from DB by Key",
		Run: func(cmd *cobra.Command, args []string) {
			err := openDB(props.storage, dbName)
			if err != nil {
				props.logger.Fatal(err)
			}
			props.logger.Infof("getting from db '%s'", dbName)

			if key == "" {
				keys := props.storage.GetKeys()
				key, _, err = fuzzy(keys, "Choose value")
				if err != nil {
					props.logger.Fatal(err)
				}
			}
			if key == "" {
				props.logger.Warn("key not specified")
				return
			}

			value, err := props.storage.GetValue(key)
			if err != nil {
				props.logger.Fatal(err)
			}
			if value == "" {
				props.logger.Warnf("no value for key '%s'", key)
				return
			}
			fmt.Println(value)
		},
	}
	c.Flags().StringVar(&dbName, "db", "", "db name (with ext)")
	c.Flags().StringVar(&key, "key", "", "key name")
	c.Flags().BoolVar(&with_add, "with-add", false, "append new entry in interactive mode")

	c.MarkFlagsMutuallyExclusive("with-add", "key")

	return c
}
