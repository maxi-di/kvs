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
			// var err error
			// if dbName == "" {
			// 	dbName, err = fuzzy(props.storage.ListDB(), "Выберите базу данных")
			// 	if err != nil {
			// 		props.logger.Fatal(err)
			// 	}
			// }
			// props.logger.Infof("getting from %s", dbName)

			// if key == "" {
			// 	keys, err := props.storage.GetKeys(dbName)
			// 	if err != nil {
			// 		props.logger.Fatal(err)
			// 	}
			// 	key, err = fuzzy(keys, "Выберите значение")
			// 	if err != nil {
			// 		props.logger.Fatal(err)
			// 	}
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
