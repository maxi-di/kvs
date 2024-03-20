package cmd

import "github.com/sirupsen/logrus"

func selectDB(dbName string, all []string, logger *logrus.Logger) string {
	var err error

	if dbName != "" {
		return dbName
	}

	if len(all) == 0 {
		logger.Warn("no one db's")
		return ""
	}

	dbName, _, err = fuzzy(all, "Выберите базу данных")
	if err != nil {
		logger.Fatal(err)
	}

	if dbName == "" {
		logger.Warn("no db specified")
	}

	return dbName
}
