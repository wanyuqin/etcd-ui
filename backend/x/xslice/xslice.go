package xslice

type T interface {
	~string | ~int | ~int64 | ~uint | ~int32 | ~int16 | ~int8
}

func UniqueSlice[K T](s []K) []K {
	m := make(map[K]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}

	us := make([]K, 0, len(m))
	for k := range m {
		us = append(us, k)
	}

	return us
}
