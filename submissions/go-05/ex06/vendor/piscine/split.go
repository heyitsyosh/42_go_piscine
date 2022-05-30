package piscine

func strLen(s string) (len int) {
	for range s {
		len++
	}
	return
}

func Split(s, sep string) []string {
	r_s, r_sep := []rune(s), []rune(sep)
	sLen, sepLen := strLen(s), strLen(sep)
	lastSep := 0
	var ret []string

	if sepLen == 0 {
		for _, r := range r_s {
			ret = append(ret, string(r))
		}
		return ret
	}

	for i := 0; i < sLen; {
		if string(r_s[i : i + sepLen]) == string(r_sep[0 : sepLen]) {
			if lastSep == 0 && i == 0 { //put in "" before first HA in case of "HA..."
				lastSep = sepLen
				i += sepLen
				ret = append(ret, "")
			} else {
				ret = append(ret, string(r_s[lastSep: i]))
				lastSep = i + sepLen
				i += sepLen
			}
		} else {
			i++
		}
	}

	ret = append(ret, string(r_s[lastSep:sLen])) //add last element
	return ret
}
