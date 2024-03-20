package cmd

import (
	"github.com/spf13/cobra"
)

func NewRemoveCmd(props *Props) *cobra.Command {
	dbName := ""
	key := ""
	c := &cobra.Command{
		Use:   "remove",
		Short: "!!!!! TO DO remove Value from DB by Key",
		Run: func(cmd *cobra.Command, args []string) {
			// dbName = selectDB(dbName, props.storage.ListDB(), props.logger)
			// if dbName == "" {
			// 	return
			// }
			// props.logger.Infof("removing from bd '%s'", dbName)

			// if key == "" {
			// 	keys, err := props.storage.GetKeys(dbName)
			// 	if err != nil {
			// 		props.logger.Fatal(err)
			// 	}
			// 	key, _, err = fuzzy(keys, "Что удаляем?")
			// 	if err != nil {
			// 		props.logger.Fatal(err)
			// 	}
			// }
			// if key == "" {
			// 	return
			// }

			// value, err := props.storage.GetValue(dbName, key)
			// if err != nil {
			// 	props.logger.Fatal(err)
			// }
			// fmt.Println(value)
		},
	}
	c.Flags().StringVar(&dbName, "db", "", "db name (with ext)")
	c.Flags().StringVar(&key, "key", "", "key name")

	return c
}
