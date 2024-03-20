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
			dbName = selectDB(dbName, props.storage.ListDB(), props.logger)
			if dbName == "" {
				return
			}
			props.logger.Infof("getting from bd '%s'", dbName)

			if key == "" {
				keys, err := props.storage.GetKeys(dbName)
				if err != nil {
					props.logger.Fatal(err)
				}
				key, _, err = fuzzy(keys, "Выберите значение")
				if err != nil {
					props.logger.Fatal(err)
				}
			}
			if key == "" {
				props.logger.Warn("key not specified")
				return
			}

			value, err := props.storage.GetValue(dbName, key)
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
