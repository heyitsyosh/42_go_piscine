package piscine

import (
	"os"
)

const readBufSize = 1024

// pathの中身をすべて1つのstringとして読み出す
func Read(path string) (string, bool) {
	f, oErr := os.Open(path)
	if oErr != nil {
		PrintStr(oErr.Error() + "\n")
		return "", false
	}
	defer f.Close()
	return ReadFromFile(f)
}

func ReadFromFile(f *os.File) (string, bool) {
	var buffer [readBufSize]byte
	data := []byte{}
	for {
		n, rErr := f.Read(buffer[:])
		if rErr != nil {
			if rErr.Error() == "EOF" {
				break
			}
			PrintStr(rErr.Error() + "\n")
			return "", false
		}
		data = append(data, buffer[:n]...)
	}
	return string(data), true
}
