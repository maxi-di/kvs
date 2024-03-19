package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewInsertCmd(props *Props) *cobra.Command {
	dbName := ""
	key := ""
	value := ""
	c := &cobra.Command{
		Use:   "insert",
		Short: "insert Key to DB",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			if dbName == "" {
				dbName, err = fuzzy(props.storage.ListDB(), "Выбирите базу данных")
				if err != nil {
					props.logger.Fatal(err)
				}
			}
			props.logger.Infof("inserting to %s", dbName)

			if key == "" {
				fmt.Scanln(key)
			}
			if value == "" {
				fmt.Scanln(value)
			}
			if err = props.storage.Insert(dbName, key, value); err != nil {
				props.logger.Fatal(err)
			}
		},
	}
	c.Flags().StringVar(&dbName, "db", "", "db name (with ext)")
	c.Flags().StringVar(&key, "key", "", "")
	c.Flags().StringVar(&value, "value", "", "")

	return c
}
