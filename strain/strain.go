package strain

type Ints []int

func (i Ints) Keep(fn func(int) bool) (res Ints) {
	for _, v := range i {
		if fn(v) {
			res = append(res, v)
		}
	}
	return
}

func (i Ints) Discard(fn func(int) bool) (res Ints) {
	for _, v := range i {
		if !fn(v) {
			res = append(res, v)
		}
	}
	return
}

type Strings []string

func (i Strings) Keep(fn func(string) bool) (res Strings) {
	for _, v := range i {
		if fn(v) {
			res = append(res, v)
		}
	}
	return
}

type Lists [][]int

func (i Lists) Keep(fn func([]int) bool) (res Lists) {
	for _, v := range i {
		if fn(v) {
			res = append(res, v)
		}
	}
	return
}
