package tfgenvars

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"path/filepath"
)

func main() {
	err := Run(os.Stdin, os.Stdout, os.Args[1:])
	if err != nil {
		panic(err)
	}
}


func Run(stdin io.Reader, stdout io.Writer, args []string) error {
	var files []string

	// var vars map[string]string
	// var outs map[string]string

	var msg []string

	variableCollector := NewCollector(VariableDeclaration)
	outputCollector := NewCollector(OutputDeclaration)

	if len(args) == 0 {
		text := ""
		scanner := bufio.NewScanner(stdin)
		for scanner.Scan() {
			text += scanner.Text() + "\n"
		}
		if err := scanner.Err(); err != nil {
			return err
		}

		vars := variableCollector.CollectFromText(text)
		outs := outputCollector.CollectFromText(text)

		for k, v := range vars {
			msg = append(msg, fmt.Sprintf("variable \"%s\" %s", k, v))
		}
		for k, v := range outs {
			msg = append(msg, fmt.Sprintf("variable \"%s\" %s", k, v))
		}
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

		varCh := make(chan map[string]string)
		outCh := make(chan map[string]string)
		for _, file := range files {
			go func() { varCh <- variableCollector.CollectFromFile(file) }()
			go func() { outCh <- outputCollector.CollectFromFile(file) }()
		}

		for i := 0; i < len(files) * 2; i++ {
			select {
			case vc := <- varCh:
				for k, v := range vc {
					msg = append(msg, fmt.Sprintf("variable \"%s\" %s", k, v))
				}
			case oc := <-outCh:
				for k, v := range oc {
					msg = append(msg, fmt.Sprintf("variable \"%s\" %s", k, v))
				}
			}
		}
	}

	for _, m := range msg {
		stdout.Write([]byte(fmt.Sprintf(m + "\n")))
	}

	return nil
}
