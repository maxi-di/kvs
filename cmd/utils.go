package cmd

import (
	"errors"
	"kvs/kvs"
)

func openDB(db kvs.Storage, name string) error {

	var err error

	if name == "" {
		all := db.ListDB()
		if len(all) == 0 {
			return errors.New("no one db's")
		}
		name, _, err = fuzzy(all, "Choose DB from list")
		if err != nil {
			return err
		}
	}

	if name == "" {
		return errors.New("no db specified")
	}

	err = db.Open(name)
	if err != nil {
		return err
	}

	return nil
}
