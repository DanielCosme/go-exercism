package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

func (il IntList) Concat(list []IntList) IntList {
	res := il
	for _, v := range list {
		res = res.Append(v)
	}
	return res
}

func (il IntList) Append(a IntList) IntList {
	return append(il, a...)
}

func (il IntList) Reverse() IntList {
	res := IntList{}
	for i := len(il) - 1; i >= 0; i-- {
		v := il[i]
		res = append(res, v)
	}
	return res
}

func (il IntList) Map(fn unaryFunc) IntList {
	res := IntList{}
	for _, v := range il {
		res = append(res, fn(v))
	}
	return res
}

func (il IntList) Length() int {
	return len(il)
}

func (il IntList) Foldl(fn binFunc, acc int) int {
	for _, v := range il {
		acc = fn(acc, v)
	}
	return acc
}

func (il IntList) Foldr(fn binFunc, acc int) int {
	for i := len(il) - 1; i >= 0; i-- {
		v := il[i]
		acc = fn(v, acc)
	}
	return acc
}

func (il IntList) Filter(fn predFunc) IntList {
	res := IntList{}
	for _, v := range il {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}
