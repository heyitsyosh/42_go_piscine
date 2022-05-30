package piscine

import "os"

func FromStdin() {
	buf := [1024]byte{}

	for {
		n, err := os.Stdin.Read(buf[:])
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			printError(err.Error() + "\n")
			return
		}
		printContent(buf[0:n])
	}
}

func FromFile(fd string, argCount int) {
	content, err := os.ReadFile(fd)
	if err != nil {
		printError("ERROR: " + err.Error() + "\n")
		if argCount == 2 {
			os.Exit(1)
		}
		return
	}
	printContent(content)
}

func printContent(content []byte) {
	_, err := os.Stdout.Write(content)
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