package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	newEntry = "add new entry..."
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
				if with_add {
					keys = append(keys, newEntry)
				}
				key, _, err = fuzzy(keys, "Choose value")
				if err != nil {
					props.logger.Fatal(err)
				}
			}
			if key == "" {
				props.logger.Warn("key not specified")
				return
			}

			var value string

			if with_add && key == newEntry {
				fmt.Println("Input key: ")
				key, _ = readLine()
				if key == "" {
					props.logger.Info("interrupt")
					os.Exit(0)
				}
				fmt.Println("Input value: ")
				value, _ = readLine()
				if value == "" {
					props.logger.Info("interrupt")
					os.Exit(0)
				}
				err := props.storage.Insert(key, value, false)
				if err != nil {
					props.logger.Fatal(err)
				}

			} else {

				value, err = props.storage.GetValue(key)
				if err != nil {
					props.logger.Fatal(err)
				}
				if value == "" {
					props.logger.Warnf("no value for key '%s'", key)
					return
				}
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
