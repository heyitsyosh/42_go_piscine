package piscine

func Split(s, sep string) []string {
	sepRunes := []rune(sep)
	sepLen := runesLen(sepRunes)
	if sepLen == 0 {
		return []string{s}
	}

	var rv []string
	sRunes := []rune(s)
	sLen := SliceLen(sRunes)
	iFrom := 0
	i := 0
	for i <= sLen {
		if i == sLen || (i+sepLen <= sLen && runesEqual(sRunes[i:i+sepLen], sepRunes, sepLen)) {
			rv = append(rv, string(sRunes[iFrom:i]))
			i += sepLen
			iFrom = i
		} else {
			i++
		}
	}
	return rv
}
