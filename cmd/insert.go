package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewInsertCmd(props *Props) *cobra.Command {
	dbName := ""
	key := ""
	value := ""
	force := false
	c := &cobra.Command{
		Use:     "insert",
		Aliases: []string{"set", "push", "add"},
		Short:   "insert Key to DB",
		Run: func(cmd *cobra.Command, args []string) {
			err := openDB(props.storage, dbName)
			if err != nil {
				props.logger.Fatal(err)
			}
			props.logger.Infof("inserting to db '%s'", dbName)

			if key == "" {
				fmt.Println("Input key: ")
				key, _ = readLine()
				if key == "" {
					props.logger.Info("interrupt")
					os.Exit(0)
				}
			}
			if value == "" {
				fmt.Println("Input value: ")
				value, _ = readLine()
				if value == "" {
					props.logger.Info("interrupt")
					os.Exit(0)
				}
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
