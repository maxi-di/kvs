package cmd

import (
	"errors"
	"fmt"
	"kvs/kvs"
	"os"

	"github.com/chzyer/readline"
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

func readLine(header string) (string, error) {
	if header != "" {
		fmt.Fprintln(os.Stderr, header)
	}

	l, err := readline.NewEx(&readline.Config{
		Prompt:            "\033[31m>>\033[0m ",
		HistoryFile:       "/tmp/kvs.history",
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
		Stdout:            os.Stderr,
	})
	if err != nil {
		return "", err
	}

	defer l.Close()

	line, err := l.Readline()

	if err == readline.ErrInterrupt {
		return "", nil
	}

	return line, nil
}
