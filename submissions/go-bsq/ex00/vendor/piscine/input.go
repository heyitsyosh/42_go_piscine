package piscine

import "os"

func From_stdin() (string, bool) {
	buf := [1024]byte{}
	input := []byte{}

	n, err := os.Stdin.Read(buf[:])
	for err == nil {
		if n > 0 {
			input = append(input, buf[:n]...)
		}
		n, err = os.Stdin.Read(buf[:])
	}
	if err.Error() == "EOF" {
		return string(input), true
	}
	PrintError(err.Error())
	return "", false
}

func From_file(file string) (string, bool) {
	input, err := os.ReadFile(file)
	if err != nil {
		PrintError(err.Error())
		return "", false
	}

	return string(input), true
}
