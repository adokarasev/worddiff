package worddiff

func worddiff(s1 string, s2 string) string {
	var buff []string
	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				buff = append(buff, string(c1)+worddiff(s1[i+1:], s2[j+1:]))
			}
		}
	}

	ret := ""
	for _, str := range buff {
		if len(ret) < len(str) {
			ret = str
		}
	}

	return ret
}
