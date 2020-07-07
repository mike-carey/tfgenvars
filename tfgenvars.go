package tfgenvars

import (
	"bufio"
	"fmt"
	"io"
	"path/filepath"
)

func Run(stdin io.Reader, stdout io.Writer, args []string) error {
	files := []string{}

	msg := []string{}
	if len(args) == 0 {
		text := ""
		scanner := bufio.NewScanner(stdin)
		for scanner.Scan() {
			text += scanner.Text() + "\n"
		}
		if err := scanner.Err(); err != nil {
			return err
		}

		msg = CollectVariablesFromText(text)
	} else {
		for _, arg := range args {
			matches, err := filepath.Glob(arg)
			if err != nil {
				return err
			}

			for _, match := range matches {
				files = append(files, match)
			}
		}

		messages := make(chan []string)
		for _, file := range files {
			go func() { messages <- CollectVariablesFromFile(file) }()
		}
		msg = <-messages
	}

	for _, m := range msg {
		stdout.Write([]byte(fmt.Sprintf(m + "\n")))
	}

	return nil
}
