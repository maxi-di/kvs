package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewInsertCmd(props *Props) *cobra.Command {
	dbName := ""
	key := ""
	value := ""
	force := false
	c := &cobra.Command{
		Use:   "insert",
		Short: "insert Key to DB",
		Run: func(cmd *cobra.Command, args []string) {
			err := openDB(props.storage, dbName)
			if err != nil {
				props.logger.Panic(err)
			}
			props.logger.Infof("inserting to db '%s'", dbName)

			if key == "" {
				fmt.Println("Input key: ")
				fmt.Scanln(&key)
			}
			if value == "" {
				fmt.Println("Input value: ")
				fmt.Scanln(&value)
			}

			if err := props.storage.Insert(key, value, force); err != nil {
				props.logger.Fatal(err)
			}
		},
	}
	c.Flags().StringVar(&dbName, "db", "", "db name (with ext)")
	c.Flags().StringVar(&key, "key", "", "")
	c.Flags().StringVar(&value, "value", "", "")
	c.Flags().BoolVarP(&force, "force", "f", false, "force inserting if exist")

	return c
}
