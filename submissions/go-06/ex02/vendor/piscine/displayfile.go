package piscine

import "os"

func DisplayFile(fd string) {
	content, err := os.ReadFile(fd)
	if err != nil {
		printError(err.Error() + "\n")
		return
	}
	printContent(string(content))
}

func printContent(s string) {
	_, err := os.Stdout.WriteString(s)
	if err != nil {
		os.Exit(1)
	}
}

func printError(s string) {
	errMsg := []byte(s)

	_, err := os.Stderr.Write(errMsg)
	if err != nil {
		os.Exit(1)
	}
}
