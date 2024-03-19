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
			var err error
			if dbName == "" {
				dbName, err = fuzzy(props.storage.ListDB(), "Выберите базу данных")
				if err != nil {
					props.logger.Fatal(err)
				}
			}
			props.logger.Infof("getting from %s", dbName)

			if key == "" {
				keys, err := props.storage.GetKeys(dbName)
				if err != nil {
					props.logger.Fatal(err)
				}
				key, err = fuzzy(keys, "Выберите значение")
				if err != nil {
					props.logger.Fatal(err)
				}
			}

			value, err := props.storage.GetValue(dbName, key)
			if err != nil {
				props.logger.Fatal(err)
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
