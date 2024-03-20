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
			dbName = selectDB(dbName, props.storage.ListDB(), props.logger)
			if dbName == "" {
				return
			}
			props.logger.Infof("inserting to bd '%s'", dbName)

			if key == "" {
				fmt.Println("Input key: ")
				fmt.Scanln(&key)
			}
			if value == "" {
				fmt.Println("Input value: ")
				fmt.Scanln(&value)
			}
			value, _ := props.storage.GetValue(dbName, key)

			if value == "" {
				if err := props.storage.Insert(dbName, key, value); err != nil {
					props.logger.Fatal(err)
				}
				return
			}

			props.logger.Infof("value already exists %s:%s", key, value)
			fmt.Printf("value already exists %s:%s, update it? [y/n] ", key, value)
			var answer string
			fmt.Scanln(&answer)
		},
	}
	c.Flags().StringVar(&dbName, "db", "", "db name (with ext)")
	c.Flags().StringVar(&key, "key", "", "")
	c.Flags().StringVar(&value, "value", "", "")

	return c
}
