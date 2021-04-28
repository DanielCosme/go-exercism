package accumulate

func Accumulate(list []string, fn func(string) string) (res []string) {
	for _, v := range list {
		res = append(res, fn(v))
	}
	return res
}
