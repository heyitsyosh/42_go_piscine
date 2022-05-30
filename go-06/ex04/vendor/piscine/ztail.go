package piscine

import "os"
import "ft"

//get filesize from os.Stat to compare with size (doesn't work with /dev/stdin)
//if using os.ReadFile, toFindLen[size - bytes:]

func Ztail(n string, file string, i int) (bool){
	bytes := atoi(n)
	buf := [1024]byte{}

	fd, o_err := os.Open(file)
	if o_err != nil {
		printError(o_err.Error() + "\n")
		return true
	}
	defer fd.Close()

	toFindLen, rf_err := os.ReadFile(file)
	if rf_err != nil {
		printError(rf_err.Error() + "\n")
		return true
	}

	size := byteLen(toFindLen)
	if bytes > size {
		bytes = size
	}
	fd.Seek(int64(-bytes), 2)

	if i > 3 {
		ft.PrintRune('\n')
	}

	for {
		f, r_err := fd.Read(buf[:])
		if r_err != nil {
			if r_err.Error() == "EOF" {
				break
			}
			printError(r_err.Error() + "\n")
			return true
		}
		printContent(buf[0:f], file, i)
	}
	return false
}

func atoi(nbr string) (uint) {
	ret := uint(0)
	for _, digit := range []rune(nbr) {
		if digit >= '0' && digit <= '9' {
		ret = ret*10 + uint(digit-'0')
		}
	}
	return ret
}

func byteLen(fileBytes []byte) (len uint) {
	for range fileBytes {
		len++
	}
	return
}

func printContent(content []byte, file string, i int) {
	if i > 0 {
		printStr("==> " + file + " <==")
	}

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

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
